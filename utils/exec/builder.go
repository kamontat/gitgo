package exec

import "io"

type OptionBuilder struct {
	opt *option
}

func (o *OptionBuilder) WithDryRun() *OptionBuilder {
	o.opt.dryrun = true
	return o
}

func (o *OptionBuilder) SetDryRun(b bool) *OptionBuilder {
	o.opt.dryrun = b
	return o
}

func (o *OptionBuilder) SetReader(reader io.Reader) *OptionBuilder {
	o.opt.in = reader
	return o
}

func (o *OptionBuilder) SetWriter(writer io.Writer) *OptionBuilder {
	o.opt.out = writer
	return o
}

func (o *OptionBuilder) SetEWriter(writer io.Writer) *OptionBuilder {
	o.opt.err = writer
	return o
}

func NewOption() *OptionBuilder {
	return &OptionBuilder{
		opt: &option{},
	}
}
