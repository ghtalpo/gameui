package ui

import "testing"

// BenchmarkDrawText-2     	200000000	         7.15 ns/op   (mbp-2010)
func BenchmarkDrawText(b *testing.B) {

	txt := NewText("HEJ", 6)
	for n := 0; n < b.N; n++ {
		txt.Draw()
	}
}

func TestTextOnly(t *testing.T) {
	txt := NewText("HEJ", 6)

	// XXX fixme height is too high
	im := txt.Draw()
	testCompareRender(t, []string{
		"# # ###   # ",
		"# # ##    # ",
		"### #   # # ",
		"# # ###  #  ",
		"            ",
	}, renderAsText(im))
}
