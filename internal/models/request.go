package models

type GenerateRequest struct {
	RawInput string `json:"rawInput"`
	Product  string `json:"product"`
	Audience string `json:"audience"`
	Tone     string `json:"tone"`
}
