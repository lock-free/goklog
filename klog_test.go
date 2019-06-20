package klog

import (
	"fmt"
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"
	"strings"
)

// copied from: https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Errorf("%v != %v", instance1, instance2)
	}
}

func TestKLog_SetListener(t *testing.T) {
	demo := func(text string) {
		fmt.Println(text)
	}

	instance := GetInstance()
	instance.SetListener(demo)
	instance.ToggleInspector()

	actualText := captureOutput(func() {
		instance.LogIn("demo title", "demo text")
	})
	if strings.HasSuffix(actualText, "(demo title) demo text\n") {
		t.Errorf("%s != %s", actualText, "")
	}
}

func TestKLog_ToggleInspector(t *testing.T) {
	instance := GetInstance()
	first := instance.inspectOpened
	instance.ToggleInspector()
	second := instance.inspectOpened

	if first == second {
		t.Errorf("toggleInspector not worked")
	}
}

func TestKLog_Info(t *testing.T) {

}

func TestKLog_LogErr(t *testing.T) {

}

func TestKLog_LogIn(t *testing.T) {

}

func TestKLog_LogVital(t *testing.T) {

}

