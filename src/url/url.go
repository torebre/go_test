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

	scheme, host := rawurl[:i], rawurl[i+3:]
	return &URL{scheme, host}, nil
}
