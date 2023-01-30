package images

import "github.com/Vellvill/frames/internal/database/image"

type ImagesService interface {
}

type imagesService struct {
	imagesRepo image.Repository
}

func NewImageService(repo image.Repository) ImagesService {
	return &imagesService{imagesRepo: repo}
}
