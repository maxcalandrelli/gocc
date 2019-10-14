package ast

import (
	"github.com/maxcalandrelli/gocc/internal/frontend/stock/token"
)

var (
	StringGetter func(interface{}) string
)

func getString(v interface{}) string {
	if StringGetter == nil {
		if str, ok := v.(string); ok {
			return str
		}
		if tok, _ := v.(*token.Token); tok != nil {
			return string(tok.Lit)
		}
	}
	return StringGetter(v)
}

func unquoteString(str string) (string, bool, rune) {
	if len(str) > 1 {
		r := str[0]
		if r == '"' || r == '`' || r == '\'' {
			str = str[1 : len(str)-1]
		}
		return str, true, rune(r)
	}
	return str, false, 0
}

/*
func unescapeString(str string) string {
	res := ""
	for s := 0; s < len(str); {
		if str[s] == '\\' {
			res = res + string(escapeCharVal(lit))
			s++
		}
		r, size := utf8.DecodeRune(lit[1:])
	}
	if size != len(lit)-2 {
		panic(fmt.Sprintf("Error decoding rune. Lit: %s, rune: %d, size %d\n", lit, r, size))
	}
	return r
}

func escapeCharVal(lit []byte) (rune, int) {
	var i, base, max uint32
	offset := 0
	switch lit[offset] {
	case 'a':
		return '\a'
	case 'b':
		return '\b'
	case 'f':
		return '\f'
	case 'n':
		return '\n'
	case 'r':
		return '\r'
	case 't':
		return '\t'
	case 'v':
		return '\v'
	case '\\':
		return '\\'
	case '\'':
		return '\''
	case '0', '1', '2', '3', '4', '5', '6', '7':
		i, base, max = 3, 8, 255
	case 'x':
		i, base, max = 2, 16, 255
		offset++
	case 'u':
		i, base, max = 4, 16, unicode.MaxRune
		offset++
	case 'U':
		i, base, max = 8, 16, unicode.MaxRune
		offset++
	default:
		panic(fmt.Sprintf("Error decoding character literal: %s\n", lit))
	}

	var x uint32
	for ; i > 0 && offset < len(lit)-1; i-- {
		ch, size := utf8.DecodeRune(lit[offset:])
		offset += size
		d := uint32(digitVal(ch))
		if d >= base {
			panic(fmt.Sprintf("charVal(%s): illegal character (%c) in escape sequence. size=%d, offset=%d", lit, ch, size, offset))
		}
		x = x*base + d
	}
	if x > max || 0xD800 <= x && x < 0xE000 {
		panic(fmt.Sprintf("Error decoding escape char value. Lit:%s, offset:%d, escape sequence is invalid Unicode code point\n", lit, offset))
	}

	return rune(x)
}

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}
*/
