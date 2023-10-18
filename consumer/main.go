package main

import (
	"consumer/proto"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"log"
	"strings"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func main() {
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	creds := insecure.NewCredentials()
	cc, err  := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	bookingClient := proto.NewBookingServiceClient(cc)
	attendanceEventHandler := services.NewAttendanceEventHandler(bookingClient)
	attendanceConsumerHandler := services.NewConsumerHandler(attendanceEventHandler)

	fmt.Println("Attendance consumer started...")
	for {
		consumer.Consume(context.Background(), events.Topics, attendanceConsumerHandler)
	}
}
