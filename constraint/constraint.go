package main

import "fmt"

type Addable interface {
	type int, float64, string
}

type Helloer interface {
	Hello() string
}

type Person struct {
	Name string
}

func (p Person) Hello() string {
	return "hello"
}


func sum[T Addable](a, b T) T {
	return a + b
}


func helloGeneric[T Helloer](a T) T {
	return a
}

func hello(a Helloer) Helloer {
	return a
}

func main() {
	fmt.Println(sum(10, 30)) // arguments must be of the same type
	fmt.Println(sum(1.5, 4.2))
	fmt.Println(sum("foo", "bar"))


	fmt.Println(helloGeneric(&Person{Name: "Dan"}).Name)

	// uncomment for erro
	fmt.Printf("%T", hello(&Person{Name: "Dan"}).Name)

	// uncomment for error
	// fmt.Println(sum("foo", 1))

}