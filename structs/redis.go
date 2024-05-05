package structs

import (
	"context"
	"time"
)

type SetCache struct {
	Ctx    context.Context
	Key    string
	Val    interface{}
	Expiry time.Duration
}
type GetCache struct {
	Ctx context.Context
	Key string
}
