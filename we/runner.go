package we

import (
	"context"
	"os"
	"os/exec"
	"syscall"

	"github.com/go-errors/errors"
)

type Runner struct {
	Commands []string `json:"commands"`
	WorkDir  string   `json:"workdir"`
	EnvVars  []string `json:"envvars"`
}

func (p *Runner) Run(ctx context.Context) error {
	localCmd := exec.CommandContext(ctx, p.Commands[0], p.Commands[1:]...)
	localCmd.Env = p.EnvVars

	if p.WorkDir != "" {
		localCmd.Dir = p.WorkDir
	}
	localCmd.Stdin = os.Stdin
	localCmd.Stdout = os.Stdout
	localCmd.Stderr = os.Stderr

	if err := localCmd.Start(); err != nil {
		return err
	}

	err := localCmd.Wait()
	var exitStatus int
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitStatus = 1

			// There is no process-independent way to get the REAL
			// exit status so we just try to go deeper.
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitStatus = status.ExitStatus()
			}
		}

		if exitStatus != 0 {
			return errors.Errorf("command exits with status %d", exitStatus)
		}
	}

	return nil
}
