package storage

// FileBackedStorage implements the sup.Storer interface.
type FileBackedStorage struct {
	data     map[string]string
	filename string
}

// New initializes a new instance of FileBackedStorage.
func New(filename string) (*FileBackedStorage, error) {
	data, err := loadJson(filename)
	return &FileBackedStorage{data: data, filename: filename}, err
}

// Filename returns the filename of the storage.
func (s *FileBackedStorage) Filename() string {
	return s.filename
}

// Load updates the state from the backing store.
func (s *FileBackedStorage) Load() error {
	data, err := loadJson(s.filename)
	if err != nil {
		return err
	}
	s.data = data
	return nil
}

// Save the state to the backing store.
func (s *FileBackedStorage) Save() error {
	return saveJson(s.filename, s.data)
}

func loadJson(filename string) (map[string]string, error) {
	return make(map[string]string), nil
}

func saveJson(filename string, data map[string]string) error {
	return nil
}
