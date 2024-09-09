package usecase

import (
	"cases/internal/entity"
	"context"
	"fmt"
	"sync"
)

type Result[T any] struct {
	Value T
	Err   error
}

type ChanUseCase struct {
	Uploader Uploader
}

func (u *ChanUseCase) UploadFiles(ctx context.Context, files []*entity.File) chan *Result[*entity.File] {
	out := make(chan *Result[*entity.File])
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file *entity.File) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				fmt.Printf("Context is done for %s\n", file.Name)
				return
			default:
				if err := u.Uploader.Upload(ctx, file); err != nil {
					out <- &Result[*entity.File]{Value: nil, Err: err}
				} else {
					out <- &Result[*entity.File]{Value: file, Err: nil}
				}
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
