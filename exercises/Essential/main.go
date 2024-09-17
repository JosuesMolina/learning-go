package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		println("Buenos días")
	} else if t.Hour() < 17 {
		println("Buenas tardes")
	} else {
		println("Buenas noches")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Es la mañana")
	case t.Hour() >= 17:
		fmt.Println("Es la tarde.")
	default:
		fmt.Println("Es la noche")

	}

	switch os := runtime.GOOS; os {
	case "windows":
		println("Go is run in Windows")
	case "linux":
		println("Go is run in Linux")
	}

	var i int
	println("Example #1 for")
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	println("Example #2 for")
	for i := 1; i <= 10; i++ {

		if i%2 > 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("Arrays: Example #1")
	var matriz = [...]int{10, 20, 30, 40, 50}
	matriz[0] = 100
	matriz[1] = 200
	fmt.Println(matriz)

	fmt.Println("Arrays: Example #2")
	for i := 0; i < len(matriz); i++ {
		println(matriz[i])
	}

	fmt.Println("Arrays: Example #3")
	for index, value := range matriz {
		fmt.Printf("indice: %d , valor: %d \n", index, value)
	}

	fmt.Println("Arrays: Example #4")
	for _, value := range matriz {
		fmt.Printf("valor: %d \n", value)
	}

	var matrizBidimensional = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrizBidimensional)

	fmt.Println("SLices: Example #1")
	diaSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}

	diaSlice := diaSemana[0:5]
	fmt.Println(diaSlice)

	diaSlice = append(diaSlice, "otro día")
	fmt.Println(diaSlice)
	fmt.Println("Slices: Delete Example #2")
	diaSlice = append(diaSlice[:2], diaSlice[3:]...)
	fmt.Println(diaSlice)

	// longitud de slice
	fmt.Println(len(diaSlice))
	// capacidad del slice
	fmt.Println(cap(diaSlice))

	// Crea un slice
	nombres := make([]string, 5, 10)
	fmt.Println(nombres)

	nombres[0] = "Josue"
	nombres[1] = "Pedro"
	nombres[2] = "Pablo"
	nombres[3] = "Maria"
	nombres[4] = "Jose"
	fmt.Println(nombres)

	nombres = append(nombres, "Juan", "Eduardo", "Mario")
	fmt.Println(nombres)

	// longitud de slice
	fmt.Println(len(nombres))
	// capacidad del slice
	fmt.Println(cap(nombres))

	rebanada := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, len(rebanada))

	copy(rebanada2, rebanada)
	fmt.Println(rebanada)
	fmt.Println(rebanada2)

	fmt.Println("Mapas:Example #1")

	colores := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colores)
	colores["negro"] = "#000000"
	fmt.Println(colores)

	if valor, color := colores["morado"]; color {
		fmt.Println(valor)
		fmt.Println(color)
	} else {
		fmt.Println("El color no existe")
	}

	delete(colores, "verde")

	for valor, clave := range colores {
		fmt.Println(valor)
		fmt.Println(clave)
	}

	fmt.Println("Structs: Example #1")
	p := Persona{"Josue", 32, "correo@gmail.com"}
	p2 := Persona{"Juan", 385, "correo@hotmail.com"}

	fmt.Println(p, p2)

	fmt.Println("Punteros: Example #1")
	// Puntero: Variable que almacena la direccion de memomria de otra variable
	var x int = 10
	// se manda la referencia de memoria de x a puntero
	var puntero *int = &x

	fmt.Println(&x)
	fmt.Println(puntero)
	fmt.Println(x)

	x = 100
	editaSinPuntero(x)
	fmt.Println(x)
	editaPuntero(&x)
	fmt.Println(x)
	p.diHola()
	fmt.Println("")
	if !true {
		// EXAMPLE OF PUNTEROS
		// intancia para entrada de datos
		leer := bufio.NewReader(os.Stdin)
		exampleListaTareas(1, leer)
	}
	fmt.Println("Control de errores: Example #1")
	var res = badNumber()
	fmt.Println(res)

	result, err := divideZero(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		println(result)
	}

	fmt.Println("defer: Example #1")
	// Orden de ejecución inverso
	// Es como una pila: el último defer se ejecuta primero.
	// LIFO
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	// útil para asegurarte de que los recursos se cierren o liberen
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, soy Josue Molina"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("panic: Example #1")
	controlPanic(false)
	controlPanic(true)
	controlPanic(false)

	fmt.Println("Control de errores: Example #1")
	log.Print("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro")
	log.SetPrefix("main():")
	// log.Fatal("¡Oye, soy un registro de errores!")
	// log.Panic("¡Oye, soy un registro de errores!")
	files, errs := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if errs != nil {
		log.Fatal(errs)
	}

	defer file.Close()

	log.SetOutput(files)
	log.Print("¡Oye, soy un Log!")
}

