package gostyle

import "testing"

func TestNew(t *testing.T) {
	a, b := new(TextStyle), new(TextStyle)

	a.New("black", "", "italic")
	b.New("black", "", "italic")

	if b.ANSI != a.ANSI {
		t.Error("*TextStyle.New() does not present the same behaviour when using the exact same arguments.")
	}

	b.New("white", "black")

	if b.ANSI == a.ANSI {
		t.Error("*TextStyle.New() presents the same behavior, even with different arguments.")
	}
}

func TestStyle(t *testing.T) {
	a := new(TextStyle)

	a.New("black", "", "italic")

	got := a.Style("test")
	expected := a.ANSI + "test" + reset

	if got != expected {
		t.Errorf("Expected %s; Got %s\n", expected, got)
	}
}

func TestGenerateANSISequence(t *testing.T) {
	a := generateANSISequence("black", "blue", "bold")
	b := generateANSISequence("white", "yellow", "bold", "italic", "underline", "strike")

	if a != "\033[30;44;1m" {
		t.Errorf(
			"Text 'black'; Background 'blue'; Style 'bold'"+
				"Expected '[30;44;1m'; Got %s\n",
			a[1:],
		)
	} else if b != "\033[37;43;1;3;4;9m" {
		t.Errorf(
			"Text 'white'; Background 'yellow'; Style 'bold', 'italic', 'underline', 'strike'"+
				"Expected '[37;43;1;3;4;9m'; Got %s\n",
			b[1:],
		)
	}
}
