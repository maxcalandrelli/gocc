
package lexer

import(
	"fmt"
	"code.google.com/p/gocc/example/astx/token"
)

type ActionTable [NumStates] ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (this ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", this.Accept, this.Ignore)
}

var ActTab = ActionTable{
 	ActionRow{ // S0
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S1
		Accept: -1,
 		Ignore: "!whitespace",
 	},
 	ActionRow{ // S2
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S3
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S4
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S5
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S6
		Accept: 5,
 		Ignore: "",
 	},
 		
}
