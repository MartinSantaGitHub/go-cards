package main

import "fmt"

func main() {
	var channel chan string
	channel2 := make(chan string)

	var slice []string
	slice2 := make([]string, 10)

	var colors map[string]string
	colors2 := make(map[string]string)

	fmt.Println(channel == nil)
	fmt.Println(channel2 == nil)

	fmt.Println()

	fmt.Println(slice == nil)
	fmt.Println(slice2 == nil)

	fmt.Println()

	fmt.Println(colors == nil)
	fmt.Println(colors2 == nil)

	// Not allowed, colors is nil
	// colors["white"] = "#ffffff"

	colors2["white"] = "#ffffff"

	delete(colors2, "white")

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	fmt.Println()

	fmt.Printf("%v\n", colors2)
}
