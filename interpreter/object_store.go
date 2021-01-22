package interpreter

type ObjectStore struct {
	data map[*string]*PingObj
}

func NewObjectStore() *ObjectStore {
	return &ObjectStore{make(map[*string]*PingObj)}
}
func (o *ObjectStore) Destroy(hash *string) {
	o.data[hash] = nil
}
func (o *ObjectStore) Add(hash *PingObj) {
	m, e := o.data[hash.hashid]
	if e {
		m.Dec(o)
	}
	o.data[hash.hashid] = hash
}
