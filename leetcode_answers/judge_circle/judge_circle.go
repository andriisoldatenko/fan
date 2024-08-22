package main

func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, c := range moves {
		if c == 'U' {
			y++
		} else if c == 'D' {
			y--
		} else if c == 'R' {
			x++
		} else {
			x--
		}
	}
	return (x == y) && (x == 0)
}

func main() {
	judgeCircle("RU")
}
