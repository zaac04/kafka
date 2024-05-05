package utilities

import (
	"log"
)

func LogError(err error, ctx string) {
	log.Println(ctx, ":", err)
}
