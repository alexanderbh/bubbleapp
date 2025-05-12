// From: https://github.com/gabe565/go-spinners/blob/main/spinners.go
// Which is originally from: https://github.com/sindresorhus/cli-spinners

package loader

import "time"

var (
	Dots = Spinner{
		Frames:   []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		Interval: 100 * time.Millisecond,
	}

	Dots2 = Spinner{
		Frames:   []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		Interval: 100 * time.Millisecond,
	}

	Dots3 = Spinner{
		Frames:   []string{"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠴", "⠲", "⠳", "⠓"},
		Interval: 100 * time.Millisecond,
	}

	Dots4 = Spinner{
		Frames:   []string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"},
		Interval: 100 * time.Millisecond,
	}

	Dots5 = Spinner{
		Frames:   []string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
		Interval: 100 * time.Millisecond,
	}

	Dots6 = Spinner{
		Frames:   []string{"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
		Interval: 100 * time.Millisecond,
	}

	Dots7 = Spinner{
		Frames:   []string{"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
		Interval: 100 * time.Millisecond,
	}

	Dots8 = Spinner{
		Frames:   []string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
		Interval: 100 * time.Millisecond,
	}

	Dots9 = Spinner{
		Frames:   []string{"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
		Interval: 100 * time.Millisecond,
	}

	Dots10 = Spinner{
		Frames:   []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		Interval: 100 * time.Millisecond,
	}

	Dots11 = Spinner{
		Frames:   []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
		Interval: 100 * time.Millisecond,
	}

	Dots12 = Spinner{
		Frames:   []string{"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀"},
		Interval: 100 * time.Millisecond,
	}

	Dots13 = Spinner{
		Frames:   []string{"⣼", "⣹", "⢻", "⠿", "⡟", "⣏", "⣧", "⣶"},
		Interval: 100 * time.Millisecond,
	}

	Dots14 = Spinner{
		Frames:   []string{"⠉⠉", "⠈⠙", "⠀⠹", "⠀⢸", "⠀⣰", "⢀⣠", "⣀⣀", "⣄⡀", "⣆⠀", "⡇⠀", "⠏⠀", "⠋⠁"},
		Interval: 100 * time.Millisecond,
	}

	Dots8Bit = Spinner{
		Frames:   []string{"⠀", "⠁", "⠂", "⠃", "⠄", "⠅", "⠆", "⠇", "⡀", "⡁", "⡂", "⡃", "⡄", "⡅", "⡆", "⡇", "⠈", "⠉", "⠊", "⠋", "⠌", "⠍", "⠎", "⠏", "⡈", "⡉", "⡊", "⡋", "⡌", "⡍", "⡎", "⡏", "⠐", "⠑", "⠒", "⠓", "⠔", "⠕", "⠖", "⠗", "⡐", "⡑", "⡒", "⡓", "⡔", "⡕", "⡖", "⡗", "⠘", "⠙", "⠚", "⠛", "⠜", "⠝", "⠞", "⠟", "⡘", "⡙", "⡚", "⡛", "⡜", "⡝", "⡞", "⡟", "⠠", "⠡", "⠢", "⠣", "⠤", "⠥", "⠦", "⠧", "⡠", "⡡", "⡢", "⡣", "⡤", "⡥", "⡦", "⡧", "⠨", "⠩", "⠪", "⠫", "⠬", "⠭", "⠮", "⠯", "⡨", "⡩", "⡪", "⡫", "⡬", "⡭", "⡮", "⡯", "⠰", "⠱", "⠲", "⠳", "⠴", "⠵", "⠶", "⠷", "⡰", "⡱", "⡲", "⡳", "⡴", "⡵", "⡶", "⡷", "⠸", "⠹", "⠺", "⠻", "⠼", "⠽", "⠾", "⠿", "⡸", "⡹", "⡺", "⡻", "⡼", "⡽", "⡾", "⡿", "⢀", "⢁", "⢂", "⢃", "⢄", "⢅", "⢆", "⢇", "⣀", "⣁", "⣂", "⣃", "⣄", "⣅", "⣆", "⣇", "⢈", "⢉", "⢊", "⢋", "⢌", "⢍", "⢎", "⢏", "⣈", "⣉", "⣊", "⣋", "⣌", "⣍", "⣎", "⣏", "⢐", "⢑", "⢒", "⢓", "⢔", "⢕", "⢖", "⢗", "⣐", "⣑", "⣒", "⣓", "⣔", "⣕", "⣖", "⣗", "⢘", "⢙", "⢚", "⢛", "⢜", "⢝", "⢞", "⢟", "⣘", "⣙", "⣚", "⣛", "⣜", "⣝", "⣞", "⣟", "⢠", "⢡", "⢢", "⢣", "⢤", "⢥", "⢦", "⢧", "⣠", "⣡", "⣢", "⣣", "⣤", "⣥", "⣦", "⣧", "⢨", "⢩", "⢪", "⢫", "⢬", "⢭", "⢮", "⢯", "⣨", "⣩", "⣪", "⣫", "⣬", "⣭", "⣮", "⣯", "⢰", "⢱", "⢲", "⢳", "⢴", "⢵", "⢶", "⢷", "⣰", "⣱", "⣲", "⣳", "⣴", "⣵", "⣶", "⣷", "⢸", "⢹", "⢺", "⢻", "⢼", "⢽", "⢾", "⢿", "⣸", "⣹", "⣺", "⣻", "⣼", "⣽", "⣾", "⣿"},
		Interval: 100 * time.Millisecond,
	}

	DotsCircle = Spinner{
		Frames:   []string{"⢎ ", "⠎⠁", "⠊⠑", "⠈⠱", " ⡱", "⢀⡰", "⢄⡠", "⢆⡀"},
		Interval: 100 * time.Millisecond,
	}

	Sand = Spinner{
		Frames:   []string{"⠁", "⠂", "⠄", "⡀", "⡈", "⡐", "⡠", "⣀", "⣁", "⣂", "⣄", "⣌", "⣔", "⣤", "⣥", "⣦", "⣮", "⣶", "⣷", "⣿", "⡿", "⠿", "⢟", "⠟", "⡛", "⠛", "⠫", "⢋", "⠋", "⠍", "⡉", "⠉", "⠑", "⠡", "⢁"},
		Interval: 100 * time.Millisecond,
	}

	Line = Spinner{
		Frames:   []string{"-", "\\", "|", "/"},
		Interval: 100 * time.Millisecond,
	}

	Line2 = Spinner{
		Frames:   []string{"⠂", "-", "–", "—", "–", "-"},
		Interval: 100 * time.Millisecond,
	}

	Pipe = Spinner{
		Frames:   []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
		Interval: 100 * time.Millisecond,
	}

	SimpleDots = Spinner{
		Frames:   []string{".  ", ".. ", "...", "   "},
		Interval: 400 * time.Millisecond,
	}

	SimpleDotsScrolling = Spinner{
		Frames:   []string{".  ", ".. ", "...", " ..", "  .", "   "},
		Interval: 200 * time.Millisecond,
	}

	Star = Spinner{
		Frames:   []string{"✶", "✸", "✹", "✺", "✹", "✷"},
		Interval: 100 * time.Millisecond,
	}

	Star2 = Spinner{
		Frames:   []string{"+", "x", "*"},
		Interval: 100 * time.Millisecond,
	}

	Flip = Spinner{
		Frames:   []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"},
		Interval: 100 * time.Millisecond,
	}

	Hamburger = Spinner{
		Frames:   []string{"☱", "☲", "☴"},
		Interval: 100 * time.Millisecond,
	}

	GrowVertical = Spinner{
		Frames:   []string{" ", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"},
		Interval: 100 * time.Millisecond,
	}

	GrowHorizontal = Spinner{
		Frames:   []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "▊", "▋", "▌", "▍", "▎"},
		Interval: 100 * time.Millisecond,
	}

	Balloon = Spinner{
		Frames:   []string{" ", ".", "o", "O", "@", "*", " "},
		Interval: 100 * time.Millisecond,
	}

	Balloon2 = Spinner{
		Frames:   []string{".", "o", "O", "°", "O", "o", "."},
		Interval: 100 * time.Millisecond,
	}

	Noise = Spinner{
		Frames:   []string{"▓", "▒", "░", "▒"},
		Interval: 200 * time.Millisecond,
	}

	Bounce = Spinner{
		Frames:   []string{"⠁", "⠂", "⠄", "⠂"},
		Interval: 100 * time.Millisecond,
	}

	BoxBounce = Spinner{
		Frames:   []string{"▖", "▘", "▝", "▗"},
		Interval: 100 * time.Millisecond,
	}

	BoxBounce2 = Spinner{
		Frames:   []string{"▌", "▀", "▐", "▄"},
		Interval: 100 * time.Millisecond,
	}

	Triangle = Spinner{
		Frames:   []string{"◢", "◣", "◤", "◥"},
		Interval: 100 * time.Millisecond,
	}

	Binary = Spinner{
		Frames:   []string{"010010", "001100", "100101", "111010", "111101", "010111", "101011", "111000", "110011", "110101"},
		Interval: 100 * time.Millisecond,
	}

	Arc = Spinner{
		Frames:   []string{"◜", "◠", "◝", "◞", "◡", "◟"},
		Interval: 100 * time.Millisecond,
	}

	Circle = Spinner{
		Frames:   []string{"◡", "⊙", "◠"},
		Interval: 100 * time.Millisecond,
	}

	SquareCorners = Spinner{
		Frames:   []string{"◰", "◳", "◲", "◱"},
		Interval: 200 * time.Millisecond,
	}

	CircleQuarters = Spinner{
		Frames:   []string{"◴", "◷", "◶", "◵"},
		Interval: 100 * time.Millisecond,
	}

	CircleHalves = Spinner{
		Frames:   []string{"◐", "◓", "◑", "◒"},
		Interval: 100 * time.Millisecond,
	}

	Squish = Spinner{
		Frames:   []string{"╫", "╪"},
		Interval: 100 * time.Millisecond,
	}

	Toggle = Spinner{
		Frames:   []string{"⊶", "⊷"},
		Interval: 200 * time.Millisecond,
	}

	Toggle2 = Spinner{
		Frames:   []string{"▫", "▪"},
		Interval: 500 * time.Millisecond,
	}

	Toggle3 = Spinner{
		Frames:   []string{"□", "■"},
		Interval: 1000 * time.Millisecond,
	}

	Toggle4 = Spinner{
		Frames:   []string{"■", "□", "▪", "▫"},
		Interval: 200 * time.Millisecond,
	}

	Toggle5 = Spinner{
		Frames:   []string{"▮", "▯"},
		Interval: 300 * time.Millisecond,
	}

	Toggle6 = Spinner{
		Frames:   []string{"ဝ", "၀"},
		Interval: 300 * time.Millisecond,
	}

	Toggle7 = Spinner{
		Frames:   []string{"⦾", "⦿"},
		Interval: 200 * time.Millisecond,
	}

	Toggle8 = Spinner{
		Frames:   []string{"◍", "◌"},
		Interval: 200 * time.Millisecond,
	}

	Toggle9 = Spinner{
		Frames:   []string{"◉", "◎"},
		Interval: 200 * time.Millisecond,
	}

	Toggle10 = Spinner{
		Frames:   []string{"㊂", "㊀", "㊁"},
		Interval: 200 * time.Millisecond,
	}

	Toggle11 = Spinner{
		Frames:   []string{"⧇", "⧆"},
		Interval: 200 * time.Millisecond,
	}

	Toggle12 = Spinner{
		Frames:   []string{"☗", "☖"},
		Interval: 300 * time.Millisecond,
	}

	Toggle13 = Spinner{
		Frames:   []string{"=", "*", "-"},
		Interval: 100 * time.Millisecond,
	}

	Arrow = Spinner{
		Frames:   []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
		Interval: 100 * time.Millisecond,
	}

	Arrow2 = Spinner{
		Frames:   []string{"⬆️ ", "↗️ ", "➡️ ", "↘️ ", "⬇️ ", "↙️ ", "⬅️ ", "↖️ "},
		Interval: 100 * time.Millisecond,
	}

	Arrow3 = Spinner{
		Frames:   []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
		Interval: 100 * time.Millisecond,
	}

	BouncingBar = Spinner{
		Frames:   []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[====]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"},
		Interval: 100 * time.Millisecond,
	}

	BouncingBall = Spinner{
		Frames:   []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"},
		Interval: 100 * time.Millisecond,
	}

	Smiley = Spinner{
		Frames:   []string{"😄 ", "😝 "},
		Interval: 200 * time.Millisecond,
	}

	Monkey = Spinner{
		Frames:   []string{"🙈 ", "🙈 ", "🙉 ", "🙊 "},
		Interval: 300 * time.Millisecond,
	}

	Hearts = Spinner{
		Frames:   []string{"💛 ", "💙 ", "💜 ", "💚 ", "❤️ "},
		Interval: 100 * time.Millisecond,
	}

	Clock = Spinner{
		Frames:   []string{"🕛 ", "🕐 ", "🕑 ", "🕒 ", "🕓 ", "🕔 ", "🕕 ", "🕖 ", "🕗 ", "🕘 ", "🕙 ", "🕚 "},
		Interval: 100 * time.Millisecond,
	}

	Earth = Spinner{
		Frames:   []string{"🌍 ", "🌎 ", "🌏 "},
		Interval: 200 * time.Millisecond,
	}

	Moon = Spinner{
		Frames:   []string{"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "},
		Interval: 100 * time.Millisecond,
	}

	Runner = Spinner{
		Frames:   []string{"🚶 ", "🏃 "},
		Interval: 100 * time.Millisecond,
	}

	Pong = Spinner{
		Frames:   []string{"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌", "▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌", "▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌", "▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌", "▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌", "▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌", "▐⠠       ▌"},
		Interval: 100 * time.Millisecond,
	}

	Shark = Spinner{
		Frames:   []string{"▐|\\____________▌", "▐_|\\___________▌", "▐__|\\__________▌", "▐___|\\_________▌", "▐____|\\________▌", "▐_____|\\_______▌", "▐______|\\______▌", "▐_______|\\_____▌", "▐________|\\____▌", "▐_________|\\___▌", "▐__________|\\__▌", "▐___________|\\_▌", "▐____________|\\▌", "▐____________/|▌", "▐___________/|_▌", "▐__________/|__▌", "▐_________/|___▌", "▐________/|____▌", "▐_______/|_____▌", "▐______/|______▌", "▐_____/|_______▌", "▐____/|________▌", "▐___/|_________▌", "▐__/|__________▌", "▐_/|___________▌", "▐/|____________▌"},
		Interval: 100 * time.Millisecond,
	}

	Dqpb = Spinner{
		Frames:   []string{"d", "q", "p", "b"},
		Interval: 100 * time.Millisecond,
	}

	Weather = Spinner{
		Frames:   []string{"☀️", "☀️", "☀️", "🌤", "⛅️", "🌥", "☁️", "🌧", "🌨", "🌧", "🌨", "🌧", "🌨", "🌨", "🌧", "🌨", "☁️", "🌥", "⛅️", "🌤", "☀️", "☀️"},
		Interval: 100 * time.Millisecond,
	}

	Christmas = Spinner{
		Frames:   []string{"🌲", "🎄"},
		Interval: 400 * time.Millisecond,
	}

	Grenade = Spinner{
		Frames:   []string{"،  ", "′  ", " ´ ", " ‾ ", "  ⸌", "  ⸊", "  |", "  ⁎", "  ⁕", " ෴ ", "  ⁓", "   ", "   ", "   "},
		Interval: 100 * time.Millisecond,
	}

	Point = Spinner{
		Frames:   []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"},
		Interval: 100 * time.Millisecond,
	}

	Layer = Spinner{
		Frames:   []string{"-", "=", "≡", "="},
		Interval: 300 * time.Millisecond,
	}

	BetaWave = Spinner{
		Frames:   []string{"ρββββββ", "βρβββββ", "ββρββββ", "βββρβββ", "ββββρββ", "βββββρβ", "ββββββρ"},
		Interval: 100 * time.Millisecond,
	}

	FingerDance = Spinner{
		Frames:   []string{"🤘 ", "🤟 ", "🖖 ", "✋ ", "🤚 ", "👆 "},
		Interval: 200 * time.Millisecond,
	}

	FistBump = Spinner{
		Frames:   []string{"🤜\u3000\u3000\u3000\u3000🤛 ", "🤜\u3000\u3000\u3000\u3000🤛 ", "🤜\u3000\u3000\u3000\u3000🤛 ", "\u3000🤜\u3000\u3000🤛\u3000 ", "\u3000\u3000🤜🤛\u3000\u3000 ", "\u3000🤜✨🤛\u3000\u3000 ", "🤜\u3000✨\u3000🤛\u3000 "},
		Interval: 100 * time.Millisecond,
	}

	SoccerHeader = Spinner{
		Frames:   []string{" 🧑⚽️       🧑 ", "🧑  ⚽️      🧑 ", "🧑   ⚽️     🧑 ", "🧑    ⚽️    🧑 ", "🧑     ⚽️   🧑 ", "🧑      ⚽️  🧑 ", "🧑       ⚽️🧑  ", "🧑      ⚽️  🧑 ", "🧑     ⚽️   🧑 ", "🧑    ⚽️    🧑 ", "🧑   ⚽️     🧑 ", "🧑  ⚽️      🧑 "},
		Interval: 100 * time.Millisecond,
	}

	Mindblown = Spinner{
		Frames:   []string{"😐 ", "😐 ", "😮 ", "😮 ", "😦 ", "😦 ", "😧 ", "😧 ", "🤯 ", "💥 ", "✨ ", "\u3000 ", "\u3000 ", "\u3000 "},
		Interval: 200 * time.Millisecond,
	}

	Speaker = Spinner{
		Frames:   []string{"🔈 ", "🔉 ", "🔊 ", "🔉 "},
		Interval: 200 * time.Millisecond,
	}

	OrangePulse = Spinner{
		Frames:   []string{"🔸 ", "🔶 ", "🟠 ", "🟠 ", "🔶 "},
		Interval: 100 * time.Millisecond,
	}

	BluePulse = Spinner{
		Frames:   []string{"🔹 ", "🔷 ", "🔵 ", "🔵 ", "🔷 "},
		Interval: 100 * time.Millisecond,
	}

	OrangeBluePulse = Spinner{
		Frames:   []string{"🔸 ", "🔶 ", "🟠 ", "🟠 ", "🔶 ", "🔹 ", "🔷 ", "🔵 ", "🔵 ", "🔷 "},
		Interval: 100 * time.Millisecond,
	}

	TimeTravel = Spinner{
		Frames:   []string{"🕛 ", "🕚 ", "🕙 ", "🕘 ", "🕗 ", "🕖 ", "🕕 ", "🕔 ", "🕓 ", "🕒 ", "🕑 ", "🕐 "},
		Interval: 100 * time.Millisecond,
	}

	Aesthetic = Spinner{
		Frames:   []string{"▰▱▱▱▱▱▱", "▰▰▱▱▱▱▱", "▰▰▰▱▱▱▱", "▰▰▰▰▱▱▱", "▰▰▰▰▰▱▱", "▰▰▰▰▰▰▱", "▰▰▰▰▰▰▰", "▰▱▱▱▱▱▱"},
		Interval: 100 * time.Millisecond,
	}
	AestheticSmall = Spinner{
		Frames:   []string{"▱▱▱", "▰▱▱", "▰▰▱", "▰▰▰", "▰▰▱", "▰▱▱"},
		Interval: 100 * time.Millisecond,
	}

	DwarfFortress = Spinner{
		Frames:   []string{" ██████£££  ", "☺██████£££  ", "☺██████£££  ", "☺▓█████£££  ", "☺▓█████£££  ", "☺▒█████£££  ", "☺▒█████£££  ", "☺░█████£££  ", "☺░█████£££  ", "☺ █████£££  ", " ☺█████£££  ", " ☺█████£££  ", " ☺▓████£££  ", " ☺▓████£££  ", " ☺▒████£££  ", " ☺▒████£££  ", " ☺░████£££  ", " ☺░████£££  ", " ☺ ████£££  ", "  ☺████£££  ", "  ☺████£££  ", "  ☺▓███£££  ", "  ☺▓███£££  ", "  ☺▒███£££  ", "  ☺▒███£££  ", "  ☺░███£££  ", "  ☺░███£££  ", "  ☺ ███£££  ", "   ☺███£££  ", "   ☺███£££  ", "   ☺▓██£££  ", "   ☺▓██£££  ", "   ☺▒██£££  ", "   ☺▒██£££  ", "   ☺░██£££  ", "   ☺░██£££  ", "   ☺ ██£££  ", "    ☺██£££  ", "    ☺██£££  ", "    ☺▓█£££  ", "    ☺▓█£££  ", "    ☺▒█£££  ", "    ☺▒█£££  ", "    ☺░█£££  ", "    ☺░█£££  ", "    ☺ █£££  ", "     ☺█£££  ", "     ☺█£££  ", "     ☺▓£££  ", "     ☺▓£££  ", "     ☺▒£££  ", "     ☺▒£££  ", "     ☺░£££  ", "     ☺░£££  ", "     ☺ £££  ", "      ☺£££  ", "      ☺£££  ", "      ☺▓££  ", "      ☺▓££  ", "      ☺▒££  ", "      ☺▒££  ", "      ☺░££  ", "      ☺░££  ", "      ☺ ££  ", "       ☺££  ", "       ☺££  ", "       ☺▓£  ", "       ☺▓£  ", "       ☺▒£  ", "       ☺▒£  ", "       ☺░£  ", "       ☺░£  ", "       ☺ £  ", "        ☺£  ", "        ☺£  ", "        ☺▓  ", "        ☺▓  ", "        ☺▒  ", "        ☺▒  ", "        ☺░  ", "        ☺░  ", "        ☺   ", "        ☺  &", "        ☺ ☼&", "       ☺ ☼ &", "       ☺☼  &", "      ☺☼  & ", "      ‼   & ", "     ☺   &  ", "    ‼    &  ", "   ☺    &   ", "  ‼     &   ", " ☺     &    ", "‼      &    ", "      &     ", "      &     ", "     &   ░  ", "     &   ▒  ", "    &    ▓  ", "    &    £  ", "   &    ░£  ", "   &    ▒£  ", "  &     ▓£  ", "  &     ££  ", " &     ░££  ", " &     ▒££  ", "&      ▓££  ", "&      £££  ", "      ░£££  ", "      ▒£££  ", "      ▓£££  ", "      █£££  ", "     ░█£££  ", "     ▒█£££  ", "     ▓█£££  ", "     ██£££  ", "    ░██£££  ", "    ▒██£££  ", "    ▓██£££  ", "    ███£££  ", "   ░███£££  ", "   ▒███£££  ", "   ▓███£££  ", "   ████£££  ", "  ░████£££  ", "  ▒████£££  ", "  ▓████£££  ", "  █████£££  ", " ░█████£££  ", " ▒█████£££  ", " ▓█████£££  ", " ██████£££  ", " ██████£££  "},
		Interval: 100 * time.Millisecond,
	}
)
