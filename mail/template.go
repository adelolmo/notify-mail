package mail

import (
	"io/ioutil"
	"strings"
)

type Template struct {
	Map map[string]string
}

func NewTemplate(placeholders map[string]string) *Template {
	return &Template{Map: placeholders}
}

func (t *Template) ReplaceContent(filename string) (string, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(input), "\n")
	for i := range lines {
		for key, value := range t.Map {
			lines[i] = strings.Replace(lines[i], key, value, -1)
		}
	}
	return strings.Join(lines, "\n"), nil
}
