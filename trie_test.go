package trie

import (
	"testing"
)

func TestFindExisting(t *testing.T) {
	trie := New()
	trie.insert("test")
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

func TestFindPartial(t *testing.T) {
	trie := New()
	trie.insert("partial")
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
	trie := New()
	trie.insert("dominique")
	node, ok := trie.find("non_existing")

	if node != nil {
		t.Errorf("Expected node to be nil, but it was %T instead.", node)
	}

	if ok {
		t.Errorf("Expected ok to be false, but it was %t instead.", ok)
	}
}

func BenchmarkInsert(b *testing.B) {
	trie := New()

	for n := 0; n < b.N; n++ {
		trie.insert("word")
	}
}

func BenchmarkInsertSimilar(b *testing.B) {
	trie := New()

	for n := 0; n < b.N; n++ {
		trie.insert("testing")
		trie.insert("tested")
	}
}

func BenchmarkInsertDifferent(b *testing.B) {
	trie := New()

	for n := 0; n < b.N; n++ {
		trie.insert("patricia")
		trie.insert("dominique")
	}
}

func BenchmarkFindExisting(b *testing.B) {
	trie := New()
	trie.insert("testing")
	trie.insert("test")
	trie.insert("tested")
	trie.insert("word")
	trie.insert("partial")
	trie.insert("dominique")
	trie.insert("patricia")

	for n := 0; n < b.N; n++ {
		_, _ = trie.find("dominique")
	}
}

func BenchmarkFindPartial(b *testing.B) {
	trie := New()
	trie.insert("testing")
	trie.insert("test")
	trie.insert("tested")
	trie.insert("word")
	trie.insert("partial")
	trie.insert("dominique")
	trie.insert("patricia")

	for n := 0; n < b.N; n++ {
		_, _ = trie.find("part")
	}
}

func BenchmarkFindNonExisting(b *testing.B) {
	trie := New()
	trie.insert("testing")
	trie.insert("test")
	trie.insert("tested")
	trie.insert("word")
	trie.insert("partial")
	trie.insert("dominique")
	trie.insert("patricia")

	for n := 0; n < b.N; n++ {
		_, _ = trie.find("non_existing")
	}
}
