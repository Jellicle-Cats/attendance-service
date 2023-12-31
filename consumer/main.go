package main

import (
	"consumer/proto"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"log"
	"os"
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
	// if err := viper.ReadInConfig(); err != nil {
	// 	panic(err)
	// }

}

func main() {
	kafkaServersEnv := os.Getenv("KAFKA_SERVERS")
	kafkaServers := strings.Split(kafkaServersEnv, ",")
	// consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	// consumer, err := sarama.NewConsumerGroup(kafkaServers, viper.GetString("kafka.group"), nil)
	consumer, err := sarama.NewConsumerGroup(kafkaServers, "attendance", nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	creds := insecure.NewCredentials()
	// cc, err  := grpc.Dial(viper.GetString("booking.server"), grpc.WithTransportCredentials(creds))
	cc, err  := grpc.Dial(os.Getenv("BOOKING_SERVER"), grpc.WithTransportCredentials(creds))

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
