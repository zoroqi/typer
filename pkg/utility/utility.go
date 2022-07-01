package utility

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/tyler-smith/go-bip39/wordlists"
)

// ReadFile returns the file contents as a string
func ReadFile(path string) (string, error) {
	contents, err := os.ReadFile(path)
	return string(contents), err
}

// RandomWords generates a strings with the specified number of random words using wordlist
func RandomWords(n int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	words := wordlists.English
	l := len(words)

	var s string
	for i := 0; i < n; i++ {
		s += words[r.Intn(l)] + " "
	}
	return s
}

// AdjustLength shortens a string if it's word count is greater than n
func AdjustLength(s string, n int) string {
	words := strings.Fields(s)
	if len(words) > n {
		words = words[:n]
		s2 := strings.Join(words, " ")
		l := len(s2)
		i := 0
		lastWorld := words[n-1]
		for i < l && i < len(s) {
			fmt.Println(i, strings.Index(s[i:], lastWorld))
			i = i + strings.Index(s[i:], lastWorld) + len(lastWorld)
		}
		if i >= len(s) {
			return s
		}
		return s[:i]
	} else {
		return s
	}
}

func AdjustTrimLine(s string) string {
	b := bufio.NewReader(strings.NewReader(s))
	sb := strings.Builder{}
	for {
		l, _, err := b.ReadLine()
		if err != nil {
			break
		}
		line := strings.TrimSpace(string(l))
		if line == "" {
			continue
		}
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	return strings.TrimSpace(sb.String())
}

// AdjustWhitespace replaces every group of whitespace characters with a single space charracter
func AdjustWhitespace(s string) (string, error) {
	reg := regexp.MustCompile(`[ \t\r]+`)
	s = reg.ReplaceAllString(s, " ")
	reg2 := regexp.MustCompile(`\n+`)
	s = reg2.ReplaceAllString(s, "\n")
	return strings.TrimSpace(s), nil
}

// RemoveNonAlpha removes all non-alphanumeric characters exept whitespace
func RemoveNonAlpha(s string) (string, error) {
	reg, err := regexp.Compile(`[^\p{L}\p{N} ]+`)
	if err != nil {
		return "", err
	}

	s = reg.ReplaceAllString(s, "")
	return s, nil
}

// Remove words of minimum length
func MinWordLength(s string, l int) (string, error) {
	words := strings.Fields(s)
	for i, word := range words {
		if len([]rune(word)) < l {
			words[i] = ""
		}
	}
	return strings.Join(words, " "), nil
}
