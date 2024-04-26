// Code generated by goctl. DO NOT EDIT.
package types

type ListVoicesRequest struct {
	Platform string `path:"platform"`
}

type ListVoicesResponse struct {
	Voices []*Voice `json:"voices"`
}

type TextToSpeechRequest struct {
	Text     string  `json:"text"`
	Voice    string  `json:"voice"`
	Platform string  `json:"platform"`
	Volume   float64 `json:"volume"`
	Pitch    float64 `json:"pitch"`
	Rate     float64 `json:"rate"`
}

type TextToSpeechResponse struct {
	Audio []byte `json:"audio"`
}

type Voice struct {
	Name               string   `json:"name"`
	ShortName          string   `json:"shortName"`
	Gender             string   `json:"gender"`
	Locale             string   `json:"locale"`
	FriendlyName       string   `json:"friendlyName"`
	ContentCategories  []string `json:"contentCategories"`
	VoicePersonalities []string `json:"voicePersonalities"`
}
