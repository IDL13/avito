package handler

import (
	"net/http"

	"github.com/IDL13/avito/internal/requests"
)

func New() Handler {
	h := &handler{}
	return h
}

type handler struct {
	db requests.Db
}

type dependenciesData struct {
	UserId         string   `json:"id"`
	DeleteSegments []string `json:"del_segments"`
	AddSegments    []string `json:"add_segments"`
}

type ttlStruct struct {
	DependenciesData dependenciesData `json:"data"`
	Start            string           `json:"start"`
	Stop             string           `json:"stop"`
}

type Segment struct {
	Name    string `json:"slug"`
	Percent int    `json:"percent"`
}

type Handler interface {
	StartServer(w http.ResponseWriter, r *http.Request)
	CreateSegment(w http.ResponseWriter, r *http.Request)
	DeletingSegment(w http.ResponseWriter, r *http.Request)
	AddDelSegments(w http.ResponseWriter, r *http.Request)
	GettingActiveUserSegments(w http.ResponseWriter, r *http.Request)
	TtlAddDelSegments(w http.ResponseWriter, r *http.Request)
	Hishtory(w http.ResponseWriter, r *http.Request)
}
