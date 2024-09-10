package controller

import (
	"cases/internal/usecase/ucboundary"
	"context"
)

type CurrentSubscriber struct {
	UseCase ucboundary.CurrentUseCase
}

func (s *CurrentSubscriber) UploadFiles(ctx context.Context, event *Event) error {
	files := event.toFiles()

	if err := s.UseCase.UploadFiles(ctx, files); err != nil {
		return err
	}
	return nil
}
