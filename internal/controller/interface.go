package controller

import (
	"cases/internal/entity"
	"context"
)

type Presenter interface {
	Success(ctx context.Context, file *entity.File) error
	Error(ctx context.Context, err error) error
}
