package domain

type App struct {
	ID         int    `json:"id"`
	AppID      string `json:"appId"`
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	LastPolled string `json:"lastPolled"`
}
