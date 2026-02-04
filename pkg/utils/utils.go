package utils

import (
	"fmt"
	"time"
)

var (
	debugEnabled bool
)

func SetDebugMode(debug bool) {
	debugEnabled = debug
}

func Log(info string) {
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%v : %v\n", date, info)
}

func LogError(info string, err any) {
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("\033[31m %v : %v : %v\n\033[0m", date, info, err)
}

func Debug(info string) {
	if debugEnabled {
		date := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("(DEBUG) %v : %v\n", date, info)
	}

}
