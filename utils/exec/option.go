package exec

import "io"

type option struct {
	dryrun bool

	in  io.Reader
	out io.Writer
	err io.Writer
}
