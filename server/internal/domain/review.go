package domain

import "time"

type Review struct {
	ID        int       `json:"id"`
	AppID     string    `json:"appId"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Rating    int       `json:"rating"`
	Timestamp time.Time `json:"timestamp"`
}
