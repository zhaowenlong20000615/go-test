package limiter

import "context"

// Limiter 限流器接口
type Limiter interface {
	// Limit 是否触发限流，返回 true 就是触发限流
	Limit(ctx context.Context, key string) (bool, error)
}
