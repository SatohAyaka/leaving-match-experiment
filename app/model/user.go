package model

type StayWatchUser struct {
	StayWatchUserId int64  `json:"id"`
	Name            string `json:"name"`
	Tags            []struct {
		TagID   int    `json:"id"`
		TagName string `json:"name"`
	} `json:"tags"`
}
