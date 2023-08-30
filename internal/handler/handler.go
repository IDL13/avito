package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/IDL13/avito/internal/CSV"
	"github.com/IDL13/avito/internal/timer"
)

func GettingData(r *http.Request, keyRequest string) (s string, err error) {
	param := r.Body
	var result map[string]string
	json.NewDecoder(param).Decode(&result)
	str := result[keyRequest]
	return str, nil

}

func Round(x float64) int {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return int(t + math.Copysign(1, x))
	}
	return int(t)
}

func (h *handler) StartServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server start"))
	data := []string{"UserID", "Segment", "Add/Remove", "Date-Time"}
	err := CSV.WriteInCSV(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var S Segment
		json.NewDecoder(param).Decode(&S)
		err := h.db.InserSegment(S.Name)
		if err != nil {
			w.Write([]byte("This segment is using"))
		} else {
			if S.Percent != 0 {
				c, err := h.db.Count()
				if err != nil {
					panic(err)
				}
				percent := Round((float64(c) * float64(S.Percent)) / float64(100))
				err = h.db.RandChoice(percent, S.Name)
				w.Write([]byte("Segment added to the database"))
			}
		}
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *handler) DeletingSegment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		jsonData, err := GettingData(r, "slug")
		if err != nil {
			panic(err)
		}
		err = h.db.DeleteSegment(jsonData)
		if err != nil {
			w.Write([]byte("This segment was not found"))
		}
		w.Write([]byte("Segment seccessfully deleted"))
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *handler) AddDelSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var d dependenciesData
		json.NewDecoder(param).Decode(&d)
		formatId, err := strconv.Atoi(d.UserId)
		if err != nil {
			panic(err)
		}
		if len(d.AddSegments) > 0 {
			err = h.db.InsertDependencies(formatId, d.AddSegments)
			if err != nil {
				panic(err)
			}
		}
		if len(d.DeleteSegments) > 0 {
			err = h.db.DeleteDependencies(formatId, d.DeleteSegments)
			if err != nil {
				panic(err)
			}
		}
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *handler) GettingActiveUserSegments(w http.ResponseWriter, r *http.Request) {
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
		info, err := h.db.SearchSegmentsForUser()
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

func (h *handler) TtlAddDelSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var ttl ttlStruct
		json.NewDecoder(param).Decode(&ttl)
		formatId, err := strconv.Atoi(ttl.DependenciesData.UserId)
		if err != nil {
			panic(err)
		}
		err = timer.CallAt(ttl.Start, h.db.InsertDependencies, formatId, ttl.DependenciesData.AddSegments)
		if err != nil {
			panic(err)
		}
		err = timer.CallAt(ttl.Stop, h.db.DeleteDependencies, formatId, ttl.DependenciesData.DeleteSegments)
		if err != nil {
			panic(err)
		}
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}

func (h *handler) Hishtory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		timeInterval, err := GettingData(r, "date")
		if err != nil {
			panic(err)
		}
		mapa := CSV.ReadInCSV(timeInterval)
		js, err := json.Marshal(mapa)
		if err != nil {
			panic(err)
		}
		w.Write(js)
	} else {
		w.Write([]byte("This url only handles POST requests"))
	}
}
