package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)


func setupLogOutPut()  {
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Log() gin.HandlerFunc {
	setupLogOutPut()
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
