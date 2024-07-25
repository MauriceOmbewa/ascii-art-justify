package main

import (
	"fmt"
	"os"
	"strings"

	"justify/alignment"
	"justify/ascii"
	"justify/terminal"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . --align=<alignment> <string> <banner>")
		os.Exit(1)
	}
	align := os.Args[1]
	input := os.Args[2]
	banner := os.Args[3]

	if strings.HasPrefix(align, "--align=") {
		align = strings.TrimPrefix(align, "--align=")
	} else {
		fmt.Println("Usage: go run . --align=<alignment> <string> <banner>")
		os.Exit(1)
	}

	if input == "" {
		fmt.Println("Usage: go run . --align=<alignment> <string> <banner>")
		os.Exit(1)
	}

	switch input {
	case "\\a", "\\0", "\\f", "\\v", "\\r":
		fmt.Println("Error: Non printable character", input)
		return
	}
	input = strings.ReplaceAll(input, "\\t", "    ")
	input = strings.ReplaceAll(input, "\\b", "\b")
	input = strings.ReplaceAll(input, "\\n", "\n")
	// Logic process for handling the backspace.
	for i := 0; i < len(input); i++ {
		indexB := strings.Index(input, "\b")
		if indexB > 0 {
			input = input[:indexB-1] + input[indexB+1:]
		}
	}
	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(input, "\n")

	// Check if the banner has an extension
	if strings.Contains(banner, ".") {
		// Check if the extension is not .txt
		if !strings.HasSuffix(banner, ".txt") {
			fmt.Println("Error: Required format: banner.txt")
			return
		}
	} else {
		// If no extension, add .txt
		banner = banner + ".txt"
	}
	// Convert to lowercase
	banner = strings.ToLower(banner)
	bannerFile := banner
	// Read the contents of banner file.
	bannerText, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	// Confirm file information.
	fileInfo, err := os.Stat(bannerFile)
	if err != nil {
		fmt.Println("Error reading file information", err)
		return
	}
	fileSize := fileInfo.Size()
	art := ""
	if fileSize == 6623 || fileSize == 4703 || fileSize == 7463 || fileSize == 4496 {
		// Split the content to a string slice and separate with newline.
		contents := strings.Split(string(bannerText), "\n")
		art = ascii.AsciiArt(words, contents)
	} else {
		fmt.Println("Error with the file size", fileSize)
		return
	}

	width, _, err := terminal.GetTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	var alignedText string

	switch align {
	case "left":
		alignedText = alignment.AlignLeft(art, width)
	case "right":
		alignedText = alignment.AlignRight(art, width)
	case "center":
		alignedText = alignment.AlignCenter(art, width)
	case "justify":
		alignedText = alignment.AlignJustify(input, banner, width)
	default:
		fmt.Println("Usage: go run . --align=<alignment> <string> <banner>")
		os.Exit(1)
	}

	banneredText := GetBanner(alignedText, banner)
	fmt.Println(banneredText)
}

func GetBanner(text, bannerType string) string {
	temp := ""
	switch bannerType {
	case "standard":
		temp = standardBanner(text)
		return temp
	case "shadow":
		temp = shadowBanner(text)
		return temp
	default:
		return text
	}
}

func standardBanner(text string) string {
	return text // Placeholder
}

func shadowBanner(text string) string {
	return text // Placeholder
}
