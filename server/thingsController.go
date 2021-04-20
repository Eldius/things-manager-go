package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Eldius/things-manager-go/logger"
	"github.com/Eldius/things-manager-go/model"
	"github.com/sirupsen/logrus"
)

var (
	log = logger.Logger()
)

func HandleThingRoot(repo *model.Repository) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createThing(rw, r, repo)
		} else if r.Method == http.MethodGet {
			listThings(rw, r, repo)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func HandleThingWithId(repo *model.Repository) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("HandleThingWithId")
		if r.Method == http.MethodGet {
			getThing(rw, r, repo)
		} else if r.Method == http.MethodPut {
			updateThing(rw, r, repo)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func updateThing(rw http.ResponseWriter, r *http.Request, repo *model.Repository) {
	log.Info("Updating things")
	if id, err := getThingId(r); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(map[string]string{
			"reason": err.Error(),
		})
	} else {
		log.WithFields(logrus.Fields{
			"id": id,
		}).Info("fetching thing")

		var t *model.Thing
		_ = json.NewDecoder(r.Body).Decode(&t)
		t.ID = id
		repo.SaveThing(t)
		rw.WriteHeader(http.StatusNoContent)
	}
}

func getThing(rw http.ResponseWriter, r *http.Request, repo *model.Repository) {
	if id, err := getThingId(r); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(map[string]string{
			"reason": err.Error(),
		})
	} else {
		log.WithFields(logrus.Fields{
			"id": id,
		}).Info("fetching thing")
		t := repo.GetThing(id)
		if t != nil {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(rw).Encode(t)
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	}
}

func getThingId(r *http.Request) (int, error) {
	id, err := strconv.Atoi(strings.Trim(r.URL.Path, "/things/"))
	log.WithError(err).WithFields(logrus.Fields{
		"url": r.URL.Path,
		"id":  id,
	}).Info("GettingThing")
	return id, err
}

func createThing(rw http.ResponseWriter, r *http.Request, repo *model.Repository) {
	log.Info("Saving thing")
	defer r.Body.Close()
	var t *model.Thing
	_ = json.NewDecoder(r.Body).Decode(&t)
	_t := repo.GetThingByName(t.Name)
	if _t == nil {
		repo.SaveThing(t)
		rw.WriteHeader(http.StatusCreated)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(rw).Encode(map[string]string{
			"reason": "Name alread exists",
		})
	}
}

func listThings(rw http.ResponseWriter, r *http.Request, repo *model.Repository) {
	log.Info("Listing things")
	defer r.Body.Close()
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(repo.ListThings())
}
