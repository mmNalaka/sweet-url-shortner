package handler

import (
	"github.com/gin-gonic/gin"
	"sweet/internal/service"
)

type DoaminHandler struct {
	*Handler
	doaminService service.DoaminService
}

func NewDoaminHandler(handler *Handler, doaminService service.DoaminService) *DoaminHandler {
	return &DoaminHandler{
		Handler:      handler,
		doaminService: doaminService,
	}
}

func (h *DoaminHandler) GetDoamin(ctx *gin.Context) {

}
