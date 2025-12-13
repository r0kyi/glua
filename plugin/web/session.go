package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/r0kyi/glua/core"
)

type Session struct {
	Name string   `lua:"name"`
	Keys []string `lua:"keys"`

	store   *cookie.Store
	session sessions.Session
}

func (s *Session) newStore() {
	var keyPairs [][]byte
	for _, key := range s.Keys {
		keyPairs = append(keyPairs, core.S2B(key))
	}
	store := cookie.NewStore(keyPairs...)
	s.store = &store
}

func (s *Session) default_(context *gin.Context) {
	s.session = sessions.Default(context)
}

func (s *Session) get(key string) string {
	value := s.session.Get(key)
	if value == nil {
		return ""
	}

	return value.(string)
}

func (s *Session) set(key string, value string) {
	s.session.Set(key, value)
}

func (s *Session) delete(key string) {
	s.session.Delete(key)
}

func (s *Session) clear() {
	s.session.Clear()
}

func (s *Session) save() error {
	err := s.session.Save()
	if err != nil {
		return err
	}

	return nil
}
