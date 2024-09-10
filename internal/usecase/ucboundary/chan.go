package ucboundary

import (
	"cases/internal/entity"
	"context"
)

type Result[T any] struct {
	Value T
	Err   error
}

type ChanUseCase interface {
	UploadFiles(ctx context.Context, files []*entity.File) chan *Result[*entity.File]
}
