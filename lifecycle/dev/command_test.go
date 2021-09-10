package dev_test

import (
	"context"
	"testing"
	"time"

	"github.com/wawandco/ox/lifecycle/dev"
	"github.com/wawandco/ox/plugins"
)

type (
	plugin struct {
		ran bool
	}

	developer struct {
		plugin

		ranAt time.Time
	}

	beforeDeveloper struct {
		plugin

		ranAt time.Time
	}
)

func (pl plugin) Name() string {
	return "aaa"
}

func (d *developer) Develop(context.Context, string) error {
	d.ran = true
	d.ranAt = time.Now()
	return nil
}

func (d *beforeDeveloper) BeforeDevelop(context.Context, string) error {
	d.ranAt = time.Now()
	d.ran = true
	time.Sleep(10 * time.Millisecond)

	return nil
}

func TestCommand(t *testing.T) {

	plb := plugin{}

	pls := []plugins.Plugin{
		&plb,
		&developer{},
		&beforeDeveloper{},
	}

	c := dev.Command{}
	c.Receive(pls)

	err := c.Run(context.Background(), "", []string{})
	if err != nil {
		t.Errorf("error should be nil, got %v", err)
	}

	if plb.ran == true {
		t.Errorf("plugin should not have ran in true")
	}

	if pls[1].(*developer).ran == false {
		t.Errorf("developer plugin should run")
	}

	if pls[2].(*beforeDeveloper).ran == false {
		t.Errorf("beforeDeveloper plugin should run")
	}

	if pls[2].(*beforeDeveloper).ranAt.After(pls[1].(*developer).ranAt) {
		t.Errorf("beforeDeveloper plugin should run before developer plugin")
	}
}
