package gopy

import "testing"

func TestLong(t *testing.T) {
	Initialize()
	defer Finalize()
	l := NewLong(10)
	d, err := NewDict()
	if err != nil {
		t.Fatal(err)
	}
	if err = d.SetItemString("gopy.count", l); err != nil {
		t.Fatal(err)
	}
	if l2, err := d.GetItemString("gopy.count"); err != nil {
		t.Fatal(err)
	} else {
		if l3, ok := l2.(*Long); !ok {
			t.Fatalf("Unable to cast to long")
		} else {
			if v := l3.Int64(); v != 10 {
				t.Fatalf("Bad value: %d", v)
			}
		}
	}
}
