package localfile

// File struct
type File struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// NewFile struct
func NewFile(name string, URL string, description string) *File {
	return &File{
		Name:        name,
		URL:         URL,
		Description: description,
	}
}
