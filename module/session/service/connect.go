package sessionsvc

import (
	"context"
	"net/url"
	"sdm/module/session"
	sessionv1 "sdm/pb/sdm/session/v1"

	"github.com/jmoiron/sqlx"
)

func (s *SessionService) Connect(
	ctx context.Context,
	req *sessionv1.ConnectRequest,
) (*sessionv1.ConnectResponse, error) {
	uri := url.URL{
		Scheme:   req.Scheme,
		Host:     req.Host,
		Path:     req.Database,
		User:     url.UserPassword(req.Username, req.Password),
		RawQuery: "sslmode=disable",
	}

	conStr := uri.String()

	db, err := sqlx.Connect(uri.Scheme, conStr)
	if err != nil {
		return nil, err
	}

	session := session.Session{
		Db:       db,
		Driver:   req.Scheme,
		Database: req.Database,
	}

	id := s.manager.Add(&session)
	s.manager.Print()

	return &sessionv1.ConnectResponse{
		Token: id,
	}, nil
}
