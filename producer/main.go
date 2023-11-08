package main

import (
	"os"
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	kafkaServersEnv := os.Getenv("KAFKA_SERVERS")
	kafkaServers := strings.Split(kafkaServersEnv, ",")
	// producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	producer, err := sarama.NewSyncProducer(kafkaServers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAttendanceService(eventProducer)
	accountController := controllers.NewAttendanceController(accountService)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	app.Post("/checkin", accountController.CheckIn)
	app.Post("/checkout", accountController.CheckOut)
	app.Listen(":3456")
}