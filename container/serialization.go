package container

type JsonSerializer interface {
	ToJSON() ([]byte, error)

	MarshalJSON() ([]byte, error)
}

type JsonDeserializer interface {
	FromJSON([]byte) error

	UnmarshalJSON([]byte) error
}
