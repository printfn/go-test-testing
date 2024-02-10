package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

type testDeps struct{}

func (*testDeps) CheckCorpus([]any, []reflect.Type) error {
	panic("unimplemented")
}

func (*testDeps) CoordinateFuzzing(time.Duration, int64, time.Duration, int64, int, []struct {
	Parent     string
	Path       string
	Data       []byte
	Values     []any
	Generation int
	IsSeed     bool
}, []reflect.Type, string, string) error {
	panic("unimplemented")
}

func (*testDeps) ImportPath() string {
	return ""
}

func (*testDeps) MatchString(pat string, str string) (bool, error) {
	panic("unimplemented")
}

func (*testDeps) ReadCorpus(string, []reflect.Type) ([]struct {
	Parent     string
	Path       string
	Data       []byte
	Values     []any
	Generation int
	IsSeed     bool
}, error) {
	panic("unimplemented")
}

func (*testDeps) ResetCoverage() {
	panic("unimplemented")
}

func (*testDeps) RunFuzzWorker(func(struct {
	Parent     string
	Path       string
	Data       []byte
	Values     []any
	Generation int
	IsSeed     bool
}) error) error {
	panic("unimplemented")
}

func (*testDeps) SetPanicOnExit0(bool) {
	panic("unimplemented")
}

func (*testDeps) SnapshotCoverage() {
	panic("unimplemented")
}

func (*testDeps) StartCPUProfile(io.Writer) error {
	panic("unimplemented")
}

func (*testDeps) StartTestLog(io.Writer) {
	panic("unimplemented")
}

func (*testDeps) StopCPUProfile() {
	panic("unimplemented")
}

func (*testDeps) StopTestLog() error {
	panic("unimplemented")
}

func (*testDeps) WriteProfileTo(string, io.Writer, int) error {
	panic("unimplemented")
}

func main() {
	testing.Init()
	m := testing.MainStart(&testDeps{}, []testing.InternalTest{
		{
			Name: "TestHello",
			F: func(t *testing.T) {
				fmt.Println("hello")
				t.Run("subtest", func(t *testing.T) {
					fmt.Println("subtest")
				})
			},
		},
	}, nil, nil, nil)
	os.Exit(m.Run())
}
