// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"fmt"

	"github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/token"
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
	ActionRow{ // S2, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S3, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S4, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S5, Accept("Λ<(>")
		Accept: 12,
		Ignore: "",
	},
	ActionRow{ // S6, Accept("Λ<)>")
		Accept: 13,
		Ignore: "",
	},
	ActionRow{ // S7, Accept("Λ<->")
		Accept: 10,
		Ignore: "",
	},
	ActionRow{ // S8, Accept("Λ<.>")
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S9, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S10, Accept("Λ<:>")
		Accept: 3,
		Ignore: "",
	},
	ActionRow{ // S11, Accept("Λ<;>")
		Accept: 4,
		Ignore: "",
	},
	ActionRow{ // S12, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S13, Accept("Λ<@>")
		Accept: 22,
		Ignore: "",
	},
	ActionRow{ // S14, Accept("prodId")
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S15, Accept("Λ<[>")
		Accept: 14,
		Ignore: "",
	},
	ActionRow{ // S16, Accept("Λ<]>")
		Accept: 15,
		Ignore: "",
	},
	ActionRow{ // S17, Accept("regDefId")
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S18, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S19, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S20, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S21, Accept("Λ<{>")
		Accept: 16,
		Ignore: "",
	},
	ActionRow{ // S22, Accept("Λ<|>")
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S23, Accept("Λ<}>")
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S24, Accept("Λ<~>")
		Accept: 11,
		Ignore: "",
	},
	ActionRow{ // S25, Accept("Λ<ε>")
		Accept: 26,
		Ignore: "",
	},
	ActionRow{ // S26, Accept("Λ<λ>")
		Accept: 24,
		Ignore: "",
	},
	ActionRow{ // S27, Accept("ignoredTokId")
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S28, Accept("string_lit")
		Accept: 20,
		Ignore: "",
	},
	ActionRow{ // S29, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S30, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S31, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S32, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S33, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S34, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S35, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S36, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S37, Accept("prodId")
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S38, Accept("prodId")
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S39, Accept("prodId")
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S40, Accept("prodId")
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S41, Accept("regDefId")
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S42, Accept("regDefId")
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S43, Accept("regDefId")
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S44, Accept("regDefId")
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S45, Accept("string_lit")
		Accept: 20,
		Ignore: "",
	},
	ActionRow{ // S46, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S47, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S48, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S49, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S50, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S51, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S52, Accept("ignoredTokId")
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S53, Accept("ignoredTokId")
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S54, Accept("ignoredTokId")
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S55, Accept("ignoredTokId")
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S56, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S57, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S58, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S59, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S60, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S61, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S62, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S63, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S64, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S65, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S66, Accept("char_lit")
		Accept: 9,
		Ignore: "",
	},
	ActionRow{ // S67, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S68, Ignore("!comment")
		Accept: -1,
		Ignore: "!comment",
	},
	ActionRow{ // S69, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S70, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S71, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S72, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S73, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S74, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S75, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S76, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S77, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S78, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S79, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S80, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S81, Ignore("!comment")
		Accept: -1,
		Ignore: "!comment",
	},
	ActionRow{ // S82, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S83, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S84, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S85, Accept("tokId")
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S86, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S87, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S88, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S89, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S90, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S91, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S92, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S93, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S94, Accept("g_sdt_lit")
		Accept: 18,
		Ignore: "",
	},
	ActionRow{ // S95, Accept("g_ctxdep_lit")
		Accept: 21,
		Ignore: "",
	},
	ActionRow{ // S96, Accept("Λ<empty>")
		Accept: 25,
		Ignore: "",
	},
	ActionRow{ // S97, Accept("Λ<error>")
		Accept: 23,
		Ignore: "",
	},
	ActionRow{ // S98, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S99, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S100, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S101, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S102, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S103, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S104, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S105, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S106, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S107, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S108, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S109, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S110, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S111, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S112, 
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S113, 
		Accept: 0,
		Ignore: "",
	},
}
