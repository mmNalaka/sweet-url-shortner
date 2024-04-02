package handler

import (
	"github.com/gin-gonic/gin"
	"sweet/internal/service"
)

type LinkHandler struct {
	*Handler
	linkService service.LinkService
}

func NewLinkHandler(handler *Handler, linkService service.LinkService) *LinkHandler {
	return &LinkHandler{
		Handler:      handler,
		linkService: linkService,
	}
}

func (h *LinkHandler) GetLink(ctx *gin.Context) {

}
