package calculadora_go

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) Operate(operacion string, operador string) int {

	valores := strings.Split(operacion, operador)
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	numero1 := parsear(valores[0])
	numero2 := parsear(valores[1])
	total := 0
	switch operador {
	case "+":
		total = numero1 + numero2
		fmt.Println(numero1 + numero2)
	case "-":
		total = numero1 - numero2
		fmt.Println(numero1 - numero2)
	case "*":
		total = numero1 * numero2
		fmt.Println(numero1 * numero2)
	default:
		fmt.Println("no existe el operador")
	}

	return total

}

func parsear(entrada string) int {
	numero, err := strconv.Atoi(entrada)
	if err != nil {
		fmt.Println("Error: format: int '<operator>' int")
		os.Exit(1)
	}
	return numero
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite una operacion")
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(">>>>>>>>>", operacion)
	return operacion
}
