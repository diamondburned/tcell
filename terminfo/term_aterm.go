// Generated automatically.  DO NOT HAND-EDIT.

package terminfo

func init() {
	// AfterStep terminal
	AddTerminfo(&Terminfo{
<<<<<<< HEAD
		Name:         "aterm",
		Columns:      80,
		Lines:        24,
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[2J",
		EnterCA:      "\x1b7\x1b[?47h",
		ExitCA:       "\x1b[2J\x1b[?47l\x1b8",
		ShowCursor:   "\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[m\x0f",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		EnterKeypad:  "\x1b=",
		ExitKeypad:   "\x1b>",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		PadChar:      "\x00",
		AltChars:     "``aaffggjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b(B\x1b)0",
		Mouse:        "\x1b[M",
		MouseMode:    "%?%p1%{1}%=%t%'h'%Pa%e%'l'%Pa%;\x1b[?1000%ga%c\x1b[?1002%ga%c\x1b[?1003%ga%c\x1b[?1006%ga%c",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\u007f",
		KeyHome:      "\x1b[7~",
		KeyEnd:       "\x1b[8~",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1bOP",
		KeyF2:        "\x1bOQ",
		KeyF3:        "\x1bOR",
		KeyF4:        "\x1bOS",
		KeyF5:        "\x1b[15~",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF13:       "\x1b[25~",
		KeyF14:       "\x1b[26~",
		KeyF15:       "\x1b[28~",
		KeyF16:       "\x1b[29~",
		KeyF17:       "\x1b[31~",
		KeyF18:       "\x1b[32~",
		KeyF19:       "\x1b[33~",
		KeyF20:       "\x1b[34~",
		KeyF21:       "\x1b[23$",
		KeyF22:       "\x1b[24$",
		KeyF23:       "\x1b[11^",
		KeyF24:       "\x1b[12^",
		KeyF25:       "\x1b[13^",
		KeyF26:       "\x1b[14^",
		KeyF27:       "\x1b[15^",
		KeyF28:       "\x1b[17^",
		KeyF29:       "\x1b[18^",
		KeyF30:       "\x1b[19^",
		KeyF31:       "\x1b[20^",
		KeyF32:       "\x1b[21^",
		KeyF33:       "\x1b[23^",
		KeyF34:       "\x1b[24^",
		KeyF35:       "\x1b[25^",
		KeyF36:       "\x1b[26^",
		KeyF37:       "\x1b[28^",
		KeyF38:       "\x1b[29^",
		KeyF39:       "\x1b[31^",
		KeyF40:       "\x1b[32^",
		KeyF41:       "\x1b[33^",
		KeyF42:       "\x1b[34^",
		KeyF43:       "\x1b[23@",
		KeyF44:       "\x1b[24@",
		KeyBacktab:   "\x1b[Z",
		KeyShfLeft:   "\x1b[d",
		KeyShfRight:  "\x1b[c",
		KeyShfUp:     "\x1b[a",
		KeyShfDown:   "\x1b[b",
		KeyCtrlLeft:  "\x1b[Od",
		KeyCtrlRight: "\x1b[Oc",
		KeyCtrlUp:    "\x1b[Oa",
		KeyCtrlDown:  "\x1b[Ob",
		KeyShfHome:   "\x1b[7$",
		KeyShfEnd:    "\x1b[8$",
		KeyCtrlHome:  "\x1b[7^",
		KeyCtrlEnd:   "\x1b[8^",
=======
		Name:          "aterm",
		Columns:       80,
		Lines:         24,
		Colors:        8,
		Bell:          "\a",
		Clear:         "\x1b[H\x1b[2J",
		EnterCA:       "\x1b7\x1b[?47h",
		ExitCA:        "\x1b[2J\x1b[?47l\x1b8",
		ShowCursor:    "\x1b[?25h",
		HideCursor:    "\x1b[?25l",
		AttrOff:       "\x1b[m\x0f",
		Underline:     "\x1b[4m",
		Bold:          "\x1b[1m",
		Italic:        "\x1b[3m",
		Strikethrough: "\x1b[9m",
		Blink:         "\x1b[5m",
		Reverse:       "\x1b[7m",
		EnterKeypad:   "\x1b=",
		ExitKeypad:    "\x1b>",
		SetFg:         "\x1b[3%p1%dm",
		SetBg:         "\x1b[4%p1%dm",
		SetFgBg:       "\x1b[3%p1%d;4%p2%dm",
		PadChar:       "\x00",
		AltChars:      "``aaffggjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:      "\x0e",
		ExitAcs:       "\x0f",
		EnableAcs:     "\x1b(B\x1b)0",
		Mouse:         "\x1b[M",
		MouseMode:     "%?%p1%{1}%=%t%'h'%Pa%e%'l'%Pa%;\x1b[?1000%ga%c\x1b[?1002%ga%c\x1b[?1003%ga%c\x1b[?1006%ga%c",
		SetCursor:     "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:   "\b",
		CursorUp1:     "\x1b[A",
		KeyUp:         "\x1b[A",
		KeyDown:       "\x1b[B",
		KeyRight:      "\x1b[C",
		KeyLeft:       "\x1b[D",
		KeyInsert:     "\x1b[2~",
		KeyDelete:     "\x1b[3~",
		KeyBackspace:  "\xff",
		KeyHome:       "\x1b[7~",
		KeyEnd:        "\x1b[8~",
		KeyPgUp:       "\x1b[5~",
		KeyPgDn:       "\x1b[6~",
		KeyF1:         "\x1bOP",
		KeyF2:         "\x1bOQ",
		KeyF3:         "\x1bOR",
		KeyF4:         "\x1bOS",
		KeyF5:         "\x1b[15~",
		KeyF6:         "\x1b[17~",
		KeyF7:         "\x1b[18~",
		KeyF8:         "\x1b[19~",
		KeyF9:         "\x1b[20~",
		KeyF10:        "\x1b[21~",
		KeyF11:        "\x1b[23~",
		KeyF12:        "\x1b[24~",
		KeyF13:        "\x1b[25~",
		KeyF14:        "\x1b[26~",
		KeyF15:        "\x1b[28~",
		KeyF16:        "\x1b[29~",
		KeyF17:        "\x1b[31~",
		KeyF18:        "\x1b[32~",
		KeyF19:        "\x1b[33~",
		KeyF20:        "\x1b[34~",
		KeyF21:        "\x1b[23$",
		KeyF22:        "\x1b[24$",
		KeyF23:        "\x1b[11^",
		KeyF24:        "\x1b[12^",
		KeyF25:        "\x1b[13^",
		KeyF26:        "\x1b[14^",
		KeyF27:        "\x1b[15^",
		KeyF28:        "\x1b[17^",
		KeyF29:        "\x1b[18^",
		KeyF30:        "\x1b[19^",
		KeyF31:        "\x1b[20^",
		KeyF32:        "\x1b[21^",
		KeyF33:        "\x1b[23^",
		KeyF34:        "\x1b[24^",
		KeyF35:        "\x1b[25^",
		KeyF36:        "\x1b[26^",
		KeyF37:        "\x1b[28^",
		KeyF38:        "\x1b[29^",
		KeyF39:        "\x1b[31^",
		KeyF40:        "\x1b[32^",
		KeyF41:        "\x1b[33^",
		KeyF42:        "\x1b[34^",
		KeyF43:        "\x1b[23@",
		KeyF44:        "\x1b[24@",
		KeyBacktab:    "\x1b[Z",
		KeyShfLeft:    "\x1b[d",
		KeyShfRight:   "\x1b[c",
		KeyShfUp:      "\x1b[a",
		KeyShfDown:    "\x1b[b",
		KeyCtrlLeft:   "\x1b[Od",
		KeyCtrlRight:  "\x1b[Oc",
		KeyCtrlUp:     "\x1b[Oa",
		KeyCtrlDown:   "\x1b[Ob",
		KeyShfHome:    "\x1b[7$",
		KeyShfEnd:     "\x1b[8$",
		KeyCtrlHome:   "\x1b[7^",
		KeyCtrlEnd:    "\x1b[8^",
>>>>>>> 703b3f6... Hack in strikethrough and italic
	})
}
