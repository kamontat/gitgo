package exec

import (
	"os"
	"os/exec"

	"github.com/kamontat/gitgo/utils/phase"
)

func Run(builder *OptionBuilder, cmd string, args ...string) (err error) {
	opt := builder.opt // get built option
	commandline := exec.Command(cmd, args...)
	if opt.dryrun {
		phase.Info(commandline.String())
		return
	}

	if opt.out != nil {
		commandline.Stdout = opt.out
	} else {
		commandline.Stdout = os.Stdout
	}

	if opt.err != nil {
		commandline.Stderr = opt.err
	} else {
		commandline.Stderr = os.Stderr
	}

	if opt.in != nil {
		commandline.Stdin = opt.in
	} else {
		commandline.Stdin = os.Stdin
	}

	return commandline.Run()
}
