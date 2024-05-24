package domain

import "errors"

var (
	ErrInternal = errors.New("internal")

	ErrUserNotFound  = errors.New("user not found")
	ErrTopicNotFound = errors.New("topic not found")
)
