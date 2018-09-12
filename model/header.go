package model

import "fmt"

// Header is struct of Key and Value, using in list.yaml.
type Header struct {
	Key   string
	Value string
}

// Format will return format of string.
func (h *Header) Format() string {
	return fmt.Sprintf("%-10s: %s", h.Key, h.Value)
}

// String will return string that show what is it.
func (h *Header) String() string {
	return fmt.Sprintf("commit key=%s, value=%s", h.Key, h.Value)
}
