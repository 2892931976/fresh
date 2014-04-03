package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestNewCommand(t *testing.T) {
	c := newCommand("build", "./build all -o foo")
	assert.Equal(t, "build", c.Name)
	assert.Equal(t, "./build all -o foo", c.CmdString)
}

func TestCommand_Build(t *testing.T) {
	c := newCommand("build", "./build all -o foo")
	assert.Nil(t, c.Cmd)

	c.build()
	assert.NotNil(t, c.Cmd)
	assert.Equal(t, "./build", c.Cmd.Path)
	assert.Equal(t, []string{"./build", "all", "-o", "foo"}, c.Cmd.Args)
}
