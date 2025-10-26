package api

import (
	"SatohAyaka/leaving-match-experiment/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetLogs() (model.GetLogResponse, error) {
	godotenv.Load("../.env")

	apiKey := os.Getenv("API_KEY")
	baseUrl := os.Getenv("STAY_WATCH_URL")
	endpoint := os.Getenv("STAY_WATCH_LOGS")

	apiUrl := fmt.Sprintf("%s%s", baseUrl, endpoint)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return model.GetLogResponse{}, fmt.Errorf("リクエスト作成に失敗しました: %v", err)
	}

	// Header に API キーを付与
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return model.GetLogResponse{}, fmt.Errorf("リクエストに失敗しました: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return model.GetLogResponse{}, fmt.Errorf("APIエラー: ステータスコード %d", res.StatusCode)
	}

	// JSONデコード
	var response model.GetLogResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return model.GetLogResponse{}, fmt.Errorf("レスポンスのデコードに失敗しました: %v", err)
	}

	return response, nil
}
