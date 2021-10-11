package hw02_testing_tool

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

type Tester interface {
	RunTests() error
}

type tester struct {
	task ITask
	path string
}

func NewTester(task ITask, path string) Tester {
	return &tester{
		task: task,
		path: path,
	}
}

func (t tester) RunTests() error {
	var i int
	fmt.Printf("=== TESTS PATH\t%s\n", t.path)

	for {
		fmt.Printf("=== RUN:\t#%d\n", i)
		in, err := ioutil.ReadFile(path.Join(t.path, fmt.Sprintf("test.%d.in", i)))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) || errors.Is(err, os.ErrPermission) {
				break
			}
			panic(err)
		}

		out, err := ioutil.ReadFile(path.Join(t.path, fmt.Sprintf("test.%d.out", i)))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) || errors.Is(err, os.ErrPermission) {
				break
			}
			panic(err)
		}
		expected := strings.TrimFunc(strings.ReplaceAll(string(out), "\r\n", "\n"), func(r rune) bool {
			return r == ' ' || r == '\t' || r == '\n' || r == '\r'
		})

		start := time.Now()
		result := t.task.Run([]string{
			strings.TrimFunc(string(in), func(r rune) bool {
				return r == ' ' || r == '\t' || r == '\n' || r == '\r'
			}),
		})
		finish := time.Since(start).Milliseconds()

		if strings.EqualFold(result, expected) {
			fmt.Printf("--- PASS:\t#%d\t%dms\n", i, finish)
			fmt.Printf("\tExpect:\t%s\n", expected)
			fmt.Printf("\tActual:\t%s\n", result)
		} else {
			fmt.Printf("--- FAILED:\t#%d\n", i)
			fmt.Printf("\tExpect:\t%s\n", expected)
			fmt.Printf("\tActual:\t%s\n", result)
			return fmt.Errorf("name: %s; test: %d", t.path, i)
		}

		i++
	}
	return nil
}
