package utils

import (
	"encoding/json"

	"github.com/labstack/echo"
)

// ParseReqBody: parse request body
func ParseReqBody(c echo.Context) (map[string]interface{}, error) {
	bodyJson := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&bodyJson)
	if err != nil {
		return bodyJson, err
	}

	return bodyJson, nil
}
