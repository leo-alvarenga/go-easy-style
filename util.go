package gostyle

func generateANSISequence(txt, bg string, style ...string) string {

	t, b, s := "", "", ""

	for key := range ansiColors {
		if txt == key {
			t = ansiColors["txt"] + ansiColors[key] + ansiColors["separator"]
		}

		if bg == key {
			b = ansiColors["bg"] + ansiColors[key] + ansiColors["separator"]
		}
	}

	for _, st := range style {
		for key := range ansiColors {
			if st == key {
				s += ansiColors[key] + ansiColors["separator"]
			}
		}
	}

	if len(s) > 0 {
		s = s[:len(s)-1]
	} else if len(b) > 0 {
		b = b[:len(b)-1]
	} else if len(t) > 0 {
		t = t[:len(t)-1]
	} else {
		return ""
	}

	return ansiColors["escape"] + t + b + s + ansiColors["end"]
}

/*
	Formats and displays the strings passed as arguments as an error message

Depending on your use case, DO NOT use this function, as it uses the buil-in
'print' and 'println' functions to display the text (not particularly thread-safe)
*/
func ShowAsError(title, msg string) {
	s1, s2 := new(TextStyle), new(TextStyle)

	s1.New("red", "black", "bold", "underline")
	s2.New("red", "black", "bold")

	s1.ShowWithStyle(title)
	s2.ShowWithStyle(msg)
}
