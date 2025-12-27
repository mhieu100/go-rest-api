package response

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageMeta struct {
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	Limit       int   `json:"limit"`
	TotalItems  int64 `json:"total_items"`
}

type PagedData struct {
	Items interface{} `json:"items"`
	Meta  PageMeta    `json:"meta"`
}

func SuccessWithPagination(c *gin.Context, items interface{}, totalItems int64, page, limit int) {
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data: PagedData{
			Items: items,
			Meta: PageMeta{
				CurrentPage: page,
				TotalPages:  totalPages,
				Limit:       limit,
				TotalItems:  totalItems,
			},
		},
	})
}
