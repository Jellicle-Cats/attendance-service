package events

import (
	"reflect"
)

var Topics = []string{
	reflect.TypeOf(CheckInEvent{}).Name(),
	reflect.TypeOf(CheckOutEvent{}).Name(),
	reflect.TypeOf(ExtentTimeEvent{}).Name(),
}

type Event interface{}


type CheckInEvent struct {
	UserID int64
	BookingID int64
}

type CheckOutEvent struct {
	UserID int64
	BookingID int64
}

type ExtentTimeEvent struct {
	UserID     int64
	BookingID  int64
	StartTime  int64
	EndTime    int64
	SeatID     int32
}
