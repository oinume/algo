package hash_table

// Map is a generic hash table interface.
type Map[K comparable, V any] interface {
	// Put adds or updates a key-value pair. Returns the old value if the key existed, or zero value if new.
	Put(key K, value V) (V, error)
	// Get returns the value for the key, or an error if not found.
	Get(key K) (V, error)
	// Remove removes the key and returns its value, or an error if not found.
	Remove(key K) (V, error)
	// Size returns the number of entries in the map.
	Size() int
}
