package handler

import (
	"net/http"

	"github.com/IDL13/avito/internal/requests"
	"github.com/IDL13/avito/internal/response"
)

func New() Handler {
	h := &handler{}
	return h
}

type handler struct {
	db   requests.Db
	resp *response.HttpResponse
	err  *response.HttpError
}

type dependenciesData struct {
	UserId         int      `json:"id"`
	DeleteSegments []string `json:"del_segments"`
	AddSegments    []string `json:"add_segments"`
}

type ttlStruct struct {
	DependenciesData dependenciesData `json:"data"`
	Start            string           `json:"start"`
	Stop             string           `json:"stop"`
}

type createSegment struct {
	Name    string `json:"slug"`
	Percent int    `json:"percent"`
}

type deleteSegment struct {
	Name string `json:"slug"`
}

type user struct {
	Id   string `json:"id"`
	Name string `json:"slug"`
}

type history struct {
	Data string `json:"date"`
}

type Handler interface {
	StartServer(w http.ResponseWriter, r *http.Request)
	CreateSegment(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeletingSegment(w http.ResponseWriter, r *http.Request)
	DeletingUser(w http.ResponseWriter, r *http.Request)
	AddDelSegments(w http.ResponseWriter, r *http.Request)
	GettingActiveUserSegments(w http.ResponseWriter, r *http.Request)
	TtlAddDelSegments(w http.ResponseWriter, r *http.Request)
	Hishtory(w http.ResponseWriter, r *http.Request)
}
