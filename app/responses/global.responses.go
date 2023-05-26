package responses

import (
	"math"

	"github.com/gin-gonic/gin"
)

func NewResponses(status int, message string, data interface{}) (response gin.H) {
	return gin.H{
		"status":  status,
		"data":    data,
		"message": message,
	}
}

func PaginationResponse(status int, message string, page int, limit int, total_data int, data interface{}) (response gin.H) {
	return gin.H{
		"data":       data,
		"limit":      limit,
		"total_data": total_data,
		"total_page": math.Ceil(float64(total_data) / float64(limit)),
		"page":       page,
		"message":    message,
		"status":     status,
	}
}
