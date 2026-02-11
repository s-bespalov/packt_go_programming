package main

import (
	"log"
	"testing"
)

func setup() {
	log.Println("setup() running")
}

func teardown() {
	log.Println("teardown() running")
}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	m.Run()
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}

func TestC(t *testing.T) {
	log.Println("TestC running")
}
