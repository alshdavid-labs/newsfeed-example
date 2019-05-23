package newsfeed

type MemoryStore struct {
	Items []Item
}

// Get gets items from the in memory database
func (s *MemoryStore) Get() []Item {
	return s.Items
}

// Set stores items to the in memory database
func (s *MemoryStore) Set(item Item) bool {
	ok := true
	s.Items = append(s.Items, item)
	return ok
}

// InMemory makes a newsfeed in memory
func InMemory() *MemoryStore {
	return &MemoryStore{
		Items: []Item{},
	}
}
