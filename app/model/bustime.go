package model

import "time"

type BusTime struct {
	BusTimeId     int64     `json:"BusTimeId"`
	RecommendedId int64     `json:"RecommendedId"`
	PreviousTime  time.Time `json:"PreviousTime"`
	NearestTime   time.Time `json:"NearestTime"`
	NextTime      time.Time `json:"NextTime"`
	CreatedDate   time.Time `json:"CreatedDate"`
	EndTime       time.Time `json:"EndTime"`
}
