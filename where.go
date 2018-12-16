package linq

// Where provides filtering on collection
func (l *Linq) Where(c interface{}, f func(d interface{}) bool) *Linq {
	for i := range c {
		if f(i) {
			l.coll = append(l.coll, i)
		}
	}
	return l
}
