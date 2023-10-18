package services

import (
	"consumer/proto"
	"context"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type attendanceEventHandler struct {
	bookingClient proto.BookingServiceClient
}

func NewAttendanceEventHandler(bookingClient proto.BookingServiceClient) EventHandler {
	return attendanceEventHandler{bookingClient}
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


		req := &proto.UpdateBookingStatusRequest{
			Id:     &proto.BookingId{Id: event.BookingID},
			Status: proto.BookingStatusEnum_CHECKED_IN,
		}

		_, err = obj.bookingClient.UpdateBookingStatus(context.Background(), req)

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
		req := &proto.UpdateBookingStatusRequest{
			Id:     &proto.BookingId{Id: event.BookingID},
			Status: proto.BookingStatusEnum_COMPLETED,
		}

		_, err = obj.bookingClient.UpdateBookingStatus(context.Background(), req)
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

		req := &proto.BookingRequest{
			User: &proto.UserId{UserId: event.UserID},
			BookingTime: &proto.BookingTime{
				StartTime: event.StartTime,
				EndTime:   event.EndTime,
			},
			Seat: &proto.Seat{SeatId: event.SeatID},
			Status: proto.BookingStatusEnum_CHECKED_IN,
		}
		_, err = obj.bookingClient.CreateBooking(context.Background(), req)

		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)
	default:
		log.Println("no event handler")
	}
}
