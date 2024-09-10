package usecase

import (
	"cases/internal/entity"
	"context"
	"fmt"
)

type CurrentUseCase struct {
	Uploader  Uploader
	Messenger Messenger
}

func (u *CurrentUseCase) UploadFiles(ctx context.Context, files []*entity.File) error {
	for _, file := range files {
		if err := u.Uploader.Upload(ctx, file); err != nil {
			if err := u.Messenger.Send(ctx, fmt.Sprintf("Failed: %s", file.Name)); err != nil {
				return err
			}
		} else {
			if err := u.Messenger.Send(ctx, fmt.Sprintf("Success: %s", file.Name)); err != nil {
				return err
			}
		}

	}

	return nil
}
