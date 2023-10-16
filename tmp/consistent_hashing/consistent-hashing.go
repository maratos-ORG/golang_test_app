package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Node структура, представляющая физический узел в кольце.
type Node struct {
	Name          string
	VirtualTokens []int
}

// Ring структура, представляющая кольцо.
type Ring struct {
	Tokens         map[int]string // Отображение токена на имя узла
	ExistingTokens map[int]bool   // Для проверки уникальности токенов
}

// NewRing создает новое кольцо.
func NewRing() *Ring {
	return &Ring{
		Tokens:         make(map[int]string),
		ExistingTokens: make(map[int]bool),
	}
}

// AddNode добавляет новый узел в кольцо.
func (r *Ring) AddNode(node Node) {
	for _, t := range node.VirtualTokens {
		r.Tokens[t] = node.Name
		r.ExistingTokens[t] = true
	}
}

// RemoveNode удаляет узел из кольца.
func (r *Ring) RemoveNode(nodeName string) {
	for t, n := range r.Tokens {
		if n == nodeName {
			delete(r.Tokens, t)
			delete(r.ExistingTokens, t)
		}
	}
}

// GenerateUniqueVirtualTokens генерирует уникальные токены.
func (r *Ring) GenerateUniqueVirtualTokens(n int, max int) []int {
	tokens := make([]int, 0, n)
	for len(tokens) < n {
		candidate := rand.Intn(max)
		if _, exists := r.ExistingTokens[candidate]; !exists {
			tokens = append(tokens, candidate)
			r.ExistingTokens[candidate] = true
		}
	}
	return tokens
}

// FindNode находит узел для данного ключа.
func (r *Ring) FindNode(key int) string {
	var keys []int
	for k := range r.Tokens {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if key <= k {
			return r.Tokens[k]
		}
	}
	return r.Tokens[keys[0]]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ring := NewRing()

	node0 := Node{Name: "Node0", VirtualTokens: ring.GenerateUniqueVirtualTokens(8, 100)}
	node1 := Node{Name: "Node1", VirtualTokens: ring.GenerateUniqueVirtualTokens(8, 100)}

	ring.AddNode(node0)
	ring.AddNode(node1)

	keys := []int{7, 22, 47, 77}

	for _, key := range keys {
		node := ring.FindNode(key)
		fmt.Printf("Key: %d, Node: %s\n", key, node)
	}

	ring.RemoveNode("Node0")

	fmt.Println("After removing Node0:")
	for _, key := range keys {
		node := ring.FindNode(key)
		fmt.Printf("Key: %d, Node: %s\n", key, node)
	}
}
