package middleware

import (
	"fmt"
	"log"

	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

// ErrorHandler adalah middleware untuk menangkap dan menangani error
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Jalankan permintaan dan tangkap semua error
		ctx.Next()
		// Tangkap error jika ada
		err := ctx.Errors.Last()
		if err != nil {
			// Tentukan status code dan pesan error yang sesuai
			errParse := utils.ParseError(err.Error())
			if errParse.StatusCode >= 500 {
				// Format [API-Muhammadiyah] IP Address - Method - Path - Status Code - Message
				message := fmt.Sprintf("[API-Muhammadiyah] %s - %s - %s - %d - %s",
					ctx.ClientIP(),
					ctx.Request.Method,
					ctx.Request.URL.Path,
					errParse.StatusCode,
					errParse.Message,
				)
				// Kirim pesan error ke Telegram
				if err := utils.SendTelegramMessage(message); err != nil {
					log.Println("Failed to send message to Telegram")
				}
			}
			utils.ErrorResponse(ctx, errParse.StatusCode, errParse.Message)
		}
	}
}
