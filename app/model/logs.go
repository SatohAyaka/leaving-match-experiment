package model

type LogJSON struct {
	ID      int64  `json:"id"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
	Room    string `json:"room"`
	Name    string `json:"name"`
}

type GetLogResponse struct {
	Logs  []LogJSON `json:"logs"`
	Count int64     `json:"count"`
}
