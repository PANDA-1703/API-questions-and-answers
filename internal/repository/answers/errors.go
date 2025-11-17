package answers

import "github.com/go-openapi/errors"

var (
	ErrAnswerNotFound            = errors.New(404, "answer with this ID not found")
	ErrAnswerNotFoundOrForbidden = errors.New(403, "answer not found or user is not the owner")
)
