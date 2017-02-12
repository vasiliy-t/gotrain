package set

type IntSet map[int]struct{}

func (i *IntSet) Contains(elm interface{}) bool {
	_, ok := (*i)[elm.(int)]
	return ok
}

func (i *IntSet) Add(elm interface{}) {
	switch elm.(type) {
	case int:
		(*i)[elm.(int)] = struct{}{}
	default:
		panic("Could not add element of type in IntSet")
	}
}

type Set interface {
	Contains(elm interface{}) bool
	Add(elm interface{})
}

func Difference(one, other Set) Set {
	iset := &IntSet{}

	for k, _ := range *one.(*IntSet) {
		if !other.Contains(k) {
			iset.Add(k)
		}
	}
	return iset
}

func Union(one, other Set) Set {
	iset := &IntSet{}

	for k, _ := range *one.(*IntSet) {
		iset.Add(k)
	}

	for k, _ := range *other.(*IntSet) {
		iset.Add(k)
	}
	return iset
}

func Intersect(one, other Set) Set {
	iset := &IntSet{}
	o := *other.(*IntSet)
	for k, _ := range *one.(*IntSet) {
		if o.Contains(k) {
			iset.Add(k)
		}
	}

	return iset
}
