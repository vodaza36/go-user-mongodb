package mongo

import (
	"gopkg.in/mgo.v2"
)

// Session struct
type Session struct {
	session *mgo.Session
}

// NewSession instance
func NewSession(url string) (*Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

// Copy a session object
func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

// GetCollection returns the given collection
func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

// Close the session
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

// DropDatabase the mongo db
func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}
