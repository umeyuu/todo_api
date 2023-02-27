package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/umeyuu/todo_api/handler"
	"github.com/umeyuu/todo_api/store"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析のエラーを回避するため明示的に戻り値を捨てている
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	// store.Tasksはグローバル変数
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux
}
