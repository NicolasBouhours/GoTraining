package main

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)
	//colors["white"] = "#ffffff"
	//delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		println(color, hex)
	}
}
