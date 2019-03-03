package server

import "context"

//interface
type AddService interface {
	Sum(_ context.Context, a, b int) (v int)
	Concat(_ context.Context, a, b string) (v string)
}

//service struct
type addService struct{}

//returns a implementation
func New() AddService {
	return addService{}
}

func (addService) Sum(_ context.Context, a, b int) (v int) {
	return a + b
}

func (addService) Concat(_ context.Context, a, b string) (v string) {
	return a + b
}
