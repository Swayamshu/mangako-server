package model

import "time"

type MangaList struct {
	Data []MangaData `json:"data"`
}

type MangaResponse struct {
	Data MangaData `json:"data"`
}

type MangaFormData struct {
	Title string `json:"title,omitempty"`
}

type MangaData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Title struct {
			En string `json:"en"`
		} `json:"title"`
		AltTitles []struct {
			Ko   string `json:"ko,omitempty"`
			My   string `json:"my,omitempty"`
			Th   string `json:"th,omitempty"`
			Bn   string `json:"bn,omitempty"`
			Ne   string `json:"ne,omitempty"`
			Zh   string `json:"zh,omitempty"`
			ZhHk string `json:"zh-hk,omitempty"`
			Mn   string `json:"mn,omitempty"`
			Ar   string `json:"ar,omitempty"`
			Fa   string `json:"fa,omitempty"`
			He   string `json:"he,omitempty"`
			Vi   string `json:"vi,omitempty"`
			Ru   string `json:"ru,omitempty"`
			Ms   string `json:"ms,omitempty"`
			Uk   string `json:"uk,omitempty"`
			Ta   string `json:"ta,omitempty"`
			Hi   string `json:"hi,omitempty"`
			Kk   string `json:"kk,omitempty"`
			Ja   string `json:"ja,omitempty"`
		} `json:"altTitles"`
		Description struct {
			En   string `json:"en"`
			Ru   string `json:"ru"`
			Uk   string `json:"uk"`
			EsLa string `json:"es-la"`
			PtBr string `json:"pt-br"`
		} `json:"description"`
		IsLocked bool `json:"isLocked"`
		Links    struct {
			Al    string `json:"al"`
			Ap    string `json:"ap"`
			Bw    string `json:"bw"`
			Kt    string `json:"kt"`
			Mu    string `json:"mu"`
			Amz   string `json:"amz"`
			Cdj   string `json:"cdj"`
			Ebj   string `json:"ebj"`
			Mal   string `json:"mal"`
			Raw   string `json:"raw"`
			Engtl string `json:"engtl"`
		} `json:"links"`
		OriginalLanguage       string `json:"originalLanguage"`
		LastVolume             string `json:"lastVolume"`
		LastChapter            string `json:"lastChapter"`
		PublicationDemographic string `json:"publicationDemographic"`
		Status                 string `json:"status"`
		Year                   int    `json:"year"`
		ContentRating          string `json:"contentRating"`
		Tags                   []struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Attributes struct {
				Name struct {
					En string `json:"en"`
				} `json:"name"`
				Description struct {
				} `json:"description"`
				Group   string `json:"group"`
				Version int    `json:"version"`
			} `json:"attributes"`
			Relationships []any `json:"relationships"`
		} `json:"tags"`
		State                          string    `json:"state"`
		ChapterNumbersResetOnNewVolume bool      `json:"chapterNumbersResetOnNewVolume"`
		CreatedAt                      time.Time `json:"createdAt"`
		UpdatedAt                      time.Time `json:"updatedAt"`
		Version                        int       `json:"version"`
		AvailableTranslatedLanguages   []string  `json:"availableTranslatedLanguages"`
		LatestUploadedChapter          string    `json:"latestUploadedChapter"`
	} `json:"attributes"`
	Relationships []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Related string `json:"related,omitempty"`
	} `json:"relationships"`
}

type CoverArtResponse struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Description string    `json:"description"`
			Volume      string    `json:"volume"`
			FileName    string    `json:"fileName"`
			Locale      string    `json:"locale"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			Version     int       `json:"version"`
		} `json:"attributes"`
		Relationships []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"relationships"`
	} `json:"data"`
}

type MangaDetails struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	CoverImage  string `json:"cover_image,omitempty"`
	MainPicture struct {
		Medium string `json:"medium,omitempty"`
		Large  string `json:"large,omitempty"`
	} `json:"main_picture,omitempty"`
	AlternativeTitles struct {
		Synonyms []string `json:"synonyms,omitempty"`
		En       string   `json:"en,omitempty"`
		Ja       string   `json:"ja,omitempty"`
	} `json:"alternative_titles,omitempty"`
	StartDate       string    `json:"start_date,omitempty"`
	Synopsis        string    `json:"synopsis,omitempty"`
	Mean            float64   `json:"mean,omitempty"`
	Rank            int       `json:"rank,omitempty"`
	Popularity      int       `json:"popularity,omitempty"`
	NumListUsers    int       `json:"num_list_users,omitempty"`
	NumScoringUsers int       `json:"num_scoring_users,omitempty"`
	Nsfw            string    `json:"nsfw,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	MediaType       string    `json:"media_type,omitempty"`
	Status          string    `json:"status,omitempty"`
	Genres          []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"genres,omitempty"`
	NumVolumes  int `json:"num_volumes,omitempty"`
	NumChapters int `json:"num_chapters,omitempty"`
	Authors     []struct {
		Node struct {
			ID        int    `json:"id,omitempty"`
			FirstName string `json:"first_name,omitempty"`
			LastName  string `json:"last_name,omitempty"`
		} `json:"node,omitempty"`
		Role string `json:"role,omitempty"`
	} `json:"authors,omitempty"`
	Pictures []struct {
		Medium string `json:"medium,omitempty"`
		Large  string `json:"large,omitempty"`
	} `json:"pictures,omitempty"`
	Background   string `json:"background,omitempty"`
	RelatedAnime []any  `json:"related_anime,omitempty"`
	RelatedManga []struct {
		Node struct {
			ID          int    `json:"id,omitempty"`
			Title       string `json:"title,omitempty"`
			MainPicture struct {
				Medium string `json:"medium,omitempty"`
				Large  string `json:"large,omitempty"`
			} `json:"main_picture,omitempty"`
		} `json:"node,omitempty"`
		RelationType          string `json:"relation_type,omitempty"`
		RelationTypeFormatted string `json:"relation_type_formatted,omitempty"`
	} `json:"related_manga,omitempty"`
	Recommendations []struct {
		Node struct {
			ID          int    `json:"id,omitempty"`
			Title       string `json:"title,omitempty"`
			MainPicture struct {
				Medium string `json:"medium,omitempty"`
				Large  string `json:"large,omitempty"`
			} `json:"main_picture,omitempty"`
		} `json:"node,omitempty"`
		NumRecommendations int `json:"num_recommendations,omitempty"`
	} `json:"recommendations,omitempty"`
	Serialization []struct {
		Node struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"node,omitempty"`
	} `json:"serialization,omitempty"`
}

type TrendingMangaList struct {
	Data []struct {
		Node struct {
			ID          int    `json:"id,omitempty"`
			Title       string `json:"title,omitempty"`
			MainPicture struct {
				Medium string `json:"medium,omitempty"`
				Large  string `json:"large,omitempty"`
			} `json:"main_picture,omitempty"`
		} `json:"node,omitempty"`
		Ranking struct {
			Rank int `json:"rank,omitempty"`
		} `json:"ranking,omitempty"`
	} `json:"data,omitempty"`
	Paging struct {
		Next string `json:"next,omitempty"`
	} `json:"paging,omitempty"`
}

type MapperData struct {
	MALID      string `json:"mal_id,omitempty"`
	MangaDexID string `json:"mangadex_id,omitempty"`
	Title      string `json:"title,omitempty"`
}
