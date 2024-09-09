package usecase

import (
	"cases/internal/entity"
	"context"
)

type Uploader interface {
	Upload(ctx context.Context, file *entity.File) error
}
