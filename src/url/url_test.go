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
