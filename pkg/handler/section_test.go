package handler

import (
	"bytes"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/romankuz19/avito-proj/pkg/service"
	service_mocks "github.com/romankuz19/avito-proj/pkg/service/mocks"
)

func TestHandler_createSection(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockSection, name string)

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *service_mocks.MockSection, name string) {
				r.EXPECT().Create(name).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"Section created"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			log.Println(test.name)
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockSection(c)
			test.mockBehavior(repo, test.name)

			services := &service.Service{Section: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/", handler.createSection)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/",
				bytes.NewBufferString(test.name))

			// Make Request
			r.ServeHTTP(w, req)
			log.Println(w.Body)
			// Asserts
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
