package goods

import (
	"github.com/Vellvill/frames/internal/database/good"
)

type GoodsService interface {
}

type goodsService struct {
	goodsRepo good.Repository
}

func NewGoodsService(repo good.Repository) GoodsService {
	return &goodsService{goodsRepo: repo}
}
