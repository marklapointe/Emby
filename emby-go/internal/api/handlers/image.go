package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/image"
	"github.com/gorilla/mux"
)

// ImageHandler handles image-related API endpoints.
type ImageHandler struct {
	imageMgr *image.Manager
}

// NewImageHandler creates a new image handler.
func NewImageHandler(imageMgr *image.Manager) *ImageHandler {
	return &ImageHandler{imageMgr: imageMgr}
}

// GetItemImage handles GET /Items/{id}/Images/{type}
func (h *ImageHandler) GetItemImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	imageType := vars["type"]

	// Get image parameters
	quality := r.URL.Query().Get("quality")
	width := r.URL.Query().Get("width")
	height := r.URL.Query().Get("height")
	tag := r.URL.Query().Get("tag")

	// Get image from manager
	img, contentType, err := h.imageMgr.GetItemImage(itemID, imageType, quality, width, height, tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("ETag", tag)
	w.Write(img)
}

// GetItemImageBlurHash handles GET /Items/{id}/Images/{type}/BlurHash
func (h *ImageHandler) GetItemImageBlurHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	imageType := vars["type"]

	blurHash, err := h.imageMgr.GetImageBlurHash(itemID, imageType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"blurHash": blurHash})
}

// GetItemImageByIndex handles GET /Items/{id}/Images/{type}/{index}
func (h *ImageHandler) GetItemImageByIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	imageType := vars["type"]
	index := vars["index"]

	quality := r.URL.Query().Get("quality")
	width := r.URL.Query().Get("width")
	height := r.URL.Query().Get("height")

	img, contentType, err := h.imageMgr.GetItemImageByIndex(itemID, imageType, index, quality, width, height)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(img)
}

// GetItemImageTag handles GET /Items/{id}/Images/{type}/Tag/{tag}
func (h *ImageHandler) GetItemImageTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	imageType := vars["type"]
	tag := vars["tag"]

	img, contentType, err := h.imageMgr.GetItemImageByTag(itemID, imageType, tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("ETag", tag)
	w.Write(img)
}

// GetItemImageCrop handles GET /Items/{id}/Images/{type}/Crop
func (h *ImageHandler) GetItemImageCrop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	imageType := vars["type"]

	// Get crop parameters
	width := r.URL.Query().GetInt("width")
	height := r.URL.Query().GetInt("height")
	cropPosition := r.URL.Query().Get("cropPosition")

	img, contentType, err := h.imageMgr.GetImageCrop(itemID, imageType, width, height, cropPosition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(img)
}

// GetItemImageResize handles GET /Items/{id}/Images/{type}/Resize
func (h *ImageHandler) GetItemImageResize(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	imageType := vars["type"]

	// Get resize parameters
	width := r.URL.Query().GetInt("width")
	height := r.URL.Query().GetInt("height")
	quality := r.URL.Query().Get("quality")

	img, contentType, err := h.imageMgr.GetImageResize(itemID, imageType, width, height, quality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(img)
}

// GetItemImageRotation handles GET /Items/{id}/Images/{type}/Rotate
func (h *ImageHandler) GetItemImageRotation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	imageType := vars["type"]

	// Get rotation parameter
	angle := r.URL.Query().GetInt("angle")

	img, contentType, err := h.imageMgr.GetImageRotation(itemID, imageType, angle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(img)
}
