package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	request "github.com/IDL13/avito/internal/requests"
	"github.com/IDL13/avito/internal/timer"
)

type Handler struct{}

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

func GettingData(r *http.Request, keyRequest string) (s string, err error) {
	param := r.Body
	var result map[string]string
	json.NewDecoder(param).Decode(&result)
	str := result[keyRequest]
	return str, nil

}

func (h *Handler) StartServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server start"))
}

func (h *Handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		jsonData, err := GettingData(r, "slug")
		if err != nil {
			panic(err)
		}
		err = request.InserSegment(jsonData)
		if err != nil {
			w.Write([]byte("This segment is using"))
		}
		w.Write([]byte("Segment added to the database"))
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *Handler) DeletingSegment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		jsonData, err := GettingData(r, "slug")
		if err != nil {
			panic(err)
		}
		err = request.DeleteSegment(jsonData)
		if err != nil {
			w.Write([]byte("This segment was not found"))
		}
		w.Write([]byte("Segment seccessfully deleted"))
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *Handler) AddDelSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var d dependenciesData
		json.NewDecoder(param).Decode(&d)
		formatId, err := strconv.Atoi(d.UserId)
		if err != nil {
			panic(err)
		}
		if len(d.AddSegments) > 0 {
			err = request.InsertDependencies(formatId, d.AddSegments)
			if err != nil {
				panic(err)
			}
		}
		if len(d.DeleteSegments) > 0 {
			err = request.DeleteDependencies(formatId, d.DeleteSegments)
			if err != nil {
				panic(err)
			}
		}
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *Handler) GettingActiveUserSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		jsonData, err := GettingData(r, "id")
		if err != nil {
			panic(err)
		}
		jsonInt, err := strconv.Atoi(jsonData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "data conversion error:%v", err)
			os.Exit(1)
		}
		info, err := request.SearchSegmentsForUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Segment serch error:%v", err)
			os.Exit(1)
		}
		for key, value := range info {
			if key == jsonInt {
				ans := make(map[int][]string)
				ans[key] = value
				js, err := json.Marshal(ans)
				if err != nil {
					fmt.Fprintf(os.Stderr, "json marshaling error:%v", err)
					os.Exit(1)
				}
				w.Write(js)
			}
		}
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *Handler) TtlAddDelSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var ttl ttlStruct
		json.NewDecoder(param).Decode(&ttl)
		formatId, err := strconv.Atoi(ttl.DependenciesData.UserId)
		if err != nil {
			panic(err)
		}
		err = timer.CallAt(ttl.Start, request.InsertDependencies, formatId, ttl.DependenciesData.AddSegments)
		if err != nil {
			panic(err)
		}
		err = timer.CallAt(ttl.Stop, request.DeleteDependencies, formatId, ttl.DependenciesData.AddSegments)
		if err != nil {
			panic(err)
		}
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}
