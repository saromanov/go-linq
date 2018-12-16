package linq

// Enumerable defines main interface for collections
type Enumerable interface {
	GetList() []interface{}
	Compare(interface{}) int
}
