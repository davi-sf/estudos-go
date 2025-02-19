package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func (note Note) ToString() {
	fmt.Println("Note title: ", note.Title)
	fmt.Println("---- Note content below ----")
	fmt.Println(note.Content)
	formatted := note.CreatedAt.Format("2006-01-02 15:04")
	fmt.Println("Created at: ", formatted)

}
