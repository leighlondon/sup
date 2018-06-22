package sup

// Storer is the usage contract that the backend agrees to, and is a
// generic interface and can swap in backends so long as they support
// a simple key-value model.
type Storer interface {
	Put(string, string)
	Get(string) (string, error)
	Load() error
	Save() error
	Filename() string
}

// Version number in semver.
const Version = "0.9.0"
