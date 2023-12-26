package shared

import (
	"bufio"
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile reads puzzle input into a string slice
func ReadFile(path string) []string {
	f, err := os.Open(path)
	Check(err)

	lines := make([]string, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	err = f.Close()
	Check(err)

	return lines
}

// RuneToInt converts rune to its numeric value
func RuneToInt(r rune) (int, error) {
	if r < '0' || r > '9' {
		return 0, fmt.Errorf("rune is non-numeric")
	}

	return int(r - '0'), nil
}
