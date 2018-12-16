package linq

// Enumerable defines main interface for collections
type Enumerable interface {
	GetList() []interface{}
	Compare(interface{}) int
}

// Linq defines main structure on collection
type Linq struct {
	coll interface{}
}

// New creates init of the linq
func New(data interface{}) *Linq {
	return &Linq{
		coll: data,
	}
}
