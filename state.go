package demo

import (
	"encoding/base64"
	"encoding/json"
)

type AppState struct {
	Domain string `json:"domain"`
}

func (s AppState) encodeState() string {
	m, _ := json.Marshal(s)
	state := base64.StdEncoding.EncodeToString(m)
	return state
}

func (n *AppState) decodeState(s string) {
	str, err := base64.StdEncoding.DecodeString(s)

	if err == nil {
		json.Unmarshal(str, n)
	}
}
