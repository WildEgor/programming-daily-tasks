package hash_set

type HashSet struct {
	set map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{set: make(map[interface{}]bool)}
}

func (h *HashSet) Add(key interface{}) {
	h.set[key] = true
}

func (h *HashSet) Remove(key interface{}) {
	delete(h.set, key)
}

func (h *HashSet) First() interface{} {
	for key := range h.set {
		return key
	}
	return nil
}

func (h *HashSet) Has(key interface{}) bool {
	return h.set[key]
}

func (h *HashSet) Size() int {
	return len(h.set)
}

func (h *HashSet) Clear() {
	h.set = make(map[interface{}]bool)
}

func (h *HashSet) Values() []interface{} {
	values := make([]interface{}, 0, h.Size())
	for key := range h.set {
		values = append(values, key)
	}
	return values
}

func (h *HashSet) Equals(other *HashSet) bool {
	if h.Size() != other.Size() {
		return false
	}
	for key := range h.set {
		if !other.Has(key) {
			return false
		}
	}
	return true
}
