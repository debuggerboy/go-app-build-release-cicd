package errorHandler

import (
	"log"
)

func CheckError(err any, ident string) {
	if err != nil {
		log.Println(":", ident, ":", err)
	}
}
