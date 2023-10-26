package main

import (
	"fmt"

	pilha "github.com/diogoX451/Pilha"
)

func main() {

	// array := []int{747, 63, 918, 884, 583, 818, 501, 68, 66, 131, 432, 710, 919, 399, 438, 436, 76, 311, 572, 104, 535, 920, 795, 709, 494, 135, 42, 834, 453, 865, 482, 571, 18, 864, 666, 55, 529, 127, 336, 657, 323, 450, 914, 854, 223, 412, 497, 156, 885, 786, 501, 628, 990, 283, 912, 824, 760, 136, 40, 763, 598, 642, 131, 578, 659, 915, 11, 647, 375, 35, 108, 13, 994, 419, 387, 872, 685, 320, 669, 897, 16, 274, 130, 657, 889, 207, 671, 578, 204, 26, 811, 197, 650, 135, 585, 329, 583, 619, 945, 375}

	// start := time.Now()

	// length := len(array)
	// ordenation.QuickSort(&array, 0, length-1)
	// end := time.Since(start)

	// fmt.Println("Demorou: ", end)

	// for i := 0; i < length; i++ {
	// 	fmt.Println(array[i])
	// }

	//Pilha

	pilha := pilha.NewPilha[string](5)

	pilha.Push("Pilha")
	pilha.Push("HuHu")
	pilha.Push("Stéphany")
	pilha.Push("Diogo")
	pilha.Push("Xavier")
	pilha.Pop()
	pilha.Push("Flavio")

	pilha.Push("Keyno")
	fmt.Println(pilha.List())
}
