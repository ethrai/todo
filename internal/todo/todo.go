package todo

import (
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
	store string
	tasks []task
}

// New creates a new todo list. It will treat file at storePath as a json store.
func New(storePath string) *List {
	return &List{
		store: storePath,
	}
}

// Add adds a new task to the list.
func (l *List) Add(name string) {
	l.tasks = append(l.tasks, task{
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
