package utils

import (
	cfg "rest-api-echo-go/src/config"
)

var DB = *cfg.DB

type CustomResponses map[string]interface{}