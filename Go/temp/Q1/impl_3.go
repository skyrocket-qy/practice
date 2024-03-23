package main

import (
	"strings"
)

type Node2 struct {
	Token    string
	Start    int
	Count    int
	End      int
	Children map[string]*Node2
}

type TokenCounterImpl2 struct {
	TokenNodeMap map[string]*Node2
	Total        int
}

func NewTokenCounterImpl2() *TokenCounterImpl2 {
	return &TokenCounterImpl2{
		TokenNodeMap: map[string]*Node2{},
		Total:        0,
	}
}

func (tc *TokenCounterImpl2) Ingest(input string) {
	tc.Total++
	tokens := strings.Split(input, ":")
	var curNode *Node2
	for i, token := range tokens {
		var node *Node2

		// get or create node
		if _, ok := tc.TokenNodeMap[token]; !ok {
			node = &Node2{
				Token:    token,
				Children: map[string]*Node2{},
			}
			tc.TokenNodeMap[token] = node
		} else {
			node = tc.TokenNodeMap[token]
		}

		// add children
		if curNode != nil {
			curNode.Children[node.Token] = node
		}
		curNode = node
		curNode.Count++

		if i == 0 {
			curNode.Start++
		}
		if i == len(tokens)-1 {
			curNode.End++
		}
	}
}

func (tc *TokenCounterImpl2) Appearance(input string) float32 {
	tokens := strings.Split(input, ":")
	var process int
	var curNode *Node2
	for i, token := range tokens {
		if curNode == nil {
			if _, ok := tc.TokenNodeMap[token]; !ok {
				return 0
			}
			curNode = tc.TokenNodeMap[token]
		} else {
			isFind := false
			for _, child := range curNode.Children {
				if child.Token == token {
					curNode = child
					isFind = true
					break
				}
			}
			if !isFind {
				return 0
			}
		}

		if i == 0 {
			process = curNode.Count
		} else if process > curNode.Count-curNode.Start {
			// update the count to rest count
			process = curNode.Count - curNode.Start
		}
		if i == len(tokens)-1 {
			return float32(process) / float32(tc.Total)
		}

		// subtract the input tail
		process -= curNode.End
		if process <= 0 {
			return 0
		}
	}

	return 0
}

func (tc *TokenCounterImpl2) Init() {
	tc.TokenNodeMap = map[string]*Node2{}
	tc.Total = 0
}
