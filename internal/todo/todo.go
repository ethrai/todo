package todo

import (
	"io"
	"os"
	"time"
)

// task is a single todo item
type task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

// List is a todo list.
type List struct {
	store *os.File
	tasks []task
}

func New(store io.ReadWriter) *List {
	panic("not implemented")
}

// Add adds a new task to the list.
func (l *List) Add() {
	panic("not implemented")
}

// Remove removes a task from the list.
func (l *List) Remove(id int) error {
	panic("not implemented")
}

// Done marks a task as done.
func (l *List) Done(id int) error {
	panic("not implemented")
}

// Undo marks a task as undone.
func (l *List) Undo(id int) error {
	panic("not implemented")
}

// Save saves the list to the store.
func (l *List) Save() error {
	panic("not implemented")
}

// Load loads the list from the store.
func (l *List) Load() error {
	panic("not implemented")
}

// String returns all tasks as a string in formatted manner.
func (l *List) String() string {
	panic("not implemented")
}
