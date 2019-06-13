package typeclass

type Eq interface {
	Eq(Eq) bool
	NotEq(Eq) bool
}
