package golang_ippanel

type PatternRequest struct {
	Code      string         `json:"code"`
	Sender    string         `json:"sender"`
	Recipient string         `json:"recipient"`
	Variable  map[string]any `json:"variable"`
}
