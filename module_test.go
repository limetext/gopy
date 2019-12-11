package gopy

import "testing"

func TestFunction(t *testing.T) {
	Initialize()
	defer Finalize()
	called := false
	f := func() (Object, error) {
		called = true
		return None, nil
	}
	if m, err := InitModule("mytest", []Method{{"mytest", f, ""}}); err != nil {
		t.Fatal(err)
	} else if t2, err := m.Dict().GetItemString("mytest"); err != nil {
		t.Fatal(err)
	} else {
		t2.Base().CallObject(nil)
	}
	if !called {
		t.Error("Function wasn't called")
	}
}

type ExampleClass struct {
	BaseObject
	called bool
}

func (e *ExampleClass) Py_Test() (Object, error) {
	panic("called")
}

func (e *ExampleClass) Py_Test2(args *Tuple, kwds *Dict) (Object, error) {
	if v, err := args.GetItem(0); err != nil {
		panic(err)
	} else if i, ok := v.(*Long); !ok {
		panic(v)
	} else if i.Int64() != 10 {
		panic(i)
	}
	panic("called2")
}

func (e *ExampleClass) PyStr() string {
	panic("strcalled")
}

var exampleClass = Class{
	Name:    "mytest.mytest",
	Pointer: &ExampleClass{},
}

func TestMethod(t *testing.T) {
	Initialize()
	defer Finalize()

	if main, err := NewDict(); err != nil {
		t.Fatal(err)
	} else if m, err := InitModule("mytest", nil); err != nil {
		t.Fatal(err)
	} else if c, err := exampleClass.Create(); err != nil {
		t.Fatal(err)
	} else if g, err := GetBuiltins(); err != nil {
		t.Fatal(err)
	} else if err := main.SetItemString("__builtins__", g); err != nil {
		t.Fatal(err)
	} else if err := m.AddObject("mytest", c); err != nil {
		t.Fatal(err)
		// } else if err := main.SetItemString("mytest", m); err != nil {
		// 	t.Fatal(err)
	} else if _, err := RunString("import mytest; a = mytest.mytest()", SingleInput, main, nil); err != nil {
		t.Fatal(err)
	} else if a, err := main.GetItemString("a"); err != nil {
		t.Fatal(err)
	} else if a == None || a.Type().String() != "<class 'mytest.mytest'>" {
		t.Error(a.Type().String())
	}
}

func TestMethod2(t *testing.T) {
	Initialize()
	defer Finalize()
	if main, err := NewDict(); err != nil {
		t.Fatal(err)
	} else if m, err := InitModule("mytest", nil); err != nil {
		t.Fatal(err)
	} else if c, err := exampleClass.Create(); err != nil {
		t.Fatal(err)
	} else if g, err := GetBuiltins(); err != nil {
		t.Fatal(err)
	} else if err := main.SetItemString("__builtins__", g); err != nil {
		t.Fatal(err)
	} else if err := m.AddObject("mytest", c); err != nil {
		t.Fatal(err)
	} else if _, err := RunString("import mytest; a = mytest.mytest()", SingleInput, main, nil); err != nil {
		t.Fatal(err)
	} else if a, err := main.GetItemString("a"); err != nil {
		t.Fatal(err)
	} else {
		type Test struct {
			m    string
			pan  string
			args []Object
		}
		tests := []Test{
			{"Test", "called", nil},
			{"Test2", "called2", []Object{NewLong(10)}},
			{"__str__", "strcalled", nil},
		}
		for _, test := range tests {
			func() {
				defer func() {
					if i := recover(); i == test.pan {
						t.Log("Success!")
					} else {
						t.Error("Panicked for some other reason:", i)
					}
				}()
				a.Base().CallMethodObjArgs(test.m, test.args...)
			}()
		}
	}
}
