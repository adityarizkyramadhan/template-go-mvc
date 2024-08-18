package utils

import "github.com/gin-gonic/gin"

// ErrorResponseData represents the structure of an error response.
// swagger:model ErrorResponseData
type ErrorResponseData struct {
	// The error message.
	// Example: "Invalid request"
	Message string `json:"message"`

	// Example: nil
	Data interface{} `json:"data"`
}

// SuccessResponseData represents the structure of a success response.
// swagger:model SuccessResponseData
type SuccessResponseData struct {
	// The success message.
	// Example: "success"
	Message string `json:"message"`

	// The actual data returned.
	Data interface{} `json:"data"`
}

func ErrorResponse(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{
		"message": message,
		"data":    nil,
	})
}

func SuccessResponse(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, gin.H{
		"message": "success",
		"data":    data,
	})
}
