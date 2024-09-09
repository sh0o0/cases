package controller

import (
	"cases/internal/entity"
	"cases/internal/usecase"
	"context"
)

type ChanUseCase interface {
	UploadFiles(ctx context.Context, files []*entity.File) chan *usecase.Result[*entity.File]
}

type ChanSubscriber struct {
	UseCase   ChanUseCase
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
