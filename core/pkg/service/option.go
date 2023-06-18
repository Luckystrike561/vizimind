package service

// Option holds generic options for Functional Options Pattern.
type Option[T any] func(T)
