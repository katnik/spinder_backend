package girls

import (
	"encoding/json"
	"net/http"
	"strings"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func Read(w http.ResponseWriter, r *http.Request) {

	type searchOutput struct {
		Name   string `bson:"name"`
		UserID string `bson:"_id"`
		Likes  string `bson:"likes"`
		Hates  string `bson:"hates"`
	}

	userid := strings.Split(r.URL.Path, "/")[2]

	query := bson.M{
		"_id": userid,
	}

	s, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}

	c := s.DB("spinder").C("girls")

	results := []searchOutput{}
	c.Find(query).Sort("n").All(&results)

	if len(results) != 0 {
		out, _ := json.MarshalIndent(results, " ", "  ")
		w.Write(out)
	} else {
		out, _ := json.MarshalIndent("Not Found", " ", "  ")
		w.Write(out)
	}
	s.Close()
}
