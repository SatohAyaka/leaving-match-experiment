package model

import "time"

type Vote struct {
	VoteId        int64     `json:"VoteId"`
	BusTimeId     int64     `json:"BusTimeId"`
	BackendUserId int64     `json:"BackendUserId"`
	Previous      bool      `json:"Previous"`
	Nearest       bool      `json:"Nearest"`
	Next          bool      `json:"Next"`
	CreatedDate   time.Time `json:"CreatedDate"`
}
