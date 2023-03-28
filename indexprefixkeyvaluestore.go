type MyStore struct {
    keyValues map[string][]byte
    indexMap  map[string][]string
}

func NewMyStore() *MyStore {
    return &MyStore{
        keyValues: make(map[string][]byte),
        indexMap:  make(map[string][]string),
    }
}

func (s *MyStore) Get(ctx sdk.Context, key []byte) []byte {
    return s.keyValues[string(key)]
}

func (s *MyStore) Set(ctx sdk.Context, key []byte, value []byte) {
    s.keyValues[string(key)] = value
    // Update index map with new key
    for prefix, _ := range s.indexMap {
        if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
            s.indexMap[prefix] = append(s.indexMap[prefix], string(key))
        }
    }
}

func (s *MyStore) Delete(ctx sdk.Context, key []byte) {
    delete(s.keyValues, string(key))
    // Remove key from index map
    for prefix, values := range s.indexMap {
        for i, value := range values {
            if value == string(key) {
                s.indexMap[prefix] = append(values[:i], values[i+1:]...)
            }
        }
    }
}

// Implement a function that retrieves all keys in the index map that match a specific prefix
func (s *MyStore) GetIndexKeys(ctx sdk.Context, prefix string) []string {
    if keys, ok := s.indexMap[prefix]; ok {
        return keys
    }
    return []string{}
}

// Implement a function that adds a new index prefix to the store
func (s *MyStore) AddIndexPrefix(ctx sdk.Context, prefix string) {
    s.indexMap[prefix] = []string{}
}
