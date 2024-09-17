package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// NotFoundError tells that task with ID is not present in the list.
type NotFoundError struct {
	ID int
}

// Error is impl of error interface
func (e NotFoundError) Error() string {
	return fmt.Sprintf("task %d not found", e.ID)
}

// Task is a single todo item
type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

// List is a todo list.
type List struct {
	storePath string
	tasks     []Task
}

// New creates a new todo list. It will treat file at storePath as a json store.
func New(storePath string) *List {
// New creates a new todo list. It will treat file at storeFile as a json store.
func New(storeFile string) *List {
	return &List{
		storePath: storeFile,
	}
}

// Add adds a new task to the list.
func (l *List) Add(name string) {
	l.tasks = append(l.tasks, Task{
		ID:          len(l.tasks) + 1,
		Name:        name,
		Desc:        "",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	})
}

// Remove removes a task from the list.
func (l *List) Remove(id int) error {
	for i, t := range l.tasks {
		if t.ID == id {
			l.tasks = append(l.tasks[:i], l.tasks[i+1:]...)
			return nil
		}
	}
	return NotFoundError{ID: id}
}

// Done marks a task as done.
func (l *List) Done(id int) error {
	for i, t := range l.tasks {
		if t.ID == id {
			l.tasks[i].Done = true
			l.tasks[i].CompletedAt = time.Now()
			return nil
		}
	}
	return NotFoundError{ID: id}
}

// Undo marks a task as undone.
func (l *List) Undo(id int) error {
	for i, t := range l.tasks {
		if t.ID == id {
			l.tasks[i].Done = false
			l.tasks[i].CompletedAt = time.Time{}
			return nil
		}
	}
	return NotFoundError{ID: id}
}

// Save saves the list to the store.
func (l *List) Save() error {
	f, err := os.OpenFile(l.storePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(l.tasks); err != nil {
		return fmt.Errorf("failed to encode tasks: %s", err)
	}

	return nil
}

// Load loads the list from the store.
func (l *List) Load() error {
	f, err := os.OpenFile(l.storePath, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&l.tasks); err != nil {
		return fmt.Errorf("failed to decode tasks: %s", err)
	}

	return nil
}

// String returns all tasks as a string in formatted manner.
func (l *List) String() string {
	panic("not implemented")
}