func controlPanic(flag bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	createError(flag)
	fmt.Println("Todo OK")
}

func createError(flag bool) {
	if flag {
		panic("Se ha generado un error")
	}
}

func badNumber() int {
	var strNumber string = "123a"
	number, err := strconv.Atoi(strNumber)
	if err != nil {
		var e = errors.New("the string is not a valid number \n")
		fmt.Println("Error:", e)
	}
	return number
}

func divideZero(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func (p *Persona) diHola() {
	fmt.Print("Hola, Mi nombre es ", p.nombre)
}

func editaSinPuntero(x int) {
	x = 200
}

func editaPuntero(x *int) {
	*x = 20
}

func Hello(name string) string {
	return "Hello," + name
}

func calculadora(a, b int) (int, int) {
	sum := a + b
	res := a - b
	return sum, res
}

func calcula(a, b int) (sum, res int) {
	sum = a + b
	res = a - b
	return
}

type Persona struct {
	nombre string
	edad   int
	correo string
}

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Add a new task
func (l *ListaTareas) addTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// method to mask a task as complete
func (l *ListaTareas) markAsComplete(index int) {
	l.tareas[index].completado = true
}

func (l *ListaTareas) patchTarea(index int, t Tarea) {
	l.tareas[index] = t
}

func (l *ListaTareas) deleteTarea(index int) {
	// append(diaSlice[:2], diaSlice[3:]...)
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func exampleListaTareas(opcion int, leer *bufio.Reader) {
	var lista = ListaTareas{}
	for {
		fmt.Println("=================================================")
		fmt.Println(
			"Selecciona una opcion \n",
			"1. agregar tarea \n",
			"2. Marcar como completada \n",
			"3. Editar tarea \n",
			"4. Eliminar tarea \n",
			"5. Salir")

		fmt.Println("Elige la opcion:")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Descripcion de la tarea:")
			t.desc, _ = leer.ReadString('\n')
			lista.addTarea(t)
		case 2:
			fmt.Println("¿Cual es el indice de la tarea que quieres marcar como acompletada?")
			var index int
			fmt.Scanln(&index)
			lista.markAsComplete(index)
		case 3:
			fmt.Println("¿Cual es el indice de la tarea que quieres modificar como acompletada?")
			var index int
			fmt.Scan(&index)
			var t Tarea
			fmt.Println("Nuevo nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Nueva descripcion de la tarea")
			t.desc, _ = leer.ReadString('\n')
			t.completado = false
			lista.patchTarea(index, t)

		case 4:
			fmt.Println("¿Cual es el indice de la tarea que quieres modificar como acompletada?")
			var index int
			fmt.Scan(&index)
			lista.deleteTarea(index)
		case 5:
			return
		default:
			fmt.Println("opcion invalida")

		}

		fmt.Println("=========================")
		for i, value := range lista.tareas {
			fmt.Printf("Tarea #%d", i)
			fmt.Printf("\n - %s - %s - Esta completada: %t\n", value.nombre, value.desc, value.completado)
		}
		fmt.Println("=========================")

	}
}
