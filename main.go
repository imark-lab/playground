package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type JSON struct {
	Hits          int    `json:"hits"`
	Start         int    `json:"start"`
	Engine        string `json:"_engine"`
	ReturnedTypes []int  `json:"returnedTypes"`
	Next          string `json:"next"`
	HasNext       bool   `json:"hasNext"`
	Results       []struct {
		ID      int    `json:"id"`
		Slug    string `json:"slug"`
		Type    int    `json:"type"`
		Name    string `json:"name"`
		Profile struct {
			ID           int    `json:"id"`
			DisplayName  string `json:"displayName"`
			Username     string `json:"username"`
			Image        string `json:"image"`
			ImageVersion string `json:"imageVersion"`
			City         string `json:"city"`
			Country      struct {
				ID        int    `json:"id"`
				Code      string `json:"code"`
				Name      string `json:"name"`
				Lat       string `json:"lat"`
				Long      string `json:"long"`
				Continent struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"continent"`
			} `json:"country"`
			Tier int `json:"tier"`
		} `json:"profile"`
		Genre struct {
			Name       string `json:"name"`
			ID         int    `json:"id"`
			IsArchived bool   `json:"isArchived"`
		} `json:"genre"`
		Performers []struct {
			ID        int         `json:"id"`
			Performer string      `json:"performer"`
			Profile   interface{} `json:"profile"`
		} `json:"performers"`
		Bpm              int           `json:"bpm"`
		Key              string        `json:"key"`
		Tags             []interface{} `json:"tags"`
		Image            string        `json:"image"`
		ImageVersion     string        `json:"imageVersion"`
		CompositionTitle string        `json:"compositionTitle"`
		HasAudio         bool          `json:"hasAudio"`
		Audio            string        `json:"audio"`
		Waveform         string        `json:"waveform"`
		PreviewWaveform  string        `json:"previewWaveform"`
		Created          time.Time     `json:"created"`
		Flags            []int         `json:"flags"`
		Status           int           `json:"status"`
		Prelicense       interface{}   `json:"prelicense"`
		Remixer          struct {
			ID            int    `json:"id"`
			DisplayName   string `json:"displayName"`
			Username      string `json:"username"`
			Image         string `json:"image"`
			ImageVersion  string `json:"imageVersion"`
			IsDeactivated bool   `json:"isDeactivated"`
			IsDeleted     bool   `json:"isDeleted"`
		} `json:"remixer"`
		Contest struct {
			ID     int    `json:"id"`
			Slug   string `json:"slug"`
			Status int    `json:"status"`
			Type   int    `json:"type"`
			Name   string `json:"name"`
			Flags  []int  `json:"flags"`
		} `json:"contest"`
		Prelicenses []interface{} `json:"prelicenses"`
		Original    struct {
			ID      int    `json:"id"`
			Slug    string `json:"slug"`
			Profile struct {
				ID          int    `json:"id"`
				Username    string `json:"username"`
				IsDeleted   bool   `json:"isDeleted"`
				DisplayName string `json:"displayName"`
			} `json:"profile"`
			Status           int    `json:"status"`
			CompositionTitle string `json:"compositionTitle"`
		} `json:"original"`
		StemCount    int       `json:"stemCount"`
		FirstPosted  time.Time `json:"firstPosted"`
		DocumentType int       `json:"__document_type"`
	} `json:"results"`
	NumberOfResults int `json:"numberOfResults"`
}

func main() {

	c := colly.NewCollector()

	//ここでタイトルを取得
	c.OnScraped(func(r *colly.Response) {
		var jsonData JSON
		json.Unmarshal(r.Body, &jsonData)

		for _, result := range jsonData.Results {
			if result.FirstPosted.Month() == 12 {
				if result.FirstPosted.Day() > 28 && result.FirstPosted.Day() < 31 {
					fmt.Println(result.CompositionTitle, result.Profile.Username, result.FirstPosted)
				}
			}
		}
	})

	for i := 1; i < 8; i++ {
		url := fmt.Sprintf("https://api.skiomusic.com/discover/projects/contest/364?page=%d&page_size=20&engine=db", i)
		c.Visit(url)
	}

}
