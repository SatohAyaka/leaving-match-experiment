package utils

import (
	"SatohAyaka/leaving-match-experiment/model"
	"time"
)

func AccompanyJadgement(data model.GetLogResponse) []model.LogJSON {
	layout := "2006-01-02 15:04:05"
	logs := data.Logs
	var result []model.LogJSON

	for i := 0; i < len(logs); i++ {
		endI, err := time.Parse(layout, logs[i].EndAt)
		if err != nil {
			continue
		}

		for j := i + 1; j < len(logs); j++ {
			endJ, err := time.Parse(layout, logs[j].EndAt)
			if err != nil {
				continue
			}

			diff := endI.Sub(endJ)
			if diff < 0 {
				diff = -diff // 絶対値を取る
			}

			// 1分以内の差なら抽出
			if diff <= time.Minute {
				result = append(result, logs[i], logs[j])
			}
		}
	}
	unique := make(map[int64]model.LogJSON)
	for _, log := range result {
		unique[log.ID] = log
	}

	uniqueLogs := make([]model.LogJSON, 0, len(unique))
	for _, log := range unique {
		uniqueLogs = append(uniqueLogs, log)
	}

	return uniqueLogs
}
