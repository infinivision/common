package functor

import "github.com/infinivision/common/typeclass"

// MapFunc map a value to a new value
type MapFunc (func(typeclass.Ord) typeclass.Ord)

// a function that analyze a recursive data structure and through use of a given combining operation
type FoldFunc (func(typeclass.Ord, interface{}) interface{})

// functional Functor
type Functor interface {
	Null() bool                       // Test whether the structure is empty
	Length() int                      // Returns the size/length of a finite structure as an Int
	Minimum() typeclass.Ord           // The least element of a non-empty structure
	Maximum() typeclass.Ord           // The largest element of a non-empty structure
	Elem(typeclass.Ord) typeclass.Ord // Does the element occur in the structure

	Map(MapFunc) Functor                     // apply f to a structure then create a new structure
	Foldl(FoldFunc, interface{}) interface{} // Left-associative fold of a structure
	Foldr(FoldFunc, interface{}) interface{} // Right-associative fold of a structure
}
