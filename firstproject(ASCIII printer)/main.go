package main

import (
	"fmt"
	"github.com/mbndr/figlet4go"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	word := getRandomWord()
	asciiRender := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorRed,
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorBlue,
		figlet4go.ColorMagenta,
		figlet4go.ColorCyan,
		figlet4go.ColorWhite,
	}

	renderStr, err := asciiRender.RenderOpts(word, options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(renderStr)
}

func getRandomWord() string {
	words := []string{"Gopher", "Hello", "Go", "Art", "ASCII", "Dangerousis"}

	randomIndex := rand.Intn(len(words))

	return words[randomIndex]
}
