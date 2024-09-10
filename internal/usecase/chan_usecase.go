package usecase

import (
	"cases/internal/entity"
	"cases/internal/usecase/ucboundary"
	"context"
	"fmt"
	"sync"
)

var _ ucboundary.ChanUseCase = (*ChanUseCase)(nil)

type ChanUseCase struct {
	Uploader Uploader
}

func (u *ChanUseCase) UploadFiles(ctx context.Context, files []*entity.File) chan *ucboundary.Result[*entity.File] {
	out := make(chan *ucboundary.Result[*entity.File])
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
					out <- &ucboundary.Result[*entity.File]{Value: nil, Err: err}
				} else {
					out <- &ucboundary.Result[*entity.File]{Value: file, Err: nil}
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
