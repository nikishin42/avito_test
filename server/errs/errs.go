package errs

import (
	"encoding/json"
)

type AppError struct {
	Err       *Err        `json:"error"`
	ErrorData interface{} `json:"errorData,omitempty"`
}

type Err struct {
	Source  error
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (e *AppError) Error() string {
	out, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(out)
}
