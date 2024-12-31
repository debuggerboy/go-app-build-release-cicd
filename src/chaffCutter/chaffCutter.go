package chaffCutter

import (
	"log"
	"strconv"
	"strings"
)

func SnipSpaceFromString(straw string) string {
	chaff := strings.TrimSpace(straw)
	return chaff
}

func SnipStringToInt64(straw string) int64 {
	chaff, err := strconv.ParseInt(strings.TrimSpace(straw), 0, 32)
	if err != nil {
		log.Println("Error in chaffCutter.SnipStringToInt64")
		log.Fatalln(err)
	}
	return chaff
}
