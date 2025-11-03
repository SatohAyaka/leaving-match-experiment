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

func GetUserData(name string, staywatchId int64) (model.User, error) {
	godotenv.Load("../.env")
	baseUrl := os.Getenv("LEAVING_MATCH_API")
	endpoint := os.Getenv("LEAVING_MATCH_USER")
	apiUrl, err := url.Parse(fmt.Sprintf("%s%s", baseUrl, endpoint))
	if err != nil {
		return model.User{}, fmt.Errorf("URLの構築に失敗しました: %v", err)
	}

	q := apiUrl.Query()
	if name != "" {
		q.Set("name", name)
	}
	if staywatchId != 0 {
		q.Set("staywatch", strconv.FormatInt(staywatchId, 10))
	}
	apiUrl.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", apiUrl.String(), nil)
	if err != nil {
		return model.User{}, fmt.Errorf("リクエスト作成に失敗しました: %v", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return model.User{}, fmt.Errorf("リクエストに失敗しました: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return model.User{}, fmt.Errorf("APIエラー: ステータスコード %d", res.StatusCode)
	}
	var response []model.User
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return model.User{}, fmt.Errorf("レスポンスのデコードに失敗しました: %v", err)
	}

	if len(response) == 0 {
		return model.User{}, fmt.Errorf("ユーザーデータが存在しません")
	}
	return response[0], nil
}
