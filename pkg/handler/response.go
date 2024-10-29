package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResp struct {
	Massage string `json:"message"`
}

type StatusResp struct{
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, massage string){
	logrus.Errorf(massage)
	c.AbortWithStatusJSON(statusCode, ErrorResp{massage})
}