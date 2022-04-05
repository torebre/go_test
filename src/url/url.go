package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
}

func Parse(rawurl string) (*URL, error) {
	i := strings.Index(rawurl, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}

	scheme, rest := rawurl[:i], rawurl[i+3:]
	host := rest
	if i := strings.Index(rest, "/"); i >= 0 {
		host = rest[:i]
	}
	return &URL{scheme, host}, nil
}
