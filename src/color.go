package src

import "math/rand"

func GetRandomColor() string {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"

	colors := []string{colorRed, colorGreen, colorYellow, colorBlue, colorPurple, colorCyan}

	return colors[rand.Intn(len(colors))]
}
