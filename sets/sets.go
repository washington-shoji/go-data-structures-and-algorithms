package sets

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (set *Set) NewSet() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]

	return exists
}

func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := &Set{}
	intersectSet.NewSet()

	var value int

	for value, _ = range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}

	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.NewSet()

	var value int
	for value, _ = range set.integerMap {
		unionSet.AddElement(value)
	}

	for value, _ = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}

	return unionSet

}

func SetsExec() {
	set := &Set{}
	set.NewSet()

	set.AddElement(1)
	set.AddElement(2)

	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))

	anotherSet := &Set{}
	anotherSet.NewSet()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)

	fmt.Println("another set", anotherSet)
	fmt.Println("intersection of sets", set.Intersect(anotherSet))
	fmt.Println("union of sets", set.Union(anotherSet))
}
