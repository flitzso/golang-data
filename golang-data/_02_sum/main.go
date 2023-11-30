package main

import "fmt"

func main() {
	var numero1, numero2 int

	// Solicitação do primeiro número
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&numero1)

	// Solicitação do segundo número
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&numero2)

	// Calcula a soma
	soma := numero1 + numero2

	// Exibe o resultado
	fmt.Printf("A soma de %d e %d é: %d\n", numero1, numero2, soma)
}
