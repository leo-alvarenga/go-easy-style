package gostyle

type TextStyle struct {
	Text       string
	Background string
	Format     []string
	ANSI       string
}

func (t *TextStyle) New(txt, background string, styles ...string) {

	t.Text += txt
	t.Background = background
	t.Format = styles

	t.ANSI = generateANSISequence(txt, background, styles...)
}

func (t *TextStyle) Style(s string) string {
	return t.ANSI + s + reset
}

/*
	Displays the string passed as argument in the using the current style

Depending on your use case, DO NOT use this function, as it uses the buil-in
'print' and 'println' functions to display the text (not particularly thread-safe)
*/
func (t *TextStyle) ShowWithStyle(s string) {
	print(t.ANSI + s)
	println(reset)
}
