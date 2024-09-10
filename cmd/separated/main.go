package main

import (
	"cases/internal/controller"
	"cases/internal/infrastructure"
	"cases/internal/usecase"
	"context"
)

func main() {
	subscriber := &controller.SeparatedSubscriber{
		UseCase: &usecase.SeparatedUseCase{
			Uploader: &infrastructure.Uploader{},
		},
		Messenger: &infrastructure.Messenger{},
	}

	subscriber.UploadFiles(context.Background(), &controller.Event{})
}
