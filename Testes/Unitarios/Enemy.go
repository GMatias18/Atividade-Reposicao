// file: /test/unit/enemy_test.go
package unit

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAddEnemy(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/enemy", func(w http.ResponseWriter, r *http.Request) {
        AddEnemy(w, r) // Função deve ser exportada ou replicada aqui para teste
    })

    body := `{"nickname":"testEnemy"}`
    req := httptest.NewRequest(http.MethodPost, "/enemy", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    rec := httptest.NewRecorder()

    mux.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    var e Enemy
    err := json.Unmarshal(rec.Body.Bytes(), &e)
    assert.NoError(t, err)
    assert.Equal(t, "testEnemy", e.Nickname)
    assert.Greater(t, e.Life, 0)
    assert.Greater(t, e.Attack, 0)
}
