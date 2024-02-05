package model

import "time"

type User struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Birthday        string    `json:"birthday,omitempty"`
	Location        string    `json:"location,omitempty"`
	JoinedAt        time.Time `json:"joined_at,omitempty"`
	AnimeStatistics struct {
		NumItemsWatching    int `json:"num_items_watching,omitempty"`
		NumItemsCompleted   int `json:"num_items_completed,omitempty"`
		NumItemsOnHold      int `json:"num_items_on_hold,omitempty"`
		NumItemsDropped     int `json:"num_items_dropped,omitempty"`
		NumItemsPlanToWatch int `json:"num_items_plan_to_watch,omitempty"`
		NumItems            int `json:"num_items,omitempty"`
		NumDaysWatched      int `json:"num_days_watched,omitempty"`
		NumDaysWatching     int `json:"num_days_watching,omitempty"`
		NumDaysCompleted    int `json:"num_days_completed,omitempty"`
		NumDaysOnHold       int `json:"num_days_on_hold,omitempty"`
		NumDaysDropped      int `json:"num_days_dropped,omitempty"`
		NumDays             int `json:"num_days,omitempty"`
		NumEpisodes         int `json:"num_episodes,omitempty"`
		NumTimesRewatched   int `json:"num_times_rewatched,omitempty"`
		MeanScore           int `json:"mean_score,omitempty"`
	} `json:"anime_statistics,omitempty"`
}
