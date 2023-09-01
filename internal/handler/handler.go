package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/IDL13/avito/internal/CSV"
	"github.com/IDL13/avito/internal/requests"
	"github.com/IDL13/avito/internal/response"
	"github.com/IDL13/avito/internal/timer"
)

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

// @Summary		CreateSegment
// @Tags			segments
// @Description	Create segment in database
// @Accept			json
// @Produce		json
// @Param			input	body	createSegment  false  "Create Segment"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			/create_segment [post]
func (h *handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	h.resp = response.NewOk()
	h.err = response.NewErr()
	if r.Method == "POST" {
		param := r.Body
		var s createSegment
		json.NewDecoder(param).Decode(&s)
		h.db = requests.New()
		err := h.db.InserSegment(s.Name)
		if err != nil {
			w.Write(h.err.NewError(("This segment is using")))
		} else {
			if s.Percent != 0 {
				c, err := h.db.Count()
				if err != nil {
					panic(err)
				}
				percent := Round((float64(c) * float64(s.Percent)) / float64(100))
				err = h.db.RandChoice(percent, s.Name)
				w.Write(h.resp.NewResponse("Segment added to the database"))
			}
		}
	} else {
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}

// @Summary		DeletingSegment
// @Tags			segments
// @Description	Delet segment in database
// @Accept			json
// @Produce		json
// @Param			input	body	deleteSegment  false  "Delete Segment"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			/deleting_segment [post]
func (h *handler) DeletingSegment(w http.ResponseWriter, r *http.Request) {
	h.resp = response.NewOk()
	h.err = response.NewErr()
	if r.Method == "POST" {
		param := r.Body
		var s deleteSegment
		json.NewDecoder(param).Decode(&s)
		h.db = requests.New()
		err := h.db.DeleteSegment(s.Name)
		if err != nil {
			w.Write(h.err.NewError("This segment was not found"))
		}
		w.Write(h.resp.NewResponse("Segment seccessfully deleted"))
	} else {
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}

// @Summary		CreateUser
// @Tags			users
// @Description	Create user in database
// @Accept			json
// @Produce		json
// @Param			input	body	user  false  "Create User"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			/create_user [post]
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.resp = response.NewOk()
	h.err = response.NewErr()
	if r.Method == "POST" {
		param := r.Body
		var user user
		json.NewDecoder(param).Decode(&user)
		h.db = requests.New()
		err := h.db.InsertUser(user.Name)
		if err != nil {
			w.Write(h.err.NewError(("This userID is using")))
		} else {
			w.Write(h.resp.NewResponse("User added"))
		}
	} else {
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}

// @Summary		DeletingUser
// @Tags			users
// @Description	Delete user in database
// @Accept			json
// @Produce		json
// @Param			input	body	user  false  "Delete User"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			/deleting_user [post]
func (h *handler) DeletingUser(w http.ResponseWriter, r *http.Request) {
	h.resp = response.NewOk()
	h.err = response.NewErr()
	if r.Method == "POST" {
		param := r.Body
		var user user
		json.NewDecoder(param).Decode(&user)
		h.db = requests.New()
		err := h.db.DeleteUser(user.Name)
		if err != nil {
			w.Write(h.err.NewError(("This userID is using")))
		} else {
			w.Write(h.resp.NewResponse("User delete"))
		}
	} else {
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}

// @Summary		AddDelSegments
// @Tags			dependencies
// @Description	Addition segments for user in database
// @Accept			json
// @Produce		json
// @Param			input	body	dependenciesData  false  "Adding Segment"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			/adding_user_to_segment [post]
func (h *handler) AddDelSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var d dependenciesData
		json.NewDecoder(param).Decode(&d)
		h.db = requests.New()
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
		w.Write(h.resp.NewResponse("Operation seccessful"))
	} else {
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}

// @Summary		GettingActiveUserSegments
// @Tags			dependencies
// @Description	User segment check
// @Accept			json
// @Produce		json
// @Param			input	body	user  false  "Check Segment"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			/getting_active_user_segments [post]
func (h *handler) GettingActiveUserSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var user user
		json.NewDecoder(param).Decode(&user)
		h.db = requests.New()
		jsonInt, err := strconv.Atoi(user.Id)
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
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}

// @Summary		TtlAddDelSegments
// @Tags			ttl
// @Description	Ttl adding or remove
// @Accept			json
// @Produce		json
// @Param			input	body	ttlStruct  false  "ttl hadler"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			/ttl_adding_user_to_segment [post]
func (h *handler) TtlAddDelSegments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var ttl ttlStruct
		json.NewDecoder(param).Decode(&ttl)
		h.db = requests.New()
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
		w.Write(h.resp.NewResponse("Operation seccessful"))
	} else {
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}

// @Summary		Hishtory
// @Tags			history
// @Description	Check history
// @Accept			json
// @Produce		json
// @Param			input	body	history  false  "ttl hadler"
// @Success		200		{object}	response.HttpResponse
// @Failure		400		{object}	response.HttpError
// @Failure		404		{object}	response.HttpError
// @Failure		500		{object}	response.HttpError
// @Router			//history [post]
func (h *handler) Hishtory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		param := r.Body
		var data history
		json.NewDecoder(param).Decode(&data)
		h.db = requests.New()
		mapa := CSV.ReadInCSV(data.Data)
		js, err := json.Marshal(mapa)
		if err != nil {
			panic(err)
		}
		w.Write(js)
	} else {
		w.Write(h.err.NewError("This url only handles POST requests"))
	}
}
