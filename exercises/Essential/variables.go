package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

func main() {

	fmt.Println("Hola mundo")

	fmt.Println(quote.Hello())

	// TYPE OF DATA
	// Integer
	// int8; -128 => 127; 8bits
	// int16; -32,768 => 32,767
	// int32; -2,147,483,648 => 2,147,483,647
	// int64; -9,223,372,036,854,775,808 => 9,223,372,036,854,775,807

	// uint; Only use positive numbers
	// Exist uint8 => uint64

	// numbers with point decimal
	// float32 is mayor than int64
	// float64 is mayor than int64

	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
	fmt.Println(uint64(math.MaxUint64))

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fullName := "Josue Molina \t(alias \"roelcode\")\n"
	fmt.Println(fullName)

	//tipo de dato byte; se utiliza para representar carecteres de tipo ASCII
	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	// Es un dato de rune; representa caracteres unicode
	var valorUnicode rune = '😀'
	fmt.Println(valorUnicode)

	fmt.Println("Valores predeterminados")

	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	//Conversiones de datos

	var integer16 int16 = 50
	var integer32 int32 = 100

	println(integer16 + int16(integer32))

	numInText := "100"

	// Atoi => Ascii to Integer
	num, _ := strconv.Atoi(numInText)
	fmt.Println(num + int(integer16))
	number := 52
	// Itoa => Integer to Ascii
	numInText = strconv.Itoa(number)
	fmt.Println(numInText)

	var age int
	var name string

	age = 32
	name = "Josue"
	fmt.Printf("Hola, mi nombre es %s y tengo %d años", name, age)
	greeting := fmt.Sprintf("Hola, mi nombre es %s y tengo %d años", name, age)
	fmt.Println(greeting)

	// almacenar data desde el teclado
	// El carácter & en Go se usa para obtener la dirección de memoria de una variable, es decir, su puntero.
	fmt.Printf("Ingresa tu nombre:")
	fmt.Scan(&name)

	fmt.Printf("Ingresa tu edad:")
	fmt.Scan(&age)

	fmt.Printf("Hola %s, tu edad es de %d", name, age)
	fmt.Printf("Hola %s, tu edad es de %d", name, age)

	var base float32
	var sideOne float32
	var sideTwo float32
	var total float32
	fmt.Printf("Ingresa la medida del primer lado: ")
	fmt.Scan(&sideOne)

	fmt.Printf("Ingresa la medida del segundo lado: ")
	fmt.Scan(&sideTwo)

	fmt.Printf("Ingresa la medida de la base: ")
	fmt.Scan(&base)

	total += base + sideOne + sideTwo
	fmt.Printf("el perimetro del triangulo es %v", total)
}
