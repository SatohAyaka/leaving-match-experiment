package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"SatohAyaka/leaving-match-experiment/lib"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("取得したいデータを選んでください\n")
	fmt.Print("1: 滞在ウォッチ\n")
	fmt.Print("2: leaving-match\n")

	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode) // 改行文字を削除

	switch mode {
	case "1":
		lib.StayWatch()
	case "2":
		lib.LeavingMatch()
	default:
		fmt.Println("エラー: 無効な入力です.1または2を入力してください.")
	}
}
