package ascii

import (
    "strings"
)

// AsciiArt creates ASCII art for the given words based on the contents provided.
func AsciiArt(words, contents []string) string {
	var countSpace int
	var artBuild strings.Builder // Using strings.Builder for efficient string concatenation.

	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ { // Each character is represented by 8 lines of ASCII art.
				for _, char := range word {
					if char == '\n' {
						continue
					}
					// Ensure the character is a printable ASCII character.
					if !(char >= 32 && char <= 126) {
						return "Error: Non printable ASCII character\n"
					}
					// Append the corresponding ASCII art line for the character.
					artBuild.WriteString(contents[int(char-' ')*9+1+i])
				}
				artBuild.WriteRune('\n')
			}
		} else { // Handle spaces between words.
			countSpace++
			if countSpace < len(words) {
				artBuild.WriteRune('\n')
			}
		}
	}
	return artBuild.String() // Return the constructed ASCII art as a string.
}