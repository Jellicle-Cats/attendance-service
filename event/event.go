package events

import (
	"reflect"
)

var Topics = []string{
	reflect.TypeOf(CheckInEvent{}).Name(),
	reflect.TypeOf(CheckOutEvent{}).Name(),
	reflect.TypeOf(ExtentTimeEvent{}).Name(),
}


type CheckInEvent struct {
	UserID string
	BookingID string
}

type CheckOutEvent struct {
	UserID string
	BookingID string
}

type ExtentTimeEvent struct {
	UserID string
	BookingID string
}