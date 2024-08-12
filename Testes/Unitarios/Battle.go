// file: /test/unit/battle_test.go
package unit

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateBattle(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/battle", func(w http.ResponseWriter, r *http.Request) {
        CreateBattle(w, r) // Assegure que os dados e métodos estejam acessíveis
    })

    body := `{"enemy":"enemy1", "player":"player1"}`
    req := httptest.NewRequest(http.MethodPost, "/battle", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    rec := httptest.NewRecorder()

    mux.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    var b Battle
    err := json.Unmarshal(rec.Body.Bytes(), &b)
    assert.NoError(t, err)
    assert.NotEmpty(t, b.ID)
    assert.GreaterOrEqual(t, b.DiceThrown, 1)
    assert.LessOrEqual(t, b.DiceThrown, 6)
}
