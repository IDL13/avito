package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_serv "github.com/IDL13/avito/mock"
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
		inputBody: `{"slug":"testAVITO", "percent":10}`,
		inputService: createSegment{
			Name:    "testAVITO",
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
			inputBody: `{"slug":"testAVITO2", "percent":0}`,
			inputService: createSegment{
				Name:    "testAVITO2",
				Percent: 0,
			},
			expectedCode: 200,
			f: func(s *mock_serv.MockDb, u createSegment) {
				s.EXPECT().InserSegment(u.Name).Return(nil).AnyTimes()
			},
			expectedBody: `{"message":"empty percent"}`,
		},
		{
			name:      "400",
			inputBody: `{"slu":"test", "percent":0}`,
			inputService: createSegment{
				Name:    "",
				Percent: 0,
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
		inputBody: `{"slug":"testAVITO2"}`,
		inputService: deleteSegment{
			Name: "testAVITO2",
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
			expectedBody: `{"message":"This segment was not found or there is dublicate in database"}`,
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
		inputBody: `{"id":"3"}`,
		inputService: user{
			Id: "3",
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u user) {
			s.EXPECT().SearchSegmentsForUser().Return(info, nil).AnyTimes()
		},
		expectedBody: `{"3":["AVITO_DISCOUNT_50","AVITO_DISCOUNT_30","VOICE_MESAGE","AVITO_DISCOUNT_30","AVITO_DISCOUNT_30","VOICE_MESAGE","AVITO_DISCOUNT_30","AVITO_DISCOUNT_30","VOICE_MESAGE","AVITO_DISCOUNT_30"]}`,
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
		inputBody: `{"id":2,"add_segments":["AVITO_MESSAGE", "AVITO"]}`,
		inputService: dependenciesData{
			UserId:      2,
			AddSegments: []string{"AVITO_MESSAGE", "AVITO"},
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u dependenciesData) {
			s.EXPECT().InsertDependencies(u.UserId, u.AddSegments).Return(nil).AnyTimes()
		},
		expectedBody: `{"message":"Operation seccessful"}{"message":"Operation seccessful"}`,
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
		inputBody: `{"id":1,"del_segments":["AVITO_MESSAGE"], "add_segments":[]}`,
		inputService: dependenciesData{
			UserId:         1,
			DeleteSegments: []string{"AVITO_MESSAGE"},
			AddSegments:    []string{},
		},
		expectedCode: 200,
		f: func(s *mock_serv.MockDb, u dependenciesData) {
			s.EXPECT().DeleteDependencies(u.UserId, u.DeleteSegments).Return(nil).AnyTimes()
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
