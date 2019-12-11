package gopy

import "testing"

func TestRunString(t *testing.T) {
	Initialize()
	defer Finalize()
	if main, err := NewDict(); err != nil {
		t.Fatal(err)
	} else if g, err := GetBuiltins(); err != nil {
		t.Fatal(err)
	} else if err := main.SetItemString("__builtins__", g); err != nil {
		t.Fatal(err)
	} else if _, err := RunString("a = 'hello world!'", FileInput, main, nil); err != nil {
		t.Fatal(err)
	} else if a, err := main.GetItemString("a"); err != nil {
		t.Fatal(err)
	} else if b, ok := a.(*Unicode); !ok || b.String() != "hello world!" {
		t.Error(b, err)
	}
}
