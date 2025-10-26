package model

import "time"

type Recommended struct {
	RecommendedId   int64     `json:"RecommendedId"`
	RecommendedTime time.Time `json:"RecommendedTime"`
	MemberIds       []int64   `json:"MemberIds"`
	Status          bool      `json:"Status"`
	CreatedDate     time.Time `json:"CreatedDate"`
}
