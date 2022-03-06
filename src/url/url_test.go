package url

import "testing"

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com"

	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
	}

	want := "https"
	got := u.Scheme

	if got != want {
		t.Logf("Parse(%q).Scheme err = %q, want %q", rawurl, got, want)
		t.Fail()
	}


	if got, want != u.Host, "foo.com"; got != want {
		t.Errorf("Parse (%q).Host = %q; want %q", rawurl, got want)
	}

}
