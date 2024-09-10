package ucboundary

import (
	"cases/internal/entity"
	"context"
)

type SeparatedUseCase interface {
	UploadFiles(ctx context.Context, files []*entity.File) error
}
