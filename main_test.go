package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MakMemEng/go-api-middle-road/handlers"
	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)
	return r
}

func TestRouterInitialization(t *testing.T) {
	router := setupRouter()
	if router == nil {
		t.Fatal("ルータの初期化に失敗しました。")
	}
}

func TestRouteRegistration(t *testing.T) {
	router := setupRouter()
	testCases := []struct {
		path   string
		method string
	}{
		{"/hello", http.MethodGet},
		{"/article", http.MethodGet},
		{"/article/list", http.MethodGet},
		{"/article/1", http.MethodGet},
		{"/article/nice", http.MethodPost},
		{"/comment", http.MethodPost},
	}

	for _, tc := range testCases {
		req, _ := http.NewRequest(tc.method, tc.path, nil)
		match := &mux.RouteMatch{}
		if !router.Match(req, match) {
			t.Errorf("ルート %s %s が見つかりません", tc.method, tc.path)
		}
	}

}
func TestServerStartup(t *testing.T) {
	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/hello")
	if err != nil {
		t.Fatalf("サーバーへのリクエストに失敗しました: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("予期しないステータスコード: got %v want %v", resp.StatusCode, http.StatusOK)
	}
}
