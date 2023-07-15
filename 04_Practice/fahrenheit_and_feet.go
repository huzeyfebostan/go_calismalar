package main

import "fmt"

func main() {
	var input, input2, output, output2 float32

	fmt.Print("Fahrenheit: ")
	fmt.Scanf("%f", &input)

	fmt.Print("Feet: ")
	fmt.Scanf("%f", &input2)

	output = (input - 32) * 5 / 9
	output2 = input2 * 0.3048

	fmt.Printf("Celsius: %.2f\nMeter: %.2f", output, output2)
}
