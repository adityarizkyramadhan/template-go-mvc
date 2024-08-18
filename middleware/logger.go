package middleware

import (
	"fmt"
	cliLog "log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerRequest(log *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Merekam waktu mulai
		startTime := time.Now()
		// Memproses permintaan
		ctx.Next()

		// Merekam waktu selesai
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Mencatat log
		go log.WithFields(logrus.Fields{
			"method": ctx.Request.Method,
			"url":    ctx.Request.URL.Path,
			"status": ctx.Writer.Status(),
			// buat format waktu jadi second
			"latency": fmt.Sprintf("%f", latency.Seconds()),
			"client":  ctx.ClientIP(),
		}).Info("Request log")

		go cliLog.Printf("%s %s %d %s \n", ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), latency)
	}
}
