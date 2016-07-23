package mongo

import (
	"labix.org/v2/mgo"
)

type empty struct{}

func GetCollectionSize(collection string) int {

	s, err := OpenSession()

	if err != nil {
		panic(err)
	}
	results := []empty{}

	c := s.DB("spinder").C(collection)
	c.Find(nil).All(&results)

	l := len(results)

	return l
}

func OpenSession() (*mgo.Session, error) {
	s, err := mgo.Dial("localhost:27017")
	if err != nil {
		return s, err
	}
	// Optional. Switch the session to a monotonic behavior.
	s.SetMode(mgo.Monotonic, true)
	return s, err
}

// CloseSession closes mongodb session
func CloseSession(session *mgo.Session) bool {
	session.Close()
	return true
}
