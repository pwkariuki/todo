package todo_test

import (
	"os"
	"testing"
	"todo"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.\n", taskName, l[0].Task)
	}
}

// TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.\n", taskName, l[0].Task)
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("new task should be completed.")
	}
}

// TestDelete tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	for i := 0; i < len(tasks); i++ {
		if l[i].Task != tasks[i] {
			t.Errorf("expected %q, got %q instead.\n", tasks[i], l[i].Task)
		}
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("expected list length %d, got %d instead.\n", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q instead.\n", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.\n", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("task %q should match %q task.\n", l1[0].Task, l2[0].Task)
	}
}
