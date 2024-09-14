package todo

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	taskName := "Foo"
	l := New("test.json")
	l.Add(taskName)
	if len(l.tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(l.tasks))
	}
	if l.tasks[0].Name != taskName {
		t.Errorf("expected task[0].Name to be %s, got %s", taskName, l.tasks[0].Name)
	}
	if l.tasks[0].ID != 1 {
		t.Errorf("expected tasks[0].ID to be 1, got %d", l.tasks[0].ID)
	}

	taskName = "Bar"
	l.Add(taskName)

	if len(l.tasks) != 2 {
		t.Errorf("expected 2 task, got %d", len(l.tasks))
	}
	if l.tasks[1].Name != taskName {
		t.Errorf("expected task[1].Name to be %s, got %s", taskName, l.tasks[1].Name)
	}
	if l.tasks[1].ID != 2 {
		t.Errorf("expected tasks[1].ID to be 2, got %d", l.tasks[1].ID)
	}
}

func TestRemove(t *testing.T) {
	task1 := "Foo"
	task2 := "Bar"
	l := New("test.json")
	l.Add(task1)
	l.Add(task2)

	if len(l.tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(l.tasks))
	}

	// Act
	if err := l.Remove(1); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if len(l.tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(l.tasks))
	}

	if l.tasks[0].Name != task2 {
		t.Errorf("expected task[0].Name to be %s, got %s", task2, l.tasks[0].Name)
	}

	if err := l.Remove(2); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if len(l.tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(l.tasks))
	}
}

func TestDoneUndo(t *testing.T) {
	task1 := "Foo"
	task2 := "Bar"
	l := New("test.json")
	l.Add(task1)
	l.Add(task2)

	if len(l.tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(l.tasks))
	}

	// Act
	if err := l.Done(1); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if l.tasks[0].Done != true {
		t.Errorf("expected task[0].Done to be true, got %t", l.tasks[0].Done)
	}

	if err := l.Undo(1); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if l.tasks[0].Done != false {
		t.Errorf("expected task[0].Done to be false, got %t", l.tasks[0].Done)
	}
}

func TestSaveLoad(t *testing.T) {
	f, err := os.CreateTemp("", "test")
	if err != nil {
		t.Errorf("failed to create temp file: %s", err)
	}
	defer os.Remove(f.Name())

	task1 := "Foo"
	task2 := "Bar"
	l := New(f.Name())
	l.Add(task1)
	l.Add(task2)

	if err := l.Save(); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	l2 := New(f.Name())
	if err := l2.Load(); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if len(l2.tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(l2.tasks))
	}

	if l2.tasks[0].Name != task1 {
		t.Errorf("expected task[0].Name to be %s, got %s", task1, l2.tasks[0].Name)
	}

	if l2.tasks[1].Name != task2 {
		t.Errorf("expected task[1].Name to be %s, got %s", task2, l2.tasks[1].Name)
	}
}
