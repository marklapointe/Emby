package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	img, contentType, err := h.imageMgr.GetItemImage(itemID, image.ImageType(imageType), quality, width, height, tag)
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

	blurHash, err := h.imageMgr.GetImageBlurHash(itemID, image.ImageType(imageType))
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

	img, contentType, err := h.imageMgr.GetItemImageByIndex(itemID, image.ImageType(imageType), 0, 0, 0, 0)
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

	img, contentType, err := h.imageMgr.GetItemImageByTag(itemID, image.ImageType(imageType), tag)
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

	width, _ := strconv.Atoi(r.URL.Query().Get("width"))
	height, _ := strconv.Atoi(r.URL.Query().Get("height"))
	cropPosition := r.URL.Query().Get("cropPosition")

	img, contentType, err := h.imageMgr.GetImageCrop(itemID, image.ImageType(imageType), width, height, cropPosition)
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

	width, _ := strconv.Atoi(r.URL.Query().Get("width"))
	height, _ := strconv.Atoi(r.URL.Query().Get("height"))
	quality, _ := strconv.Atoi(r.URL.Query().Get("quality"))

	img, contentType, err := h.imageMgr.GetImageResize(itemID, image.ImageType(imageType), width, height, quality)
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

	angle, _ := strconv.Atoi(r.URL.Query().Get("angle"))

	img, contentType, err := h.imageMgr.GetImageRotation(itemID, image.ImageType(imageType), angle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(img)
}
