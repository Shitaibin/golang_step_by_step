//go:build stb

package build_tag

func NewTag() Tag {
	return &TagDarwin{}
}

type TagDarwin struct {
}

func (t *TagDarwin) GetTag() string {
	return "darwin"
}
