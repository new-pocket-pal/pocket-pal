package service

import "pocket-pal/src/facades"

// Service is a struct that represent the service
type Service struct {
	*facades.Facades
}

func NewService(f *facades.Facades) *Service {
	return &Service{
		f,
	}
}
