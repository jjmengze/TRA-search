package cmd

import (
	"io"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

type PluginHandler interface {
	Execute(executablePath string, cmdArgs, environment []string) error
}

type defaultPluginHandler struct{}

func (h *defaultPluginHandler) Execute(executablePath string, cmdArgs, environment []string) error {
	return syscall.Exec(executablePath, cmdArgs, environment)
}

func NewDefaultSearchTRACommand() *cobra.Command {
	return NewDefaultSearchTRACommandWithArgs(&defaultPluginHandler{}, os.Args, os.Stdin, os.Stdout, os.Stderr)
}
func NewDefaultSearchTRACommandWithArgs(pluginHandler PluginHandler, args []string, in io.Reader, out, errout io.Writer) *cobra.Command {
	return SearchTRACommand(in, out, errout)
}

func SearchTRACommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "tra",
		Short: "Search tra data",
		Long:  "This is very sample cmd of tra search engine",
		Run:   runHelp,
	}
	return cmds
}
func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
