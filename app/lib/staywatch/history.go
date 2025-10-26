package staywatch

import (
	"SatohAyaka/leaving-match-experiment/api"
	"SatohAyaka/leaving-match-experiment/model"
	"SatohAyaka/leaving-match-experiment/utils"
	"log"
)

func History() []model.LogJSON {
	//履歴取得（user-id, offset, limit）
	//user-id指定なし，offsetを回す，limitはデフォルト30件
	//一旦最新30件で回してる
	Logs, err := api.GetLogs()
	if err != nil {
		log.Fatalf("取得失敗: %v", err)
	}
	converted := utils.AccompanyJadgement(Logs)
	return converted
}
