package good

import "github.com/Vellvill/frames/internal/database"

type Repository interface {
	database.Pool
}

type repository struct {
	database.Pool
}

func NewGoodsRepository(pool database.Pool) Repository {
	return &repository{pool}
}
