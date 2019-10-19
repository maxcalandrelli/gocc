// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"fmt"

	"github.com/maxcalandrelli/gocc/example/ctx/ctx1.grammar/ctx1/internal/token"
)

type ActionTable [NumStates]ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (a ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", a.Accept, a.Ignore)
}

var ActTab = ActionTable{
	ActionRow{ // S0, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S1, Ignore("!whitespace")
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S2, Accept("id")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S3, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S4, Accept("id")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S5, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S6, Accept("id")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S7, Accept("id")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S8, Accept("id")
		Accept: 2,
		Ignore: "",
	},
}
