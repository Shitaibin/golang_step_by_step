package build_tag

import "testing"

func TestTag(t *testing.T) {
	tag := NewTag()
	t.Log(tag.GetTag())
}
