package main

import (
	"fmt"
	"strings"
	"time"
)

type TokenCounter interface {
	Ingest(input string)
	Appearance(input string) float32
	Init()
}

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

func main() {
	fmt.Println("Runnning Implementation 1")
	tc1 := NewTokenCounterImpl1()
	startTime := time.Now()
	Test(tc1)
	fmt.Println("Execute time =", time.Since(startTime))
	fmt.Println()

	fmt.Println("Runnning Implementation 2")
	startTime = time.Now()
	tc2 := NewTokenCounterImpl2()
	Test(tc2)
	fmt.Println("Execute time =", time.Since(startTime))
}

func Test(tc TokenCounter) {
	fmt.Println("Original case")
	tc.Ingest("oursky:uk:dev")
	tc.Ingest("oursky:hk:design")
	tc.Ingest("oursky:hk:pm")
	tc.Ingest("oursky:hk:dev")
	tc.Ingest("skymaker")
	fmt.Println("Appearance('oursky'):", tc.Appearance("oursky"))
	fmt.Println("Appearance('oursky:hk'):", tc.Appearance("oursky:hk"))

	tc.Ingest("skymaker:london:ealing:dev")
	tc.Ingest("skymaker:london:design")
	tc.Ingest("skymaker:london:design")
	tc.Ingest("skymaker:man:pm")
	tc.Ingest("skymaker:man:pm")
	fmt.Println("Appearance('skymaker:man'):", tc.Appearance("skymaker:man"))
	fmt.Println("Appearance('skymaker:lon'):", tc.Appearance("skymaker:lon"))
	fmt.Println("Appearance('london:skymaker'):", tc.Appearance("london:skymaker"))

	// fmt.Println("Customize case")
	// fmt.Println("case 1")
	// tc.Init()
	// tc.Ingest("a:b:c")
	// fmt.Println(tc.Appearance("a"))     // 1
	// fmt.Println(tc.Appearance("b"))     // 1
	// fmt.Println(tc.Appearance("c"))     // 1
	// fmt.Println(tc.Appearance("a:b"))   // 1
	// fmt.Println(tc.Appearance("b:c"))   // 1
	// fmt.Println(tc.Appearance("a:b:c")) // 1
	// fmt.Println(tc.Appearance("a:c"))   // 0

	// fmt.Println("case 2")
	// tc.Init()
	// tc.Ingest("a:b")
	// fmt.Println(tc.Appearance("a"))   // 1
	// fmt.Println(tc.Appearance("b"))   // 1
	// fmt.Println(tc.Appearance("a:b")) // 1
	// fmt.Println(tc.Appearance("c"))   // 0

	// fmt.Println("case 3")
	// tc.Init()
	// tc.Ingest("a")
	// fmt.Println(tc.Appearance("a")) // 1
	// fmt.Println(tc.Appearance("c")) // 0

	// fmt.Println("case 4")
	// tc.Init()
	// tc.Ingest("a:b")
	// tc.Ingest("b:c")
	// fmt.Println(tc.Appearance("a"))     // 0.5
	// fmt.Println(tc.Appearance("b"))     // 1
	// fmt.Println(tc.Appearance("c"))     // 0.5
	// fmt.Println(tc.Appearance("a:b"))   // 0.5
	// fmt.Println(tc.Appearance("a:b:c")) // 0
}
