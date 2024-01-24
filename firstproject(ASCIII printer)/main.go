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

	// Получаем случайное слово для отображения в ASCII-арт
	word := getRandomWord()

	// Создаем объект Figlet
	asciiRender := figlet4go.NewAsciiRender()

	// Настройки для ASCII-арта
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

	// Выводим ASCII-арт
	fmt.Println(renderStr)
}

func getRandomWord() string {
	// Пример списка слов для ASCII-арта
	words := []string{"Gopher", "Hello", "Go", "Art", "ASCII", "Dangerousis"}

	// Генерируем случайный индекс из списка слов
	randomIndex := rand.Intn(len(words))

	// Возвращаем случайное слово
	return words[randomIndex]
}
