package trie

import (
	"testing"
)

var trie = New()

func TestInsert(t *testing.T) {
	trie.insert("testing")
	trie.insert("test")
	trie.insert("tested")
	trie.insert("dominique")
	trie.insert("partial")
}

func TestFindTesting(t *testing.T) {
	node, ok := trie.find("testing")

	if node.value != 'g' {
		t.Errorf("Expected value to be g, but it was %s instead.", string(node.value))
	}

	if node.terminal != true {
		t.Errorf("Expected value to be true, but it was %t instead.", node.terminal)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindTest(t *testing.T) {
	node, ok := trie.find("test")

	if node.value != 't' {
		t.Errorf("Expected value to be t, but it was %s instead.", string(node.value))
	}

	if node.terminal != true {
		t.Errorf("Expected value to be true, but it was %t instead.", node.terminal)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindTested(t *testing.T) {
	node, ok := trie.find("tested")

	if node.value != 'd' {
		t.Errorf("Expected value to be d, but it was %s instead.", string(node.value))
	}

	if node.terminal != true {
		t.Errorf("Expected value to be true, but it was %t instead.", node.terminal)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindDominique(t *testing.T) {
	node, ok := trie.find("dominique")

	if node.value != 'e' {
		t.Errorf("Expected value to be e, but it was %s instead.", string(node.value))
	}

	if node.terminal != true {
		t.Errorf("Expected value to be true, but it was %t instead.", node.terminal)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindPartial(t *testing.T) {
	node, ok := trie.find("partial")

	if node.value != 'l' {
		t.Errorf("Expected value to be l, but it was %s instead.", string(node.value))
	}

	if node.terminal != true {
		t.Errorf("Expected value to be true, but it was %t instead.", node.terminal)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindPart(t *testing.T) {
	node, ok := trie.find("part")

	if node.value != 't' {
		t.Errorf("Expected value to be t, but it was %s instead.", string(node.value))
	}

	if node.terminal != false {
		t.Errorf("Expected value to be false, but it was %t instead.", node.terminal)
	}

	if !ok {
		t.Errorf("Expected ok to be true, but it was %t instead.", ok)
	}
}

func TestFindNonExisting(t *testing.T) {
	node, ok := trie.find("non_existing")

	if node != nil {
		t.Errorf("Expected node to be nil, but it was %T instead.", node)
	}

	if ok {
		t.Errorf("Expected ok to be false, but it was %t instead.", ok)
	}
}
