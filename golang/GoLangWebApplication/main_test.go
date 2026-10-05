package main

import "testing"

func TestMain(t *testing.T) {
	go main()
	t.Log("TestMain executed")
}