
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
			case 65 <= r && r <= 90 : // ['A','Z']
				return 2
			case r == 95 : // ['_','_']
				return 3
			case 97 <= r && r <= 100 : // ['a','d']
				return 2
			case r == 101 : // ['e','e']
				return 4
			case 102 <= r && r <= 104 : // ['f','h']
				return 2
			case r == 105 : // ['i','i']
				return 5
			case 106 <= r && r <= 115 : // ['j','s']
				return 2
			case r == 116 : // ['t','t']
				return 6
			case 117 <= r && r <= 122 : // ['u','z']
				return 2
			
			
			
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
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S3
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S4
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 107 : // ['a','k']
				return 8
			case r == 108 : // ['l','l']
				return 10
			case 109 <= r && r <= 122 : // ['m','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S5
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 101 : // ['a','e']
				return 8
			case r == 102 : // ['f','f']
				return 11
			case 103 <= r && r <= 122 : // ['g','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S6
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 103 : // ['a','g']
				return 8
			case r == 104 : // ['h','h']
				return 12
			case 105 <= r && r <= 122 : // ['i','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S7
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S8
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S9
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S10
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 114 : // ['a','r']
				return 8
			case r == 115 : // ['s','s']
				return 13
			case 116 <= r && r <= 122 : // ['t','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S11
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S12
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 100 : // ['a','d']
				return 8
			case r == 101 : // ['e','e']
				return 14
			case 102 <= r && r <= 122 : // ['f','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S13
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 100 : // ['a','d']
				return 8
			case r == 101 : // ['e','e']
				return 15
			case 102 <= r && r <= 122 : // ['f','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S14
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 109 : // ['a','m']
				return 8
			case r == 110 : // ['n','n']
				return 16
			case 111 <= r && r <= 122 : // ['o','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S15
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S16
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 9
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
		},
	
}
