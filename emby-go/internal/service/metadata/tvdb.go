package metadata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TVDbProvider struct {
	apiKey     string
	language   string
	httpClient *http.Client
}

type TVDbSeries struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Overview    string    `json:"overview"`
	Poster      string    `json:"poster"`
	Fanart      string    `json:"fanart"`
	Banner      string    `json:"banner"`
	FirstAired  string    `json:"first_aired"`
	Network     string    `json:"network"`
	IMDBID      string    `json:"imdb_id"`
	SiteRating  float64   `json:"site_rating"`
	Episodes    []TVDbEpisode `json:"episodes"`
}

type TVDbEpisode struct {
	ID            int    `json:"id"`
	SeasonNumber  int    `json:"aired_season"`
	EpisodeNumber int    `json:"aired_episode_number"`
	Name          string `json:"episode_name"`
	Overview      string `json:"overview"`
	FirstAired    string `json:"first_aired"`
	GuestStars    []string `json:"guest_stars"`
	Director      string `json:"director"`
	Filepath      string `json:"filename"`
	Runtime       int    `json:"runtime"`
}

type TVDbActors struct {
	Actors []TVDbActor `json:"actors"`
}

type TVDbActor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Image     string `json:"image"`
	SortOrder int    `json:"sortOrder"`
}

func NewTVDbProvider(apiKey string) *TVDbProvider {
	return &TVDbProvider{
		apiKey:   apiKey,
		language: "en",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (p *TVDbProvider) GetSeries(tvdbID int) (*TVDbSeries, error) {
	url := fmt.Sprintf("https://api.thetvdb.com/series/%d", tvdbID)
	series, err := p.fetchSeries(url)
	if err != nil {
		return nil, err
	}

	episodes, err := p.GetAllEpisodes(tvdbID)
	if err != nil {
		return series, nil
	}
	series.Episodes = episodes

	return series, nil
}

func (p *TVDbProvider) SearchSeries(query string) ([]TVDbSeries, error) {
	url := fmt.Sprintf("https://api.thetvdb.com/search/series?name=%s", query)
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TVDb API error: HTTP %d", resp.StatusCode)
	}

	var result struct {
		Data []TVDbSeries `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (p *TVDbProvider) GetAllEpisodes(seriesID int) ([]TVDbEpisode, error) {
	url := fmt.Sprintf("https://api.thetvdb.com/series/%d/episodes", seriesID)
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TVDb API error: HTTP %d", resp.StatusCode)
	}

	var result struct {
		Data []TVDbEpisode `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (p *TVDbProvider) GetEpisode(seriesID int, season, episode int) (*TVDbEpisode, error) {
	url := fmt.Sprintf("https://api.thetvdb.com/series/%d/episodes/query?airedSeason=%d&airedEpisode=%d",
		seriesID, season, episode)
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TVDb API error: HTTP %d", resp.StatusCode)
	}

	var result struct {
		Data []TVDbEpisode `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("episode not found")
	}

	return &result.Data[0], nil
}

func (p *TVDbProvider) GetActors(seriesID int) ([]TVDbActor, error) {
	url := fmt.Sprintf("https://api.thetvdb.com/series/%d/actors", seriesID)
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TVDb API error: HTTP %d", resp.StatusCode)
	}

	var result TVDbActors
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Actors, nil
}

func (p *TVDbProvider) fetchSeries(url string) (*TVDbSeries, error) {
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TVDb API error: HTTP %d", resp.StatusCode)
	}

	var result struct {
		Data TVDbSeries `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func (p *TVDbProvider) GetImageURL(imageType, path string) string {
	switch imageType {
	case "poster":
		return fmt.Sprintf("https://artworks.thetvdb.com/banners/%s", path)
	case "fanart":
		return fmt.Sprintf("https://artworks.thetvdb.com/banners/%s", path)
	case "banner":
		return fmt.Sprintf("https://artworks.thetvdb.com/banners/%s", path)
	default:
		return fmt.Sprintf("https://artworks.thetvdb.com/banners/%s", path)
	}
}