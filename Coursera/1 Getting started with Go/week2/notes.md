## 2.1.3 Deallocation of memory


## 2.1.4 Garbage Collection
- Go is compiled languages which enables garbage collection
- Implementation is fast
- Garbage collected in background

## 2.2.1 Integers
- int8 int16 int32 int64
- uint8 ...

## 2.2.2 Ints, Floats, Strings
### Type conversion
```go
var x int32 = 1
var y int16 = 2
x = int32(y)
```
### Float
float32 - ~6 digits of precision
float64 - ~15 digits of precision

### Complex
`var z complex128 = complex(2,3)`

## 2.2.3 Strings
Strings are immutable. Modified strings are returned
### String package
- `Compare(a,b)`
- `Contains(s, substr)`
- `HasPrefix(s, prefix)`
- `Index(s, substr)`
- `TrimSpaces(s)`

### Strconv Package
Atoi(s) - String s to int
Itoa(s) - Int to string
ParseFloat(s) - String to Float

## 2.3.1 Constants
```go
type Grades int
const (
	A Grades = iota
	B
	C
	D
	F
)
```

## 2.3.3 Flow Control
### Tagless Switch
```go
switch {
	case x > 1:
		fmt.Printf("case 1")
	case x < -1:
		fmt.Println("case 2")
	default:
		fmt.Println("nocase")
}
```
