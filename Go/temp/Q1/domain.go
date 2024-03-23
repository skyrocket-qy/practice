package main

type TokenCounter interface {
	Ingest(input string)
	Appearance(input string) float32
	Init()
}
