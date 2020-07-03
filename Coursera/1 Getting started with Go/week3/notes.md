## 3.1.1 Classes and Encapsulation
Go doesnot contain the "class" keyword like other languages

## 3.1.2 Support for classes
### Associating Methods with data

```go
type MyInt int
func (mi MyInt) Double () int {
    return int(mi * 2)
}
func main(){
        v := MyInt(3)
        fmt.Println(v.Double())
}
```
Here, v is *implicit argument* and is *passed by value*

## 3.2.1 Packages and Encapsulation
```go
package data
var x int = 1
func PrintX(){
	fmt.Println(x)
}
```

```go
package main
import "data"
func main(){
	data.PrintX()
}
```

## 3.2.2 Pointer Recievers
Methods called on data use *passing by value*. This could be costly and doesn't change the original object.

Receivers could be a pointer to the object
```go 
func (p *Point) OffsetX(v float64){
    p.x = p.x + v
}
```

## 3.2.2 Point recievers, referencing, dereferencing
### No need to Derefernce
p.x works **fine**. (Unlike p->x in C/C++)

```go 
func (p *Point) OffsetX(v float64){
    p.x = p.x + v
}
```

### No need to Reference
Even though the receiver for `Offset` is `*Point`, no need to reference on calling the method.

```go
p.Offset(5)
```


