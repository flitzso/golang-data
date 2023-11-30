package main

import "fmt"

func main() {
	var student1 string = "John" // Declaração de variável com tipo explicitamente definido
	var student2 = "Jane"        // Tipo inferido pelo compilador
	x := 2                       // Declaração curta de variável com tipo inferido

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
