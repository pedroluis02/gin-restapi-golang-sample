package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedroluis02/gin-restapi-golang-sample/src/api"
	"github.com/stretchr/testify/assert"
)

func TestGeneralRoutes(t *testing.T) {
	server := api.NewServer(false)

	t.Run("Ping", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &res)

		assert.Nil(t, err)
		assert.Equal(t, "pong", res["message"])
	})
}
