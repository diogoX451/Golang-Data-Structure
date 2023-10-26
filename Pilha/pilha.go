package pilha

import "fmt"

type Pilha[T any] struct {
	size     int
	array    []T
	capacity int
}

func NewPilha[T any](capacity int) *Pilha[T] {
	return &Pilha[T]{
		size:     0,
		array:    make([]T, 0),
		capacity: capacity,
	}
}

func (p *Pilha[T]) isEmpty() bool { return len(p.array) == 0 }

func (p *Pilha[T]) isFull() bool {
	return len(p.array) == p.capacity
}

func (p *Pilha[T]) Push(value T) {
	if p.isFull() {
		fmt.Println("Pilha cheia")
		return
	}

	p.array = append(p.array, value)
	p.size++
}

func (p *Pilha[T]) Pop() interface{} {
	if p.size == 0 {
		return nil
	}

	value := p.array[p.size-1]
	p.array = p.array[:p.size-1]
	p.size--
	return value
}

func (p *Pilha[T]) List() []T {
	return p.array
}
