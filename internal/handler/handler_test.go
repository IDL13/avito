package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
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
func TestInserSegment(t *testing.T)          {}
func TestDeleteSegment(t *testing.T)         {}
func TestSearchSegmentsForUser(t *testing.T) {}
func TestInsertDependencies(t *testing.T)    {}
func TestDeleteDependencies(t *testing.T)    {}
func TestCount(t *testing.T)                 {}
func TestRandChoice(t *testing.T)            {}
