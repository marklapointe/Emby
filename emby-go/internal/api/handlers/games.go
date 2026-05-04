package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

// GamesHandler handles games-related API endpoints.
type GamesHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewGamesHandler creates a new games handler.
func NewGamesHandler(cfg *config.Config, repo *repository.ItemRepository) *GamesHandler {
	return &GamesHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetGames handles GET /Videos/Games
func (h *GamesHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	isFavorite := r.URL.Query().Get("IsFavorite")
	_ = isFavorite

	games, err := h.repo.GetGames(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

// GetGame handles GET /Videos/Games/{id}
func (h *GamesHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "id")

	game, err := h.repo.GetGame(gameId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

// GetGameGenres handles GET /Videos/Games/Genres
func (h *GamesHandler) GetGameGenres(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	genres, err := h.repo.GetGameGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// GetGameStudios handles GET /Videos/Games/Studios
func (h *GamesHandler) GetGameStudios(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	studios, err := h.repo.GetGameStudios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
}

// GetGameCompanies handles GET /Videos/Games/Companies
func (h *GamesHandler) GetGameCompanies(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	companies, err := h.repo.GetGameCompanies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}

// GetGameImage handles GET /Videos/Games/{id}/Images/{type}
func (h *GamesHandler) GetGameImage(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "id")
	imageType := chi.URLParam(r, "type")

	_ = gameId
	_ = imageType

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "game image"})
}

// GetGameBackdropImage handles GET /Videos/Games/{id}/BackdropImage
func (h *GamesHandler) GetGameBackdropImage(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "id")

	_ = gameId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "game backdrop image"})
}

// GetGameLogoImage handles GET /Videos/Games/{id}/LogoImage
func (h *GamesHandler) GetGameLogoImage(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "id")

	_ = gameId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "game logo image"})
}

// GetGameThumbImage handles GET /Videos/Games/{id}/ThumbImage
func (h *GamesHandler) GetGameThumbImage(w http.ResponseWriter, r *http.Request) {
	gameId := chi.URLParam(r, "id")

	_ = gameId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "game thumb image"})
}
