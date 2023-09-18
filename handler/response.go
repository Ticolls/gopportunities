package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}

type CreateOpeningResponse struct {
	Message string          `json:"message"`
	Data    OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string          `json:"message"`
	Data    OpeningResponse `json:"data"`
}

type GetOpeningByIdResponse struct {
	Message string          `json:"message"`
	Data    OpeningResponse `json:"data"`
}

type GetAllOpeningsResponse struct {
	Message string            `json:"message"`
	Data    []OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string          `json:"message"`
	Data    OpeningResponse `json:"data"`
}
