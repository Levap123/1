package funcs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// main struct
var All API

func ParseJson() error {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlRelations := "https://groupietrackers.herokuapp.com/api/relation"
	if err := parseInfo(urlArtists, &All.Artists); err != nil {
		return err
	}

	if err := parseInfo(urlRelations, &All.Relation); err != nil {
		return err
	}

	for i, v := range All.Relation.Index {
		All.Artists[i].Rel = v.DatesLocations
	}
	return nil
}

func parseInfo(url string, temp interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	text, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	return json.Unmarshal(text, &temp)
}
