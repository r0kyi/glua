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

	sess struct {
		key   string
		value string
	}

	store   *cookie.Store
	session sessions.Session
	context *gin.Context
}

func (s *Session) newStore() {
	var keyPairs [][]byte
	for _, key := range s.Keys {
		keyPairs = append(keyPairs, core.S2B(key))
	}
	store := cookie.NewStore(keyPairs...)
	s.store = &store
}

func (s *Session) default_() {
	s.session = sessions.Default(s.context)
}

func (s *Session) get() {
	value := s.session.Get(s.sess.key)
	if value == nil {
		s.sess.value = ""
	} else {
		s.sess.value = value.(string)
	}
}

func (s *Session) set() {
	s.session.Set(s.sess.key, s.sess.value)
}

func (s *Session) delete() {
	s.session.Delete(s.sess.key)
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
