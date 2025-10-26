package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"SatohAyaka/leaving-match-experiment/lib/staywatch"
)

func StayWatch() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("3: 退室履歴\n")
		fmt.Print("4: 退室同行割合\n")
		fmt.Print("5: 平均退室時同行人数\n")

		mode, _ := reader.ReadString('\n')
		mode = strings.TrimSpace(mode)

		switch mode {
		case "0":
			fmt.Println("終了します.")
			os.Exit(0)
		case "3":
			staywatch.History()
			return
		case "4":
			staywatch.Ratio()
			return
		case "5":
			staywatch.Average()
			return
		default:
			fmt.Println("エラー: 無効な入力です.3〜5を入力してください.")
		}
	}
}
