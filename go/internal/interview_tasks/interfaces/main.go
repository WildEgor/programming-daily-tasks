package main

type Foo struct {
}

func (f *Foo) A() {
}

func (f *Foo) B() {
}

func (f *Foo) C() {
}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

func main() {
	var f AB = &Foo{}
	y := f.(BC) // ok
	_ = y
	// y.A() no

	if y, ok := f.(BC); ok {
		y.C() // it's ok
	}

	switch f.(type) {
	case AB:
		f.A()
		f.B()
	case BC:
		f.B()
		// f.C() is not allowed here because f's type is not guaranteed to be BC
	}
}
