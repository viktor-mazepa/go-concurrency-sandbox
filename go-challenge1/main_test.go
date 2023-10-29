package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	testMessage := "test_message"

	wg.Add(1)

	go updateMessage(testMessage, &wg)

	wg.Wait()

	if strings.Compare(msg, testMessage) != 0 {
		t.Error("Test_updateMessage. Expected to find:", testMessage, "but found:", msg)
	}
}

func Test_printMessage(t *testing.T) {
	testMessage := "test_message"

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	wg.Add(1)

	go updateMessage(testMessage, &wg)

	wg.Wait()

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, testMessage) {
		t.Error("Expected to find ", testMessage, ", but it is no there")
	}
}

func Test_main(t *testing.T) {

	messages := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	for _, str := range messages {
		if !strings.Contains(output, str) {
			t.Error("Expected to find ", str, ", but it is no there")
		}
	}
}
