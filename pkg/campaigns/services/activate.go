package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/helpers"
)

func (s *Service) Activate(ctx context.Context, name string, id int) error {
	err := s.Repository.Activate(ctx, name, id)
	if err != nil {
		if errors.Is(err, helpers.ErrTimeout) {
			log.Println("time_out_error")
			srvErr := helpers.Error{
				Code: helpers.ErrCodeTimeout,
				Msg:  "error activating campaign: time out error",
			}
			return srvErr
		}
		log.Println(err.Error())
		return err
	}
	return nil
}
func (s *Service) Inactivate(ctx context.Context, name string, id int) error {
	err := s.Repository.Inactivate(ctx, name, id)
	if err != nil {
		if errors.Is(err, helpers.ErrTimeout) {
			log.Println("time_out_error")
			srvErr := helpers.Error{
				Code: helpers.ErrCodeTimeout,
				Msg:  "error inactivating campaign: time out error",
			}
			return srvErr
		}
		log.Println(err.Error())
		return err
	}
	return nil
}

func (r *Service) SetActivate(ctx context.Context, name string, id int) error {
	return nil
}

func (r *Service) SetInactivate(ctx context.Context, name string, id int) error {
	return nil
}
