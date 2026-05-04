package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

// ChannelHandler handles channel-related API endpoints.
type ChannelHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewChannelHandler creates a new channel handler.
func NewChannelHandler(cfg *config.Config, repo *repository.ItemRepository) *ChannelHandler {
	return &ChannelHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetChannels handles GET /Channels
func (h *ChannelHandler) GetChannels(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	isFavorite := r.URL.Query().Get("IsFavorite")
	_ = isFavorite

	channels, err := h.repo.GetChannels(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}

// GetChannel handles GET /Channels/{id}
func (h *ChannelHandler) GetChannel(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")

	channel, err := h.repo.GetChannel(channelId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channel)
}

// GetChannelFolders handles GET /Channels/{id}/Folders
func (h *ChannelHandler) GetChannelFolders(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")

	folders, err := h.repo.GetChannelFolders(channelId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

// GetChannelItems handles GET /Channels/{id}/Items
func (h *ChannelHandler) GetChannelItems(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")
	userId := r.URL.Query().Get("UserId")

	items, err := h.repo.GetChannelItems(channelId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetChannelImage handles GET /Channels/{id}/Images/{type}
func (h *ChannelHandler) GetChannelImage(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")
	imageType := chi.URLParam(r, "type")

	_ = channelId
	_ = imageType

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "image"})
}

// GetChannelLogoImage handles GET /Channels/{id}/LogoImage
func (h *ChannelHandler) GetChannelLogoImage(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")

	_ = channelId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "logo image"})
}

// GetChannelBannerImage handles GET /Channels/{id}/BannerImage
func (h *ChannelHandler) GetChannelBannerImage(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")

	_ = channelId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "banner image"})
}

// GetChannelBackdropImage handles GET /Channels/{id}/BackdropImage
func (h *ChannelHandler) GetChannelBackdropImage(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")

	_ = channelId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "backdrop image"})
}

// GetChannelThumbImage handles GET /Channels/{id}/ThumbImage
func (h *ChannelHandler) GetChannelThumbImage(w http.ResponseWriter, r *http.Request) {
	channelId := chi.URLParam(r, "id")

	_ = channelId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "thumb image"})
}
