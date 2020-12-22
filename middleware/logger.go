package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[MYAPP] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
			param.Latency,
			param.ClientIP,
			param.MethodColor(), param.Method, param.ResetColor(),
			param.Path, param.ErrorMessage)

	})
}
