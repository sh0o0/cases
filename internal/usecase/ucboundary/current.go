package ucboundary

import (
	"cases/internal/entity"
	"context"
)

type CurrentUseCase interface {
	UploadFiles(ctx context.Context, files []*entity.File) error
}
