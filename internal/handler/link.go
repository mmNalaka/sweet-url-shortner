package handler

import (
	"sweet/internal/service"

	"github.com/gin-gonic/gin"
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

func (h *LinkHandler) CreateLink(ctx *gin.Context) {
}
