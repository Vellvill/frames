package image

import "database/sql"

type Image struct {
	ID        int64        `db:"id"`
	KeyS3     int64        `db:"key_s3"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
