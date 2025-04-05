package ratelimit

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/webook/pkg/limiter"
	"log"
	"net/http"
)

type Builder struct {
	prefix  string
	limiter limiter.Limiter
}

// NewBuilder 限流
// 参数二：限流的窗口多大
// 参数三：最多多少请求
func NewBuilder(l limiter.Limiter) *Builder {
	return &Builder{
		prefix:  "ip-limiter",
		limiter: l,
	}
}

func (b *Builder) Prefix(prefix string) *Builder {
	b.prefix = prefix
	return b
}

// Build 构造Gin的middleware
func (b *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := fmt.Sprintf("%s:%s", b.prefix, ctx.ClientIP())
		/*
			limit内部会调用lua脚本，让redis执行
			万一你和redis之间的网络出现了问题，又或者redis本身出现问题了，就会返回err
			你拿到了这个err，你是限流还是不限流？
			这就有两种做法，
			保守做法：因为是借助于redis做限流，那人家redis都崩了，为了防止系统崩溃，所以选择直接限流。
				ctx.AbortWithStatus(http.StatusInternalServerError)
			激进做法：虽然redis崩溃了，但是这个时候还是要尽量服务正常的用户，所以选择不限流。
				ctx.Next()
		*/
		limited, err := b.limiter.Limit(ctx, key)
		if err != nil {
			log.Println(err)
			// 保守做法
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if limited { // 如果被限流了
			log.Println(err)
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		ctx.Next()
	}
}
