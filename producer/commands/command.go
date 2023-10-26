package commands

type CheckInCommand struct {
	UserID int64
	BookingID int64
}

type CheckOutCommand struct {
	UserID int64
	BookingID int64
}
