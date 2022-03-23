package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

const testFile = `[
	{
		"date": "2022-01-01",
		"description": "foo"
	}
]`

func Test(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())
	if _, err := f.Write([]byte(testFile)); err != nil {
		log.Fatal(err)
	}

	want := []Event{{Date: "2022-01-01", Description: "foo"}}
	got := ReadEvents(f.Name())
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected output: got %v, want %v", got, want)
	}
}
