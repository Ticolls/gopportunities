package handler

import (
	"fmt"
	"net/http"

	"github.com/Ticolls/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "parameter").Error())
		return
	}

	opening := schemas.Opening{}

	//find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "get-opening-by-id", opening)
}
