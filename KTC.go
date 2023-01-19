package main
import "fmt"

const ebulicaoK = 373.15

func main() {

	tempK := ebulicaoK // varialvel do valor da temperatura em kelvin
	tempC := tempK - 273.15 // variavel do valor da temperatura em celsius
	
	fmt.Printf("A temperatura de ebulição da água em °K é de %g ou em °C %.4g", tempK, tempC)
	
}