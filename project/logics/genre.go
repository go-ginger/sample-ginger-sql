package logics

import (
	gl "github.com/go-ginger/logic"
	gm "github.com/go-ginger/models"
	"github.com/go-ginger/sample-ginger-sql/project/models"
)

type genreLogic struct {
	gl.IBaseLogicHandler
}

func (l *genreLogic) Model(request gm.IRequest) {
	req := request.(*models.GingerRequest)
	req.Model = &models.Genre{}
}

func (l *genreLogic) Models(request gm.IRequest) {
	req := request.(*models.GingerRequest)
	req.Models = &[]models.Genre{}
}
