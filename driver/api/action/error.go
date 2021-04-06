package action

import "encoding/json"

type appError struct {
	Message string `json:"error"`
}

func FormatJSONError(message string) []byte {
	appE := appError{
		Message: message,
	}
	body, err := json.Marshal(appE)

	if err != nil {
		return []byte(err.Error())
	}
	return body
}
