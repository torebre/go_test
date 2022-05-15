package url

import "testing"

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"

	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
	}

	want := "https"

	if got := u.Scheme; got != want {
		t.Errorf("Parse (%q).Scheme = %q; want %q", rawurl, got, want)
	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawurl, got, want)
	}

}

func TestURLPort(t *testing.T) {
	tests := []struct {
		name string
		in   string
		port string
	}{
		{name: "with port", in: "foo.com:80", port: "80"},
		{name: "with empty port", in: "foo.com:", port: ""},
		{name: "without port", in: "foo.com", port: ""},
		{name: "ip with port", in: "1.2.3.4:90", port: "90"},
		{name: "ip without port", in: "1.2.3.4", port: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("%s: for host %q; got %q; want %q", tt.name, tt.in, got, want)
			}
		})
	}
}
