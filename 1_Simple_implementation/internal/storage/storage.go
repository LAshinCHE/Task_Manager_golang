package storage

type Storage struct {
	filename string
}

func NewStorage(f string) *Storage {
	return &Storage{
		filename: f,
	}
}
