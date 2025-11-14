package repository

import "context"

type ContextGetter interface {
	Context() context.Context
}
