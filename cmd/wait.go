package cmd

import (
	"fmt"
	"strconv"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// Wait represents `wait` subcommand.
type Wait struct {
	cli  CLI
	repo repository.Repository
}

// NewWait returns a new Wait.
func NewWait(cli CLI, version string, repo repository.Repository) Command {
	return &Wait{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Wait) Run(args []string) int {
	if len(args) < 3 {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	td, err := c.repo.Get(id)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	td.State = todo.Waiting
	err = c.repo.Update(td)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	return 0
}