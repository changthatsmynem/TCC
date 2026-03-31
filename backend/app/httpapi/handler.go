package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"tcc/backend/app/model"
	"tcc/backend/app/store"
)

type Handler struct {
	commentStore *store.CommentStore
}

func NewHandler(commentStore *store.CommentStore) *Handler {
	return &Handler{commentStore: commentStore}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/feed", h.handleFeed)
	mux.HandleFunc("/api/comments", h.handleAddComment)
	return mux
}

func (h *Handler) handleFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJSON(w, http.StatusOK, model.FeedResponse{
		PageTitle: "IT 08-1",
		Post: model.Post{
			Author:    "Change can",
			Avatar:    "C",
			CreatedAt: "16 October 2021 16:00",
			ImagePath: "/image1.png",
		},
		Comments: h.commentStore.List(),
	})
}

func (h *Handler) handleAddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.AddCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Message) == "" {
		http.Error(w, "message is required", http.StatusBadRequest)
		return
	}

	entry, err := h.commentStore.Add(req.Message)
	if err != nil {
		http.Error(w, "unable to save comment", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, entry)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("write JSON: %v", err)
	}
}
