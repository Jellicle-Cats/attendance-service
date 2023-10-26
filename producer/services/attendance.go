package services

import (
	"errors"
	"events"
	"log"
	"producer/commands"
)

type AttendanceService interface {
	CheckIn(command commands.CheckInCommand) error
	CheckOut(command commands.CheckOutCommand) error
}

type attendanceService struct {
	eventProducer EventProducer
}

func NewAttendanceService(eventProducer EventProducer) AttendanceService {
	return attendanceService{eventProducer}
}


func (obj attendanceService) CheckIn(command commands.CheckInCommand) error {
	if command.UserID == 0 || command.BookingID == 0 {
		return errors.New("bad request")
	}

	event := events.CheckInEvent{
		UserID:     command.UserID,
		BookingID: command.BookingID,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj attendanceService) CheckOut(command commands.CheckOutCommand) error {
	if command.UserID == 0 || command.BookingID == 0 {
		return errors.New("bad request")
	}

	event := events.CheckOutEvent{
		UserID:     command.UserID,
		BookingID: command.BookingID,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}
