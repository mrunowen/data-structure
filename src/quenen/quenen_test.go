package quenen

import (
	"fmt"
	"testing"
)

func TestQuenen(t *testing.T) {
	fmt.Println("Test Quenen...")
	fmt.Println("new quenen")

	quenen := New()
	size := quenen.GetLen()
	if size != 0 {
		t.Fatal("func GetLen() error")
	}

	e := &Elem{0}
	quenen.Enquenen(*e)
	size = quenen.GetLen()
	if size != 1 {
		t.Fatal("func GetLen() error")
	}

	e.Value = 1
	quenen.Enquenen(*e)
	size = quenen.GetLen()
	if size != 2 {
		t.Fatal("func GetLen() error")
	}

	e = quenen.Dequenen()
	if e.Value != 1 {
		t.Fatal("func Dequenen() error")
	}
	size = quenen.GetLen()
	if size != 1 {
		t.Fatal("func GetLen() error")
	}

	e = quenen.Dequenen()
	if e.Value != 0 {
		t.Fatal("func Dequenen() error")
	}
	size = quenen.GetLen()
	if size != 0 {
		t.Fatal("func GetLen() error")
	}

	e = quenen.Dequenen()
	if e != nil {
		t.Fatal("func Dequenen() error")
	}
}
