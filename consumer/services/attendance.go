package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type attendanceEventHandler struct {
	attendanceRepo repositories.AttendanceRepository
}

func NewAttendanceEventHandler(attendanceRepo repositories.AttendanceRepository) EventHandler {
	return attendanceEventHandler{attendanceRepo}
}

func (obj attendanceEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.CheckInEvent{}).Name():
		event := &events.CheckInEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		attendanceRecord := repositories.AttendanceRecord{
			UserID:    event.UserID,
			BookingID: event.BookingID,
		}
		err = obj.attendanceRepo.RecordCheckIn(attendanceRecord)

		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)
	case reflect.TypeOf(events.CheckOutEvent{}).Name():
		event := &events.CheckOutEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		attendanceRecord := repositories.AttendanceRecord{
			UserID:    event.UserID,
			BookingID: event.BookingID,
		}
		err = obj.attendanceRepo.RecordCheckOut(attendanceRecord)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)
	case reflect.TypeOf(events.ExtentTimeEvent{}).Name():
		event := &events.ExtentTimeEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		attendanceRecord := repositories.AttendanceRecord{
			UserID:    event.UserID,
			BookingID: event.BookingID,
		}
		err = obj.attendanceRepo.RecordExtentTime(attendanceRecord)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)
	default:
		log.Println("no event handler")
	}
}
