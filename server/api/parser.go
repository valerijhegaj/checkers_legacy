package api

import "encoding/json"

type Helper struct {
	Password string `json:"password,omitempty"`
	UserName string `json:"username,omitempty"`
	MaxAge   int    `json:"max_age,omitempty"`
}

func Parse(data []byte) (
	Helper,
	error,
) {
	var h Helper
	err := json.Unmarshal(data, &h)
	return h, err
}

func UnParse(data Helper) []byte {
	rawData, _ := json.Marshal(data)
	return rawData
}
