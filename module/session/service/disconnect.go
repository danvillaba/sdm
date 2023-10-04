package sessionsvc

import (
	"context"
	sessionv1 "sdm/pb/sdm/session/v1"
)

func (s *SessionService) Disconnect(
	ctx context.Context,
	req *sessionv1.DisconnectRequest,
) (*sessionv1.DisconnectResponse, error) {
	s.manager.Remove(req.Token)

	return &sessionv1.DisconnectResponse{}, nil
}
