package utils

import (
	"fmt"
	"time"
)

func Log(info string) {
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%v : %v\n", date, info)
}

func LogError(info string, err any) {
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%v : %v : %v\n", date, info, err)
}
