package test

import "github.com/wawandco/ox/plugins"

// Receive takes BeforeTesters, AfterTesters and Testers
// from the passed list of pugins and save those in the
// instance of Command so these get used later on.
func (b *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {

		if ptool, ok := plugin.(BeforeTester); ok {
			b.beforeTesters = append(b.beforeTesters, ptool)
		}

		if ptool, ok := plugin.(Tester); ok {
			b.testers = append(b.testers, ptool)
		}

		if ptool, ok := plugin.(AfterTester); ok {
			b.afterTesters = append(b.afterTesters, ptool)
		}
	}
}
