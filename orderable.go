package maps

// orderable is an internal constraint that allows any type that supports
// the operators: <, <=, >=, >.
//
// Source: https://pkg.go.dev/golang.org/x/exp/constraints#Ordered
type orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}
