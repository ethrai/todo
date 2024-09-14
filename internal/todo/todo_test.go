package todo

import "testing"

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
