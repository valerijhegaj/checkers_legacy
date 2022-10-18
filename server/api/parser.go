package api

import (
	"encoding/json"

	"checkers/core"
	"checkers/server/pkg/defines"
)

type Helper struct {
	Password string            `json:"password,omitempty"`
	UserName string            `json:"username,omitempty"`
	MaxAge   int               `json:"max_age,omitempty"`
	GameName string            `json:"gamename,omitempty"`
	Settings defines.Settings  `json:"settings,omitempty"`
	From     core.Coordinate   `json:"from"`
	Way      []core.Coordinate `json:"way"`
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
