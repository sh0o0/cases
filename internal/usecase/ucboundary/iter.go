package ucboundary

import (
	"cases/internal/entity"
	"context"
	"iter"
)

type IterUseCase interface {
	UploadFiles(ctx context.Context, files []*entity.File) iter.Seq2[*entity.File, error]
}
