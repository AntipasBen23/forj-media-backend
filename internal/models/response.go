package models

type GenerateResponse struct {
	Hooks        []string `json:"hooks"`
	PostOutlines []string `json:"postOutlines"`
	FullPosts    []string `json:"fullPosts"`
}
