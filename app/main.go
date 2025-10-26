package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("取得したいデータを選んでください")
	fmt.Print("1: 滞在ウォッチ")
	fmt.Print("2: leaving-match")

	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode) // 改行文字を削除

	switch mode {
	case "1":
	case "2":
	default:
		fmt.Println("エラー: 無効な入力です。1または2を入力してください。")
		os.Exit(1)
	}
}
