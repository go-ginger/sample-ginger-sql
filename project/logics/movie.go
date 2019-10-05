package logics

import (
	gl "github.com/go-ginger/logic"
	gm "github.com/go-ginger/models"
	"github.com/go-ginger/sample-ginger-sql/project/models"
)

type movieLogic struct {
	gl.IBaseLogicHandler
}

func (l *movieLogic) Model(request gm.IRequest) {
	req := request.(*models.GingerRequest)
	req.Model = &models.Movie{}
}

func (l *movieLogic) Models(request gm.IRequest) {
	req := request.(*models.GingerRequest)
	req.Models = &[]models.Movie{}
}
