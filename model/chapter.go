package model

import "time"

type ChapterResponse struct {
	Data   []Data `json:"data,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Total  int    `json:"total,omitempty"`
}

type Data struct {
	ID            string          `json:"id,omitempty"`
	Type          string          `json:"type,omitempty"`
	Attributes    Attributes      `json:"attributes,omitempty"`
	Relationships []Relationships `json:"relationships,omitempty"`
}

type Attributes struct {
	Volume             any       `json:"volume,omitempty"`
	Chapter            string    `json:"chapter,omitempty"`
	Title              any       `json:"title,omitempty"`
	TranslatedLanguage string    `json:"translatedLanguage,omitempty"`
	ExternalURL        any       `json:"externalUrl,omitempty"`
	PublishAt          time.Time `json:"publishAt,omitempty"`
	ReadableAt         time.Time `json:"readableAt,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	Pages              int       `json:"pages,omitempty"`
	Version            int       `json:"version,omitempty"`
}

type Relationships struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type ChapterData struct {
	BaseURL string  `json:"baseUrl,omitempty"`
	Chapter ChapterList `json:"chapter,omitempty"`
}

type ChapterList struct {
	Hash      string   `json:"hash,omitempty"`
	Data      []string `json:"data,omitempty"`
	DataSaver []string `json:"dataSaver,omitempty"`
}
