package controllers

import (
	"log"
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AttendanceController interface {
	CheckIn(c *fiber.Ctx) error
	CheckOut(c *fiber.Ctx) error
}

type attendanceController struct {
	attendanceService services.AttendanceService
}

func NewAttendanceController(attendanceService services.AttendanceService) AttendanceController {
	return attendanceController{attendanceService}
}



func (obj attendanceController) CheckIn(c *fiber.Ctx) error {
	command := commands.CheckInCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = obj.attendanceService.CheckIn(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Check in success",
	})
}

func (obj attendanceController) CheckOut(c *fiber.Ctx) error {
	command := commands.CheckOutCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = obj.attendanceService.CheckOut(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Check out success",
	})
}
