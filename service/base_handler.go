package service

import "github.com/gin-gonic/gin"

type Handler interface {
	Handle(ctx *gin.Context)
}
