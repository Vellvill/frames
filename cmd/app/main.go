package main

import (
	"context"

	"github.com/Vellvill/frames/internal/database"
	goodsService "github.com/Vellvill/frames/internal/database/good"
	imagesRepo "github.com/Vellvill/frames/internal/database/image"
	. "github.com/Vellvill/frames/internal/logger"
	goodsRepo "github.com/Vellvill/frames/internal/service/goods"
	imagesService "github.com/Vellvill/frames/internal/service/images"
)

func main() {
	ctx := context.Background()

	db, err := database.New(ctx)
	if err != nil {
		Logger.Fatal(err)
	}

	_ = imagesService.NewImageService(imagesRepo.NewImageRepository(db))
	_ = goodsRepo.NewGoodsService(goodsService.NewGoodsRepository(db))

}
