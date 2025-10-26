package model

type User struct {
	BackendUserId   int64   `json:"BackendUserId"`
	StayWatchUserId *int64  `json:"StayWatchUserId"`
	SlackUserId     *string `json:"SlackUserId"`
	ChannelId       *string `json:"ChannelId"`
	UserName        *string `json:"UserName"`
}

type UseUserData struct {
	BackendUserId   int64
	StayWatchUserId *int64
	UserName        *string
}
