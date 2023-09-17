package handler

import (
	"fmt"
	"net/http"

	"github.com/Ticolls/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary get opening by id
// @Description show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path string true "Opening identification"
// @Success 200 {object} GetOpeningByIdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [get]
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
