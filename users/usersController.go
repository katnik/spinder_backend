package users

import (
	"encoding/json"
	"net/http"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/katnik/spinder/utils/misc"
)

func Create(w http.ResponseWriter, r *http.Request) {

	type searchOutput struct {
		Username string `bson:"un"`
		UserID   string `bson:"id"`
		Password string `bson:"pw"`
	}
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	inputValues := r.Form
	name := inputValues.Get("username")
	password := inputValues.Get("password")

	s, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}

	c := s.DB("spinder").C("users")
	// Revisit userID generation. Maybe a singleton with a global state is not the
	// Best approach. 
	// The argument in GetInstance is irrelevant since the generator is a singleton
	id := misc.GetInstance(0).Generate()
	newPlayer := bson.M{
		"un": name,
		"id": id,
		"pw": password,
	}

	err = c.Insert(newPlayer)

	if err != nil {
		panic(err)
	}
	out, _ := json.MarshalIndent(newPlayer, " ", "  ")
	w.Write(out)
	s.Close()
}
