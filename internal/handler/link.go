package handler

import (
	"net/http"
	v1 "sweet/api/v1"
	"sweet/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LinkHandler struct {
	*Handler
	linkService service.LinkService
}

func NewLinkHandler(handler *Handler, linkService service.LinkService) *LinkHandler {
	return &LinkHandler{
		Handler:     handler,
		linkService: linkService,
	}
}

func (h *LinkHandler) GetLink(ctx *gin.Context) {

}

// CreateLink godoc
// @Summary Create a link
// @Schemes
// @Description
// @Tags Link module
// @Accept json
// @Produce json
// @Param request body v1.CreateLinkRequest true "params"
// @Success 200 {object} v1.CreateLinkResponse
// @Router /link [post]
func (h *LinkHandler) CreateLink(ctx *gin.Context) {
	req := new(v1.CreateLinkRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	data, err := h.linkService.CreateLink(ctx, req)
	if err != nil {
		h.logger.WithContext(ctx).Error("linkService.CreateLink error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, data)
}
