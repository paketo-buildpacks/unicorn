package main

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

func (p GemfileParser) Parse(path string) (bool, bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false, nil
		}

		return false, false, fmt.Errorf("failed to parse Gemfile: %w", err)
	}
	defer file.Close()

	quotes := `["']`
	mriRe := regexp.MustCompile(`^ruby .*`)
	unicornRe := regexp.MustCompile(fmt.Sprintf(`^gem %sunicorn%s`, quotes, quotes))

	hasMri := false
	hasUnicorn := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())

		if !hasMri {
			hasMri = mriRe.Match(line)
		}

		if !hasUnicorn {
			hasUnicorn = unicornRe.Match(line)
		}

		if hasMri && hasUnicorn {
			return true, true, nil
		}
	}

	return hasMri, hasUnicorn, nil
}
