
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 34 : // ['"','"']
				return 2
			case r == 38 : // ['&','&']
				return 3
			case r == 40 : // ['(','(']
				return 4
			case r == 41 : // [')',')']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 6
			case r == 60 : // ['<','<']
				return 7
			case r == 62 : // ['>','>']
				return 8
			case r == 102 : // ['f','f']
				return 9
			case r == 105 : // ['i','i']
				return 10
			case r == 116 : // ['t','t']
				return 11
			case r == 124 : // ['|','|']
				return 12
			
			
			
			}
			return NoState
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 13
			
			
			default:
				return 2
			
			}
			return NoState
		},
	
		// S3
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S6
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 6
			
			
			
			}
			return NoState
		},
	
		// S7
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S8
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 97 : // ['a','a']
				return 14
			
			
			
			}
			return NoState
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 110 : // ['n','n']
				return 15
			
			
			
			}
			return NoState
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 114 : // ['r','r']
				return 16
			
			
			
			}
			return NoState
		},
	
		// S12
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S13
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 108 : // ['l','l']
				return 17
			
			
			
			}
			return NoState
		},
	
		// S15
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 117 : // ['u','u']
				return 18
			
			
			
			}
			return NoState
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 115 : // ['s','s']
				return 19
			
			
			
			}
			return NoState
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 20
			
			
			
			}
			return NoState
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 21
			
			
			
			}
			return NoState
		},
	
		// S20
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S21
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
}
