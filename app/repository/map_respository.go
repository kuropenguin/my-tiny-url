package repository

func NewMapRepository() IRepository {
	return &MapRepository{
		tinyURLStorage: make(map[string]string),
		urlStorage:     make(map[string]string),
	}
}

type MapRepository struct {
	tinyURLStorage map[string]string
	urlStorage     map[string]string
}

func (m *MapRepository) Create(url string, tinyURL string) error {
	m.tinyURLStorage[tinyURL] = url
	m.urlStorage[url] = tinyURL
	return nil
}

func (m *MapRepository) FindbyTinyURL(tinyURL string) (string, error) {
	if ok := m.tinyURLStorage[tinyURL]; ok != "" {
		return m.tinyURLStorage[tinyURL], nil
	}
	return "", ErrNotFound
}

func (m *MapRepository) FindbyURL(url string) (string, error) {
	if ok := m.urlStorage[url]; ok != "" {
		return m.urlStorage[url], nil
	}
	return "", ErrNotFound
}
