package questions

import "github.com/go-openapi/errors"

var (
	ErrQuestionNotFound = errors.New(404, "question with this ID not found")
)
