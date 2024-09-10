package infrastructure

import (
	"cases/internal/usecase"
	"context"
	"fmt"
)

var _ usecase.Messenger = (*Messenger)(nil)

type Messenger struct{}

func (m *Messenger) Send(ctx context.Context, msg string) error {
	fmt.Println(msg)
	return nil
}
