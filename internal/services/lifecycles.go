package services

type lifecycle int

const (
	Transietnt lifecycle = iota
	Singleton
	Scoped
)
