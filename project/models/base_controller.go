package models

import (
	"github.com/go-ginger/ginger"
	gm "github.com/go-ginger/models"
)

type BaseController struct {
	ginger.BaseController
}

type BaseItemController struct {
	ginger.BaseItemController
}

type BaseItemsController struct {
	ginger.BaseItemsController
}

func (c *BaseController) GetRequestSample() gm.IRequest {
	return &GingerRequest{}
}

func (c *BaseItemController) GetRequestSample() gm.IRequest {
	return &GingerRequest{}
}

func (c *BaseItemsController) GetRequestSample() gm.IRequest {
	return &GingerRequest{}
}
