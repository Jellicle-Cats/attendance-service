package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AttendanceRecord struct {
	UserID    string
	BookingID string
}

type AttendanceRepository interface {
	RecordCheckIn(record AttendanceRecord) error
	RecordCheckOut(record AttendanceRecord) error
	RecordExtentTime(record AttendanceRecord) error
}

type attendanceRepository struct {
	baseURL string
}

func NewAttendanceRepository(baseURL string) AttendanceRepository {
	return &attendanceRepository{baseURL}
}

func (r *attendanceRepository) RecordCheckIn(record AttendanceRecord) error {
	fmt.Println("RecordCheckIn")
	return r.sendHTTPRequest("POST", "/checkin", record)
}

func (r *attendanceRepository) RecordCheckOut(record AttendanceRecord) error {
	fmt.Println("RecordCheckOut")
	return r.sendHTTPRequest("POST", "/checkout", record)
}


func (r *attendanceRepository) RecordExtentTime(record AttendanceRecord) error {
	fmt.Println("RecordExtentTime")
	return r.sendHTTPRequest("POST", "/extenttime", record)
}

func (r *attendanceRepository) sendHTTPRequest(method, endpoint string, data interface{}) error {
	url := r.baseURL + endpoint

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	return nil
}
