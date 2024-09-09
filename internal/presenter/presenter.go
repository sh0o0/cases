package presenter

import (
	"cases/internal/entity"
	"context"
	"fmt"
)

type Presenter struct{}

func (p *Presenter) Success(ctx context.Context, file *entity.File) error {
	fmt.Printf("Sent message for %s\n", file.Name)
	return nil
}

func (p *Presenter) Error(ctx context.Context, err error) error {
	fmt.Printf("Sent error massage: %v\n", err)
	return nil
}
