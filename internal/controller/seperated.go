package controller

import (
	"cases/internal/usecase/ucboundary"
	"context"
)

type SeparatedSubscriber struct {
	UseCase   ucboundary.SeparatedUseCase
	Presenter Presenter
}

func (s *SeparatedSubscriber) UploadFiles(ctx context.Context, event *Event) error {
	files := event.toFiles()

	if err := s.UseCase.UploadFiles(ctx, files); err != nil {
		if err := s.Presenter.Error(ctx, err); err != nil {
			return err
		}
		return nil
	}

	for _, file := range files {
		if err := s.Presenter.Success(ctx, file); err != nil {
			return err
		}
	}
	return nil
}
