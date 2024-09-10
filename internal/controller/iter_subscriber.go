package controller

import (
	"cases/internal/usecase/ucboundary"
	"context"
)

type IterSubscriber struct {
	UseCase   ucboundary.IterUseCase
	Presenter Presenter
}

func (s *IterSubscriber) UploadFiles(ctx context.Context, event *Event) error {
	files := event.toFiles()

	out := s.UseCase.UploadFiles(ctx, files)
	for file, err := range out {
		if err != nil {
			if err := s.Presenter.Error(ctx, err); err != nil {
				return err
			}
		} else {
			if err := s.Presenter.Success(ctx, file); err != nil {
				return err
			}
		}
	}

	return nil
}
