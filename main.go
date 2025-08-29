package main

import (
	"bufio"
	"fmt"
	"os"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

	var hola int = 10

	fmt.Println(hola)

	var chau int = 50

	fmt.Println(chau)

	var arrayName = [4]string{"Hola", "Chau", "Chao", "Chao"}
	fmt.Println(arrayName)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful", "No"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	var caracter string = "b"
	fmt.Println(caracter)

	if caracter == "b" {
		fmt.Println("B")
	}

	if chau > hola {
		fmt.Println("Adios")
	}

	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	fruits := [3]string{"apple", "orange", "banana"}

	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	myMessage()

	fmt.Print("Presioná Enter para salir...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
