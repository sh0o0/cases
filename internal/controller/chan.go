package controller

import (
	"cases/internal/usecase/ucboundary"
	"context"
)

type ChanSubscriber struct {
	UseCase   ucboundary.ChanUseCase
	Presenter Presenter
}

func (s *ChanSubscriber) UploadFiles(ctx context.Context, event *Event) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	files := event.toFiles()

	out := s.UseCase.UploadFiles(ctx, files)
	for r := range out {
		if r.Err != nil {
			if err := s.Presenter.Error(ctx, r.Err); err != nil {
				return err
			}
		} else {
			if err := s.Presenter.Success(ctx, r.Value); err != nil {
				return err
			}
		}
	}

	return nil
}
