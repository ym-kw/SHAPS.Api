package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type SubscriptionController struct {
	create usecase.SubscriptionCreater
}

func NewSubscriptionController(c usecase.SubscriptionCreater) *SubscriptionController {
	return &SubscriptionController{
		create: c,
	}
}

func (s *SubscriptionController) Post(c *gin.Context) {
	s.create.Excecute(c)
	c.JSON(200, nil)
}