package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* name */
			shift(3), /* hello */
			shift(4), /* hiya */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* name */
			nil,          /* hello */
			nil,          /* hiya */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(5), /* name */
			nil,      /* hello */
			nil,      /* hiya */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(2), /* name, reduce: Saying */
			nil,       /* hello */
			nil,       /* hiya */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(3), /* name, reduce: Saying */
			nil,       /* hello */
			nil,       /* hiya */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Hello */
			nil,       /* name */
			nil,       /* hello */
			nil,       /* hiya */

		},
	},
}
