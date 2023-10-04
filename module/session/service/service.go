package sessionsvc

import (
	"sdm/module/session"
	sessionv1 "sdm/pb/sdm/session/v1"
)

type SessionService struct {
	manager *session.SessionManager
	sessionv1.UnimplementedSessionServiceServer
}

func CreateSessionService(man *session.SessionManager) *SessionService {
	return &SessionService{manager: man}
}
