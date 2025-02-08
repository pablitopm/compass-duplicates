package writer

import (
	"fmt"
	"main/model"
)

type StdOut struct {
}

func NewStdOut() *StdOut {
	return &StdOut{}
}

func (s StdOut) Write(result []model.CompareResult) error {
	for _, r := range result {
		fmt.Printf("Original User ID: %d, Compared User ID: %d, Accuracy: %s\n", r.ContactIDSource, r.ContactIDMatch, r.Accuracy)
	}
	return nil
}
