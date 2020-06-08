package unicorn

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type GemfileParser struct{}

func NewGemfileParser() GemfileParser {
	return GemfileParser{}
}

func (p GemfileParser) Parse(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, fmt.Errorf("failed to parse Gemfile: %w", err)
	}
	defer file.Close()

	quotes := `["']`
	unicornRe := regexp.MustCompile(fmt.Sprintf(`^gem %sunicorn%s`, quotes, quotes))

	hasUnicorn := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())

		if !hasUnicorn {
			hasUnicorn = unicornRe.Match(line)
		}

		if hasUnicorn {
			return true, nil
		}
	}

	return false, nil
}
