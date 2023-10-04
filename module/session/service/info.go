package sessionsvc

import (
	"context"
	sessionv1 "sdm/pb/sdm/session/v1"

	"github.com/gofiber/fiber/v2"
)

func (s *SessionService) Info(
	ctx context.Context,
	req *sessionv1.InfoRequest,
) (*sessionv1.InfoResponse, error) {

	current := s.manager.Get(req.Token)
	if current == nil {
		return nil, fiber.ErrUnauthorized
	}

	return &sessionv1.InfoResponse{
		Id:       current.Id,
		Scheme:   current.Driver,
		Database: current.Database,
	}, nil
}
