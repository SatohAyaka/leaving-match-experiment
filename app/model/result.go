package model

import "time"

type Result struct {
	ResultId    int64     `json:"ResultId"`
	BusTimeId   int64     `json:"BusTimeId"`
	BusTime     time.Time `json:"BusTime"`
	Member      int64     `json:"Member"`
	CreatedDate time.Time `json:"CreatedDate"`
}
