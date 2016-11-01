package algo

const (
	DEFAULT_HASH_SIZE = 4096
)

type HashEntry struct {
	originalKey string
	data string
	next *HashEntry
}

type HashTable struct {
	table []*HashEntry
	size uint32
	entries uint32
}

func NewHashTable(size uint32) *HashTable {
	return &HashTable{ table: make([]*HashEntry, size), size: size  }
}

func (h *HashTable) Set(key string, value string) {
	hash := crc32.ChecksumIEEE([]byte(key))
	idx := hash % h.size
	entry := h.table[idx]

	newEntry := &HashEntry{originalKey: key, data: value, next: nil}

	// no collision
	if entry == nil {
		h.table[idx] = newEntry
		h.entries += 1
		return
	}

	// collision
	prev := entry
	next := entry.next
	for next != nil {
		prev = next
		next = entry.next
	}

	prev.next = newEntry
	h.entries += 1
}

func (h *HashTable) loadFactor() float32 {
	return float32(h.entries) / float32(h.size)
}

func (h *HashTable) findEntry(originalKey string, e *HashEntry) *HashEntry {
	if e.originalKey == originalKey {
		return e
	}
	if e.next == nil {
		return nil
	}
	
	return h.findEntry(originalKey, e.next)
}

func (h *HashTable) Get(key string) string {
	hash := crc32.ChecksumIEEE([]byte(key))
	idx := hash % h.size
	rootEntry := h.table[idx]

	if rootEntry == nil {
		panic("key not found")
	}
	foundEntry := h.findEntry(key, rootEntry)

	if foundEntry == nil {
		panic("key not found")
	}

	return foundEntry.data
}
