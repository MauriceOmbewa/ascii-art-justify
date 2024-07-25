package alignment

import (
    "fmt"
    "strings"
    "os"
	"justify/ascii"
)

func AlignLeft(text string, width int) string {
	return text
}

func AlignRight(text string, width int) string {
	var result strings.Builder
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		result.WriteString(fmt.Sprintf("%*s\n", width, line))
	}
	return result.String()
}

func AlignCenter(text string, width int) string {
	var result strings.Builder
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		padding := (width - len(line)) / 2
		result.WriteString(fmt.Sprintf("%*s%*s\n", padding+len(line), line, padding, ""))
	}
	return result.String()
}

func AlignJustify(input, banner string, width int) string {
    bannerFile := banner
    bannerText, err := os.ReadFile(bannerFile)
    if err != nil {
        fmt.Println("Error reading from file:", err)
        os.Exit(0)
    }
    fileInfo, err := os.Stat(bannerFile)
    if err != nil {
        fmt.Println("Error reading file information", err)
        os.Exit(0)
    }
    fileSize := fileInfo.Size()
    var contents []string
    if fileSize == 6623 || fileSize == 4703 || fileSize == 7463 || fileSize == 4496 {
        contents = strings.Split(string(bannerText), "\n")
    } else {
        return "Error: Invalid banner file size\n"
    }

    var result strings.Builder
    lines := strings.Split(input, "\n")
    
    for _, line := range lines {
        words := strings.Fields(line)
        if len(words) == 0 {
            result.WriteString("\n")
            continue
        }
        if len(words) == 1 {
            asciiWord := ascii.AsciiArt([]string{words[0]}, contents)
            result.WriteString(asciiWord)
            continue
        }

        asciiWords := make([]string, len(words))
        totalLength := 0
        for i, word := range words {
            asciiWords[i] = ascii.AsciiArt([]string{word}, contents)
            totalLength += len(strings.Split(asciiWords[i], "\n")[0])
        }

        spaceCount := len(words) - 1
        totalSpaces := width - totalLength
        spaces := make([]int, spaceCount)
        for i := range spaces {
            spaces[i] = totalSpaces / spaceCount
        }
        for i := 0; i < totalSpaces%spaceCount; i++ {
            spaces[i]++
        }

        asciiLines := make([][]string, 8)
        for i := range asciiLines {
            asciiLines[i] = make([]string, len(words))
        }

        for i, asciiWord := range asciiWords {
            lines := strings.Split(asciiWord, "\n")
            for j := 0; j < 8; j++ {
                asciiLines[j][i] = lines[j]
            }
        }

        for j := 0; j < 8; j++ {
            for i, word := range asciiLines[j] {
                result.WriteString(word)
                if i < len(asciiLines[j])-1 {
                    result.WriteString(strings.Repeat(" ", spaces[i]))
                }
            }
            result.WriteString("\n")
        }
    }

    return result.String()
}