package ucboundary

import (
	"cases/internal/entity"
	"cases/internal/usecase"
	"context"
)

type ChanUseCase interface {
	UploadFiles(ctx context.Context, files []*entity.File) chan *usecase.Result[*entity.File]
}
