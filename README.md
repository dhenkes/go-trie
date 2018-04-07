# go-trie

This is a simple go trie implementation.

## How does it work?

```go
// Create a new trie
trie := New()

// Insert new values
trie.insert("testing")
trie.insert("test")
trie.insert("tested")

// Find inserted values
node, ok := trie.find("abc") // Returns nil, false
node, ok = trie.find("testing") // Returns *Node, true (terminal true)
node, ok = trie.find("tested") // Returns *Node, true (terminal true)
node, ok = trie.find("teste") // Returns *Node, true (terminal false)
```
