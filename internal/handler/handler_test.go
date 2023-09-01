package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	mock_serv "github.com/IDL13/avito/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		inputService string
		expectedCode int
		f            func(s *mock_serv.MockDb, u string)
		expectedBody string
	}{{
		name:         "OK",
		inputBody:    `{"slug":"test"}`,
		inputService: "test",
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u string) {
			s.EXPECT().InsertUser(u).Return(nil).AnyTimes()
		},
		expectedBody: `{"message":"User added"}`,
	},
		{
			name:         "400",
			inputBody:    `{"slu":"test"}`,
			inputService: "test",
			expectedCode: 200,
			f: func(s *mock_serv.MockDb, u string) {
				s.EXPECT().InsertUser(u).Return(nil).AnyTimes()
			},
			expectedBody: `{"message":"This userID is using"}`,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockDb(ctrl)
			test.f(srvc, test.inputService)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/create_user", service.CreateUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create_user", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		inputService string
		expectedCode int
		f            func(s *mock_serv.MockDb, u string)
		expectedBody string
	}{{
		name:         "OK",
		inputBody:    `{"slug":"test"}`,
		inputService: "test",
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u string) {
			s.EXPECT().DeleteUser(u).Return(nil).AnyTimes()
		},
		expectedBody: `{"message":"User delete"}`,
	},
		{
			name:         "400",
			inputBody:    `{"slu":"test"}`,
			inputService: "test",
			expectedCode: 200,
			f: func(s *mock_serv.MockDb, u string) {
				s.EXPECT().DeleteUser(u).Return(nil).AnyTimes()
			},
			expectedBody: `{"message":"This userID is using"}`,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockDb(ctrl)
			test.f(srvc, test.inputService)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/deleting_user", service.DeletingUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/deleting_user", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestInserSegment(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		inputService createSegment
		expectedCode int
		f            func(s *mock_serv.MockDb, u createSegment)
		expectedBody string
	}{{
		name:      "OK",
		inputBody: `{"slug":"test", "percent":10}`,
		inputService: createSegment{
			Name:    "test",
			Percent: 10,
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u createSegment) {
			s.EXPECT().InserSegment(u.Name).Return(nil).AnyTimes()
		},
		expectedBody: `{"message":"Segment added to the database"}`,
	},
		{
			name:      "Simple create",
			inputBody: `{"slug":"test", "percent":}`,
			inputService: createSegment{
				Name:    "test",
				Percent: 10,
			},
			expectedCode: 200,
			f: func(s *mock_serv.MockDb, u createSegment) {
				s.EXPECT().InserSegment(u.Name).Return(nil).AnyTimes()
			},
			expectedBody: `{"message":"Segment added to the database"}`,
		},
		{
			name:      "400",
			inputBody: `{"slu":"test", "percent":10}`,
			inputService: createSegment{
				Name:    "test",
				Percent: 10,
			},
			expectedCode: 200,
			f: func(s *mock_serv.MockDb, u createSegment) {
				s.EXPECT().InserSegment(u.Name).Return(nil).AnyTimes()
			},
			expectedBody: `{"message":"This segment is using"}`,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockDb(ctrl)
			test.f(srvc, test.inputService)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/create_segment", service.CreateSegment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create_segment", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestDeleteSegment(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		inputService deleteSegment
		expectedCode int
		f            func(s *mock_serv.MockDb, u deleteSegment)
		expectedBody string
	}{{
		name:      "OK",
		inputBody: `{"slug":"test"}`,
		inputService: deleteSegment{
			Name: "test",
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u deleteSegment) {
			s.EXPECT().InserSegment(u.Name).Return(nil).AnyTimes()
		},
		expectedBody: `{"message":"Segment seccessfully deleted"}`,
	},
		{
			name:      "400",
			inputBody: `{"slu":"test"}`,
			inputService: deleteSegment{
				Name: "test",
			},
			expectedCode: 200,
			f: func(s *mock_serv.MockDb, u deleteSegment) {
				s.EXPECT().InserSegment(u.Name).Return(nil).AnyTimes()
			},
			expectedBody: `{"message":"This segment was not found"}`,
		}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockDb(ctrl)
			test.f(srvc, test.inputService)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/deleting_segment", service.DeletingSegment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/deleting_segment", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestSearchSegmentsForUser(t *testing.T) {
	info := make(map[int][]string)
	tests := []struct {
		name         string
		inputBody    string
		inputService user
		expectedCode int
		f            func(s *mock_serv.MockDb, u user)
		expectedBody string
	}{{
		name:      "OK",
		inputBody: ``,
		inputService: user{
			Name: "test",
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u user) {
			s.EXPECT().SearchSegmentsForUser().Return(info, nil).AnyTimes()
		},
		expectedBody: `{"message":"Segment seccessfully deleted"}`,
	},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockDb(ctrl)
			test.f(srvc, test.inputService)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/getting_active_user_segments", service.GettingActiveUserSegments)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/getting_active_user_segments", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestInsertDependencies(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		inputService dependenciesData
		expectedCode int
		f            func(s *mock_serv.MockDb, u dependenciesData)
		expectedBody string
	}{{
		name:      "OK",
		inputBody: `{"id":"1","del_segments":[], "add_segments":["AVITO_MESSAGE"]}`,
		inputService: dependenciesData{
			UserId:         "1",
			DeleteSegments: []string{},
			AddSegments:    []string{"AVITO_MESSAGE"},
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u dependenciesData) {
			formatId, _ := strconv.Atoi(u.UserId)
			s.EXPECT().InsertDependencies(formatId, u.AddSegments).Return(nil).AnyTimes()
		},
		expectedBody: `{"message":"Operation seccessful"}`,
	},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockDb(ctrl)
			test.f(srvc, test.inputService)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/adding_user_to_segment", service.AddDelSegments)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/adding_user_to_segment", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestDeleteDependencies(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		inputService dependenciesData
		expectedCode int
		f            func(s *mock_serv.MockDb, u dependenciesData)
		expectedBody string
	}{{
		name:      "OK",
		inputBody: `{"id":"1","del_segments":["AVITO_MESSAGE"], "add_segments":[]}`,
		inputService: dependenciesData{
			UserId:         "1",
			DeleteSegments: []string{"AVITO_MESSAGE"},
			AddSegments:    []string{},
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u dependenciesData) {
			formatId, _ := strconv.Atoi(u.UserId)
			s.EXPECT().DeleteDependencies(formatId, u.DeleteSegments).Return(nil).AnyTimes()
		},
		expectedBody: `{"message":"Operation seccessful"}`,
	},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockDb(ctrl)
			test.f(srvc, test.inputService)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/adding_user_to_segment", service.AddDelSegments)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/adding_user_to_segment", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
