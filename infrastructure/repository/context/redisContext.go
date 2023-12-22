package context

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/redis"
)

// RedisContext Redis实例
var RedisContext *redisContext

type redisContext struct {
	redis.IClient `inject:"default"` // 使用farseer.yaml的Redis.default配置节点，并自动注入
}

// InitRedisContext 初始化上下文
func InitRedisContext() {
	RedisContext = container.ResolveIns(&redisContext{})
}
