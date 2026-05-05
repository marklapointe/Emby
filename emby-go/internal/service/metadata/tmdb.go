package metadata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TMDbProvider struct {
	apiKey     string
	language   string
	httpClient *http.Client
}

type TMDbMovie struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Overview     string   `json:"overview"`
	PosterPath   string   `json:"poster_path"`
	BackdropPath string   `json:"backdrop_path"`
	ReleaseDate  string   `json:"release_date"`
	VoteAverage  float64  `json:"vote_average"`
	Genres       []TMDbGenre `json:"genres"`
	Runtime      int      `json:"runtime"`
	Tagline      string   `json:"tagline"`
}

type TMDbTVShow struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Overview     string      `json:"overview"`
	PosterPath   string      `json:"poster_path"`
	BackdropPath string      `json:"backdrop_path"`
	FirstAirDate string      `json:"first_air_date"`
	VoteAverage  float64     `json:"vote_average"`
	Genres       []TMDbGenre `json:"genres"`
	EpisodeRunTime []int    `json:"episode_run_time"`
	Networks     []TMDbNetwork `json:"networks"`
}

type TMDbGenre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TMDbNetwork struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewTMDbProvider(apiKey string) *TMDbProvider {
	return &TMDbProvider{
		apiKey:   apiKey,
		language: "en-US",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (p *TMDbProvider) GetMovie(tmdbID int) (*TMDbMovie, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?api_key=%s&language=%s", tmdbID, p.apiKey, p.language)
	return p.fetchMovie(url)
}

func (p *TMDbProvider) SearchMovie(query string) ([]TMDbMovie, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&language=%s&query=%s", p.apiKey, p.language, query)
	return p.searchMovies(url)
}

func (p *TMDbProvider) SearchTVShow(query string) ([]TMDbTVShow, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/tv?api_key=%s&language=%s&query=%s", p.apiKey, p.language, query)
	return p.searchTVShows(url)
}

func (p *TMDbProvider) GetTVShow(tmdbID int) (*TMDbTVShow, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d?api_key=%s&language=%s", tmdbID, p.apiKey, p.language)
	return p.fetchTVShow(url)
}

func (p *TMDbProvider) fetchMovie(url string) (*TMDbMovie, error) {
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDb API error: HTTP %d", resp.StatusCode)
	}

	var movie TMDbMovie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}

	return &movie, nil
}

func (p *TMDbProvider) searchMovies(url string) ([]TMDbMovie, error) {
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDb API error: HTTP %d", resp.StatusCode)
	}

	var result struct {
		Results []TMDbMovie `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (p *TMDbProvider) fetchTVShow(url string) (*TMDbTVShow, error) {
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDb API error: HTTP %d", resp.StatusCode)
	}

	var show TMDbTVShow
	if err := json.NewDecoder(resp.Body).Decode(&show); err != nil {
		return nil, err
	}

	return &show, nil
}

func (p *TMDbProvider) searchTVShows(url string) ([]TMDbTVShow, error) {
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDb API error: HTTP %d", resp.StatusCode)
	}

	var result struct {
		Results []TMDbTVShow `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func (p *TMDbProvider) GetImageURL(path string, size string) string {
	return fmt.Sprintf("https://image.tmdb.org/t/p/%s%s", size, path)
}