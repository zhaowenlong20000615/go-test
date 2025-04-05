package limiter

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisSlidingWindowLimiter 限流器实现-Redis-滑动窗口
type RedisSlidingWindowLimiter struct {
	cmd      redis.Cmdable
	interval time.Duration // 窗口大小
	rate     int           // 阈值
}

// 可以在编译阶段将静态资源文件打包进编译好的程序中，并提供访问这些文件的能力。
//
//go:embed slide_window.lua
var luaScript string

// NewRedisSlidingWindowLimiter
// 参数二：限流的窗口多大
// 参数三：最多多少请求
func NewRedisSlidingWindowLimiter(cmd redis.Cmdable, interval time.Duration, rate int) *RedisSlidingWindowLimiter {
	return &RedisSlidingWindowLimiter{
		cmd:      cmd,
		interval: interval,
		rate:     rate,
	}
}

// Limit 让Redis执行lua脚本
func (r *RedisSlidingWindowLimiter) Limit(ctx context.Context, key string) (bool, error) {
	return r.cmd.Eval(ctx, luaScript, []string{key},
		r.interval.Milliseconds(), r.rate, time.Now().UnixMilli()).Bool()
}
