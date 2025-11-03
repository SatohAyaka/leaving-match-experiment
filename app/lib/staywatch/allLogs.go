package staywatch

import (
	"SatohAyaka/leaving-match-experiment/api"
	"SatohAyaka/leaving-match-experiment/model"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func AllLogs() []model.LogJSON {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ユーザ名:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	userData, err := api.GetUserData(name, 0)
	if err != nil {
		log.Fatalf("取得失敗: %v", err)
	}
	staywatchId := *userData.StayWatchUserId

	fmt.Print("表示数:")
	limitStr, _ := reader.ReadString('\n')
	limitStr = strings.TrimSpace(limitStr)
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		fmt.Println("数値を入力してください")
		return nil
	}

	response, err := api.GetLogs(staywatchId, limit)
	if err != nil {
		log.Fatalf("取得失敗: %v", err)
	}
	logs := response.Logs
	var converted []model.LogJSON
	for i := 0; i < len(logs); i++ {
		if logs[i].EndAt == "2016-01-01 00:00:00" {
			continue
		}
		converted = append(converted, logs[i])
	}
	return converted
}
