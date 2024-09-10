package usecase

import (
	"cases/internal/entity"
	"cases/internal/usecase/ucboundary"
	"context"
)

var _ ucboundary.SeparatedUseCase = (*SeparatedUseCase)(nil)

type SeparatedUseCase struct {
	Uploader Uploader
}

func (u *SeparatedUseCase) UploadFiles(ctx context.Context, files []*entity.File) error {
	for _, file := range files {
		if err := u.Uploader.Upload(ctx, file); err != nil {
			return err
		}
	}

	return nil
}
