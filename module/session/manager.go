package session

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

type SessionManager struct {
	sessions map[string]*Session
	mu       sync.Mutex
}

func CreateSessionManager() *SessionManager {
	return &SessionManager{
		sessions: map[string]*Session{},
		mu:       sync.Mutex{},
	}
}

func (s *SessionManager) Print() {
	for _, s := range s.sessions {
		log.Info(s.Id)
	}
}
func (s *SessionManager) Get(id string) *Session {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.sessions[id]
}

func (s *SessionManager) Add(session *Session) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if session.Id == "" {
		oid := rand.Int()
		session.Id = fmt.Sprint(oid)
	}

	s.sessions[session.Id] = session

	return session.Id
}

func (s *SessionManager) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	con, ok := s.sessions[id]
	if ok {
		con.Db.Close()
		delete(s.sessions, id)
	}
}

func (s *SessionManager) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key, con := range s.sessions {
		fmt.Printf("deleting session %s", key)
		con.Db.Close()
		delete(s.sessions, key)
	}
}
