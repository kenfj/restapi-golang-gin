package app

// https://qiita.com/268iop/items/dc6e73138ee97dff2f73

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetRouter(t *testing.T) {

	router := SetRouter()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "{\"msg\":\"pong\"}")
}
