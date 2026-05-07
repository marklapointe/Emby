package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
)

type CollectionHandler struct {
	itemRepo *repository.ItemRepository
}

func NewCollectionHandler(itemRepo *repository.ItemRepository) *CollectionHandler {
	return &CollectionHandler{itemRepo: itemRepo}
}

func (h *CollectionHandler) CreateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *CollectionHandler) AddToCollection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}