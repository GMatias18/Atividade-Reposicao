// file: /test/unit/player_test.go
package unit

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAddPlayer(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
        AddPlayer(w, r) // Esta função precisa ser acessível
    })

    body := `{"nickname":"testPlayer","life":50,"attack":5}`
    req := httptest.NewRequest(http.MethodPost, "/player", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    rec := httptest.NewRecorder()

    mux.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    var p PlayerRequest
    err := json.Unmarshal(rec.Body.Bytes(), &p)
    assert.NoError(t, err)
    assert.Equal(t, "testPlayer", p.Nickname)
}

