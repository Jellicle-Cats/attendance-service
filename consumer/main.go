package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
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

	// Initialize the AttendanceRepository with the base URL of the external service
	attendanceRepo := repositories.NewAttendanceRepository("https://external-service-url")

	attendanceEventHandler := services.NewAttendanceEventHandler(attendanceRepo)
	attendanceConsumerHandler := services.NewConsumerHandler(attendanceEventHandler)

	fmt.Println("Attendance consumer started...")
	for {
		consumer.Consume(context.Background(), events.Topics, attendanceConsumerHandler)
	}
}
