package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullAPIFlow(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			AddPlayer(w, r)
		} else if r.Method == http.MethodGet {
			LoadPlayers(w, r)
		}
	})
	mux.HandleFunc("/enemy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			AddEnemy(w, r)
		} else if r.Method == http.MethodGet {
			LoadEnemies(w, r)
		}
	})
	mux.HandleFunc("/battle", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateBattle(w, r)
		} else if r.Method == http.MethodGet {
			LoadBattles(w, r)
		}
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	playerData := `{"nickname":"Hero","life":100,"attack":10}`
	enemyData := `{"nickname":"Villain"}`
	battleData := `{"player":"Hero","enemy":"Villain"}`

	resp, err := http.Post(server.URL+"/player", "application/json", bytes.NewBufferString(playerData))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()

	resp, err = http.Post(server.URL+"/enemy", "application/json", bytes.NewBufferString(enemyData))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()

	resp, err = http.Post(server.URL+"/battle", "application/json", bytes.NewBufferString(battleData))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}
