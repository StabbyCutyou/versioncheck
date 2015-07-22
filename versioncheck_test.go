package versioncheck

import "testing"

func TestVersionCheck(t *testing.T) {
	Register("Floof", "1.0.0", "= 1.0.0")
	Register("Zoof", "1.0.0", "<~ 1.0")
	Register("Toof", "1.0.0", "< 1.0.1")
	errs := Check()
	if len(errs) > 0 {
		for i := 0; i < len(errs); i++ {
			t.Error(errs[i])
		}
	}
}
