package controller

import (
	"cases/internal/entity"
	"context"
	"iter"
)

type IterUseCase interface {
	UploadFiles(ctx context.Context, files []*entity.File) iter.Seq2[*entity.File, error]
}

type IterSubscriber struct {
	UseCase   IterUseCase
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
