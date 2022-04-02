package main

import (
	"log"
	"testing"
)

func TestInvoke(t *testing.T) {
	server := &HelloServerImpl{}
	invoke, err := Invoke(server, "/SayHello", "World")
	if err != nil {
		t.Errorf("Invoke error: %s", err)
	}
	log.Println(invoke)
}
