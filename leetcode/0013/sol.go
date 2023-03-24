package sol

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

func romanToInt(s string) int {
	total := 0
	l := len(s)
	next := byte(0)

	for i := 0; i < l; i++ {
		if i == l-1 {
			next = 0
		} else {
			next = s[i+1]
		}

		switch s[i] {
		case 'I':
			switch next {
			case 'V':
				total += 4
				i++
			case 'X':
				total += 9
				i++
			default:
				total += 1
			}
		case 'V':
			total += 5
		case 'X':
			switch next {
			case 'L':
				total += 40
				i++
			case 'C':
				total += 90
				i++
			default:
				total += 10
			}
		case 'L':
			total += 50
		case 'C':
			switch next {
			case 'D':
				total += 400
				i++
			case 'M':
				total += 900
				i++
			default:
				total += 100
			}
		case 'D':
			total += 500
		case 'M':
			total += 1000
		}
	}

	return total
}
