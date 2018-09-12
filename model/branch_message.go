package model

import (
	"path"
	"strings"
)

// BranchName is the object the create branch name.
type BranchName struct {
	Iter  string
	Key   string
	Title string
	Desc  string
	Issue string
}

// GetKey will try to format the key to right way.
// Otherwise, return normal Key
func (b *BranchName) GetKey() string {
	arr := strings.Split(b.Key, ":")
	if len(arr) > 0 {
		s := arr[0]
		return strings.TrimSpace(s)
	}
	return b.Key
}

func (b *BranchName) Name() string {
	return path.Join(b.Iter, b.GetKey(), b.Title, b.Desc, b.Issue)
}
