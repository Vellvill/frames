package image

import "github.com/Vellvill/frames/internal/database"

type Repository interface {
	database.Pool
}

type repository struct {
	database.Pool
}

func NewImageRepository(pool database.Pool) Repository {
	return &repository{pool}
}

func (s *repository) BatchCreate(images ...Image) error {
	if len(images) == 0 {
		return nil
	}

	q := database.PgQb().
		Insert(tableName).
		Columns("id_s3")

	for i := range images {
		q = q.Values(images[i].IDS3)
	}

	return s.Exec(q)
}
