package controller

import (
	"cases/internal/usecase"
	"cases/internal/usecase/ucboundary"
	"context"
)

type SeparatedSubscriber struct {
	UseCase   ucboundary.SeparatedUseCase
	Messenger usecase.Messenger
}

func (s *SeparatedSubscriber) UploadFiles(ctx context.Context, event *Event) error {
	files := event.toFiles()

	if err := s.UseCase.UploadFiles(ctx, files); err != nil {
		if err := s.Messenger.Send(ctx, "Failed"); err != nil {
			return err
		}
		return nil
	}

	if err := s.Messenger.Send(ctx, "Success"); err != nil {
		return err
	}

	return nil
}
