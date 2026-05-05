package server

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type StaticHandler struct {
	fs       http.FileSystem
	basePath string
	version  string
}

func NewStaticHandler(basePath, version string) *StaticHandler {
	return &StaticHandler{
		fs:       http.Dir(basePath),
		basePath: basePath,
		version:  version,
	}
}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := h.fs.Open(r.URL.Path)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

	if fi.IsDir() {
		indexPath := path.Join(r.URL.Path, "index.html")
		f2, err := h.fs.Open(indexPath)
		if err != nil {
			http.ServeFile(w, r, r.URL.Path)
			return
		}
		defer f2.Close()
		fi2, _ := f2.Stat()
		f = f2
		fi = fi2
		r.URL.Path = indexPath
	}

	contentType := mimeType(path.Ext(r.URL.Path))
	w.Header().Set("Content-Type", contentType)

	if strings.HasSuffix(r.URL.Path, ".html") {
		data, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, "500 Internal Server Error", 500)
			return
		}
		html := string(data)
		html = h.modifyHTML(r.URL.Path, html)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}

func (h *StaticHandler) modifyHTML(filePath string, html string) string {
	isIndex := strings.EqualFold(filePath, "index.html") || strings.EqualFold(filePath, "/index.html")

	if isIndex {
		if !strings.Contains(html, "data-culture") {
			html = strings.Replace(html, "<html", `<html lang="en"`, 1)
		}

		if strings.Contains(html, "<script") {
			html = strings.Replace(html, "<script", "<!--<script", 1)
			html = strings.Replace(html, "</script>", "</script>-->", 1)
		}
	}

	if isIndex {
		scriptTag := `<script>window.dashboardVersion='` + h.version + `';</script>
<script src="scripts/apploader.js?v=` + h.version + `" defer></script>`

		html = strings.Replace(html, "</body>", scriptTag+"</body>", 1)
	}

	return html
}

func mimeType(ext string) string {
	ext = strings.ToLower(ext)
	switch ext {
	case ".html", ".htm":
		return "text/html; charset=utf-8"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".ttf":
		return "font/ttf"
	case ".eot":
		return "application/vnd.ms-fontobject"
	case ".otf":
		return "font/otf"
	default:
		return "application/octet-stream"
	}
}