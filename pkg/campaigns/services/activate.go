package services

import (
	"context"
	"errors"
	"log"

	"github.com/leal/pkg/campaigns/domain"
)

func (s *Service) Activate(ctx context.Context, name string, id int) error {
	err := s.Repository.Activate(ctx, name, id)
	if err != nil {
		if errors.Is(err, domain.ErrTimeout) {
			log.Println("time_out_error")
			srvErr := domain.Error{
				Code: domain.ErrCodeTimeout,
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
		if errors.Is(err, domain.ErrTimeout) {
			log.Println("time_out_error")
			srvErr := domain.Error{
				Code: domain.ErrCodeTimeout,
				Msg:  "error inactivating campaign: time out error",
			}
			return srvErr
		}
		log.Println(err.Error())
		return err
	}
	return nil
}
