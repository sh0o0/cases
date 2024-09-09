package main

import (
	"cases/internal/controller"
	"cases/internal/infrastructure"
	"cases/internal/presenter"
	"cases/internal/usecase"
	"context"
)

func main() {
	subscriber := &controller.IterSubscriber{
		UseCase: &usecase.IterUseCase{
			Uploader: &infrastructure.Uploader{},
		},
		Presenter: &presenter.Presenter{},
	}

	subscriber.UploadFiles(context.Background(), &controller.Event{})
}
