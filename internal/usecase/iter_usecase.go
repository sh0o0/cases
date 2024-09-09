package usecase

import (
	"cases/internal/entity"
	"context"
	"iter"
)

type IterUseCase struct {
	Uploader Uploader
}

func (u *IterUseCase) UploadFiles(ctx context.Context, files []*entity.File) iter.Seq2[*entity.File, error] {
	return func(yield func(*entity.File, error) bool) {
		for _, file := range files {
			if err := u.Uploader.Upload(ctx, file); err != nil {
				if !yield(nil, err) {
					return
				}
			} else {
				if !yield(file, nil) {
					return
				}

			}
		}
	}
}
