package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(hexToDecimal("1E"))
	fmt.Println(binToDecimal("10"))
	fmt.Println(capitalize("hELLo"))
	fmt.Println(capitalizeTheNth([]string{"go", "is", "really", "fun"}, 3))
	fmt.Println(joinWithPunctuation([]string{"hello", ",", "world", "!"}))
	fmt.Println(isPunctuation("!"))
	fmt.Println(isPunctuation("x"))
	fmt.Println(aOrAn("apple"))
	fmt.Println(fixArticles("There it was. A amazing rock. A honest man. A book."))

	result := fixSingleQuotes("' hello world '")
	fmt.Printf("%q\n", result)
}

// 1
func hexToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

// 2
func binToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}

// 3
func capitalize(w string) string {
	return strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
}

// 4
func capitalizeTheNth(words []string, n int) []string {
	for i := len(words) - n; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

// 5
func joinWithPunctuation(tokens []string) string {
	return strings.ReplaceAll(strings.Join(tokens, ""), ",", ", ")
}

// 6
func isPunctuation(s string) bool {
	return strings.Contains(",!", s)
}

// 7
func aOrAn(word string) string {
	if word == "" {
		return "a"
	}
	if strings.ContainsAny(strings.ToLower(string(word[0])), "aeiouh") {
		return "an"
	}
	return "a"
}

// 8
func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "A" || words[i] == "a" {
			next := strings.Trim(words[i+1], ".,!?")
			article := aOrAn(next)

			if words[i] == "A" {
				article = strings.Title(article)
			}
			words[i] = article
		}
	}
	return strings.Join(words, " ")
}

// 9
func fixSingleQuotes(text string) string {
	return "'" + strings.TrimSpace(strings.Trim(text, "'")) + "'"
}
