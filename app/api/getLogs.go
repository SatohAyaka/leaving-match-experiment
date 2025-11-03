package api

import (
	"SatohAyaka/leaving-match-experiment/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetLogs(userId int64, limit int64) (model.GetLogResponse, error) {
	godotenv.Load("../.env")

	apiKey := os.Getenv("API_KEY")
	baseUrl := os.Getenv("STAY_WATCH_URL")
	endpoint := os.Getenv("STAY_WATCH_LOGS")

	apiUrl, err := url.Parse(fmt.Sprintf("%s%s", baseUrl, endpoint))
	if err != nil {
		return model.GetLogResponse{}, fmt.Errorf("URLの構築に失敗しました: %v", err)
	}
	q := apiUrl.Query()
	if userId != 0 {
		q.Set("user-id", strconv.FormatInt(userId, 10))
	}
	if limit != 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	apiUrl.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", apiUrl.String(), nil)
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
