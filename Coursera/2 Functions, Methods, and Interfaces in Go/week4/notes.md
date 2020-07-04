## 4.1.1 Polymorphism
### Inheritence
Subclass inherits methods/data of the superclass.

Golang doesnot have Inheritence.

### Overriding
Subclasses override the implementation of the superclass.

Golang  doesnot have overriding

## 4.1.2 Interfaces
Set of method signatures: name, parameters and return values

Ex: `Shape2D` interface: All shapes must have `Area()` and `Perimeter()`

### Satisfying an Interface
A Type satisfies the interface if it defines all the methods specified in the interface.

Ex: Types `Rectangle` and `Triangle` satisfying `Shape2D` interface:
- must have `Area()` and `Perimeter()` methods 
- Additional methoda are OK

This is similar to inheritence with overriding

### Defining an interface type
```go
type Shape2D interface{
	Area() float64
	Perimeter() float64
}

type Triangle {...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimeter() float64 {...}
```

`Triangle` type satisfies the `Shape2D` interface.
No need to state explicitely that `Triangle` satisfies the interface, compiler does the rest.

## 4.1.3 Interface vs Concrete types
|Interface| Concrete Type|
|---|---|
|Specify **some** method signature| Specify the exact representation of data and methods|
|Implemenations are abstracted| Complete method implementation are included|

### Interface Values
Can be treated like other values
- Assigned to values
- Passed, Returned

Two components:
- **Dynamic Type** : Concrete type which it is assigned to 
- **Dynamic Value** : Value of the dynamic type

### Defining an Interface type
```go
type Speaker interface {speak()}

type Dog struct {name string}

func (d Dog) Speak(){
	fmt.Println(d.name)
}

func main(){
	var s1 Speaker 
	var d1 Dog("Brian")
	s1 = d1
	s1.speak()
}
```

`s1` is speaker type, interface value.
`Dog` is the dynamic type
`d1` is the dynamic value

### Interface with Nil dynamic value 
```go
var s1 Speaker
var d1 *Dog
s1 = d1
```
`s1` has a dynamic type but no dynamic value since `d1` has no concrete value.

### Nil Dynamic Value
```go
func (d *Dog) Speak(){
	if d == nil {fmt.Println("<noise>")}
	else {fmt.Println(d.name)}
}

var s1 Speaker
var d1 *Dog
s1 = d1
s1.Speak()
```
### Nil Dynamic Type
```go
var s1 Speaker
s1.Speak() //ERROR
```
Nil Dynamic Types don't have implementations specified and thus methods can't be called.

## 4.2.1 Using interfaces

### Empty interface
All types satisfy empty interface
Used to accept any type as parameter
```go
func PrintMe(val {}interface){
	fmt.Println(val)
}
```
## 4.2.2 Type assertions
[Type assertions](https://tour.golang.org/methods/15) in go 

### Type assertion for diasmbiguation
Concrete type in paranthesis
```go
func DrawShape(s Shape2D) bool {
	rect , ok := s.(Rectangle)
	if ok {
		DrawRect(rect)
	}
	tri, ok := s.(Triange)
	if ok {
		DrawTriangle(tri)
	}
}
```

### Type Switch
Switch statement with type assertion
```go
	switch:= sh:= s.(type){
	case Rectangle:
		DrawRect(sh)
	case Triangle:
		DrawTri(sh)
	}
```

## 4.2.3 Error Handling
