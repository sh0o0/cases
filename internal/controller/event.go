package controller

import "cases/internal/entity"

type Event struct{}

func (s *Event) toFiles() []*entity.File {
	return []*entity.File{
		{
			Name: "file1",
		},
		{
			Name: "file2",
		},
		{
			Name: "file3",
		},
	}
}
