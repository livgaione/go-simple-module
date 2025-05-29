package main

import "fmt"

// func main() {

// 	for i := range 11 {

// 		fmt.Printf("%d\n", i)

// 	}
// }

// func main() {

// 	for i := 0; i <= 10; i++ {

// 		fmt.Printf("%d\n", i)

// 	}
// }

// func main() {

// 	for i := 10; i > 0; i-- {

// 		fmt.Printf("%d\n", i)

// 	}
// }

// 4. Tabuada do 5

// func main() {

// 	for i := 0; i <= 10; i++ {

// 		fmt.Printf("%d\n", i*5)

// 	}
// }

// func main() {

// 	for i := range 10 {

// 		fmt.Printf("%d\n", i*5)

// 	}
// }

// 5. Imprimir apenas números pares de 1 a 20

func main() {
	for i := range 20 {

		if i%2 == 0 {
			fmt.Printf("%d\n", i)

		}
	}
}
