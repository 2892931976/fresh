package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/pilu/config"
)

type runner struct {
	Sections []*section
	DoneChan chan bool
	SigChan  chan os.Signal
}

func newRunner() *runner {
	r := &runner{
		DoneChan: make(chan bool),
		SigChan:  make(chan os.Signal),
	}

	signal.Notify(r.SigChan, os.Interrupt)
	signal.Notify(r.SigChan, os.Kill)

	return r
}

func newRunnerWithFreshfile(freshfilePath string) (*runner, error) {
	r := newRunner()

	sections, err := config.ParseFile(freshfilePath, "main: *")
	if err != nil {
		return r, err
	}

	for s, opts := range sections {
		section := r.NewSection(s)
		for name, cmd := range opts {
			section.NewCommand(name, cmd)
		}
	}

	return r, nil
}

func (r *runner) NewSection(description string) *section {
	s := newSection(description)
	r.Sections = append(r.Sections, s)
	return s
}

func (r *runner) Run() {
	logger.log("Running...")
	logger.log("%d goroutines", runtime.NumGoroutine())
	go r.ListenForSignals()

	for _, s := range r.Sections {
		go func(s *section) {
			s.Run()
		}(s)
	}
}

func (r *runner) Stop() {
	logger.log("Stopping all sections")
	for _, s := range r.Sections {
		s.Stop()
	}
}

func (r *runner) ListenForSignals() {
	logger.log("Listening for signals")
	<-r.SigChan
	fmt.Printf("Interrupt a second time to quit\n")
	logger.log("Waiting for a second signal")
	select {
	case <-r.SigChan:
		logger.log("Second signal received")
		r.DoneChan <- true
	case <-time.After(1 * time.Second):
		logger.log("Timeout")
		logger.log("Stopping...")
		r.Stop()
		logger.log("Calling Run...")
		r.Run()
	}
}
