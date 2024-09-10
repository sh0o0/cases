package infrastructure

import (
	"cases/internal/entity"
	"context"
	"time"
)

type Uploader struct{}

func (u *Uploader) Upload(ctx context.Context, file *entity.File) error {
	time.Sleep(1 * time.Second)
	return nil
}
