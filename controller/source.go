package controller

import (
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
)

type SourceController interface {
	Token(ctx *gin.Context)
	Upload(ctx *gin.Context) // 回调
}

type sourceController struct {
	sourceService service.SourceService
}

func NewSourceController(sourceService service.SourceService) SourceController {
	return &sourceController{
		sourceService: sourceService,
	}
}

func (s *sourceController) Token(ctx *gin.Context) {

}

func (s *sourceController) Upload(ctx *gin.Context) {

}
