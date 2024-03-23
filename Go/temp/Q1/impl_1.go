package main

import (
	"strings"
)

type TokenCounterImpl1 struct {
	tokens []string
}

func NewTokenCounterImpl1() *TokenCounterImpl1 {
	return &TokenCounterImpl1{
		tokens: []string{},
	}
}

func (tc *TokenCounterImpl1) Ingest(input string) {
	tc.tokens = append(tc.tokens, input)
}

func (tc *TokenCounterImpl1) Appearance(input string) float32 {
	var appear float32
	for _, token := range tc.tokens {
		findIndex := strings.Index(token, input)
		if findIndex != -1 {
			lastIndex := len(input) + findIndex - 1
			if lastIndex != len(token)-1 && token[lastIndex+1] != ':' {
				continue
			}
			appear++
		}
	}

	return appear / float32(len(tc.tokens))
}

func (tc *TokenCounterImpl1) Init() {
	tc.tokens = []string{}
}
