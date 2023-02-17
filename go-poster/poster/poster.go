package poster

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	OmdbURL   = "https://omdbapi.com/demo.aspx/"
	OmdbToken = "03AFY_a8VKA30bVg79yxdp2SseRICRhRd3MIS8JmVxD30z_e0GExlUP1IjbfiyWgsU3hqRo-yUTQstf8xSSuT5XdCql2s9JhzzVZA1ogTPTsDGRAKkf_8mQ_UhDvLEGMappaN8qLGRJBJi2lUg8Op7i5pQEwsbeYL5qCq4L_c7F4N4WS5S1zvVU_B7dJEvzyP8iYn22nwCBXRB1FhdJ7CQhEqWX8gTIr-zpMAKMX_juLDr0852cifgZQqUJUdvQIVVy-uEPJCn060O-zpUSYFuFhGGjS0slzzmHYQG14_qrrkvc26EB8P1pvJg2dOLjqprSngChFMpFOvIkQoyqebuvchrDgSTsPHwwCWRGDOPZKi4RVHuM3DHc4uvnIJAE1-DZz_pmSr9KFaqm6fRocU4i3sArpG6DxBmztlQq5DeUB-HI9yxo1f1dN-QfIjapGO5DbIomlTXjO31OkTz8wAnzwVM1dj8b5C3UScs19GOrzT3cHchhP3SIsHIWXBy2sdatdprKFz_h2_I"
)

type MovieInfo struct {
	Title    string
	Year     string
	Rated    string
	Director string
	Poster   string
}
type Poster struct {
}

func NewPoster() *Poster {
	return &Poster{}
}

func (p *Poster) Query(name string) (string, error) {
	url := fmt.Sprintf("%s?t=%s&token=%s", OmdbURL, name, OmdbToken)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var movie MovieInfo
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return "", err
	}
	resp2, err := http.Get(movie.Poster)
	if err != nil {
		return "", err
	}
	defer resp2.Body.Close()
	data, err := io.ReadAll(resp2.Body)
	if err != nil {
		return "", err
	}
	fname := fmt.Sprintf("%s.jpg", name)
	file, err := os.Create(fname)
	if err != nil {
		return "", err
	}
	defer file.Close()
	file.Write(data)
	return fname, nil
}
