package note

import (
	"encoding/json"
	"errors"
	"os"
)

type note struct {
	Title   string
	Content string
}

func New(title string, content string) note {
	newNote := note{
		Title:   title,
		Content: content,
	}

	return newNote
}

func (n note) SaveToFile() (string, error) {
	jsonData, err := json.Marshal(n)

	if err != nil {
		return "", errors.New("Something went wrong")
	}

	os.WriteFile("example.json", []byte(jsonData), 0644)

	return "File saved", nil
}