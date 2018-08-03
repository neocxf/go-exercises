package models

import "github.com/pkg/errors"

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR       = errors.New("Your requested item was not found")
	CONFLICT_ERROR        = errors.New("Your item already exists")
)
