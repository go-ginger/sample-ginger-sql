package models

import gm "github.com/go-ginger/models"

// extend ginger request for future requirements
type GingerRequest struct {
	gm.Request
}

func (request *GingerRequest) SetBaseRequest(req *gm.Request) {
	request.Request = *req
}

func (request *GingerRequest) GetBaseRequest() *gm.Request {
	return &request.Request
}
