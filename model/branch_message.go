package model

import (
	"path"
	"strings"
)

// BranchName is the object the create branch name.
type BranchName struct {
	Iter  string
	Type  string
	Title string
	Desc  string
	Issue string
}

// GetType will try to format the key to right way.
// Otherwise, return normal Key
func (b *BranchName) GetType() string {
	arr := strings.Split(b.Type, ":")
	if len(arr) > 0 {
		s := arr[0]
		return strings.TrimSpace(s)
	}
	return b.Type
}

func (b *BranchName) Name() string {
	return path.Join(b.Iter, b.GetType(), b.Title, b.Desc, b.Issue)
}
