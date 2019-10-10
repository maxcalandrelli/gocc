// Code generated by gocc; DO NOT EDIT.

package iface

import (
	"github.com/maxcalandrelli/gocc/example/errorrecovery/internal/token"
	"io"
)

type (
	Scanner interface {
		Scan() (tok *token.Token)
		GetStream() TokenStream
	}

	CheckPoint interface {
		DistanceFrom(CheckPoint) int
	}

	CheckPointable interface {
		GetCheckPoint() CheckPoint
		GotoCheckPoint(CheckPoint)
	}

	TokenStream interface {
		io.Reader
		io.RuneScanner
		io.Seeker
	}
)
