package utilities

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type ContextKey string

const (
	ginKey = "gin.context.key"
)

var ginContextMap map[string]*gin.Context

func init() {
	ginContextMap = map[string]*gin.Context{}
}

func AddGinContext(ctx context.Context, c *gin.Context) context.Context {
	key := generateNewKey()

	ctx = setResKey(ctx, key)

	ginContextMap[key] = c

	return ctx
}

func GetGinContext(ctx context.Context) *gin.Context {
	key := getResKey(ctx)
	res, ok := ginContextMap[key]
	if !ok {
		_, cancel := context.WithCancel(ctx)
		cancel()
	}
	return res
}

func DeleteGinContext(ctx context.Context) {
	key := getResKey(ctx)
	if _, ok := ginContextMap[key]; ok {
		delete(ginContextMap, key)
	}
}

func getResKey(ctx context.Context) string {
	return getKey(ctx, ContextKey(ginKey))
}

func setResKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextKey(ginKey), value)
}

func getKey(ctx context.Context, ctxKey ContextKey) string {
	key, _ := ctx.Value(ctxKey).(string)
	return key
}

func generateNewKey() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Int())
}
