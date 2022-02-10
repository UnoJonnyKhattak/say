package say_test

import (
	"fmt"
	"github.com/UnoJonnyKhattak/say"
	"testing"
)

func ExampleHello() {
	greeting := say.Hello("Jonny")
	fmt.Println(greeting)
}

func TestHello(t *testing.T) {
	n := "Jonny"
	expected := "Hello Jonny"
	actual := say.Hello(n)

	if expected != actual {
		t.Logf("Hello: expected [%s] got [%s]", expected, actual)
		t.Fail()
	}
}

func ExampleBye() {
	greeting := say.Bye("Jonny")
	fmt.Println(greeting)
}

func TestBye(t *testing.T) {
	n := "Jonny"
	expected := "Bye Jonny"
	actual := say.Bye(n)

	if expected != actual {
		t.Logf("Bye: expected [%s] got [%s]", expected, actual)
		t.Fail()
	}
}
