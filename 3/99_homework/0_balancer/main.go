package main

import (
	"sync"
)


type RoundRobinBalancer struct {
	mu    sync.RWMutex
	servers []int
}

func (rrb *RoundRobinBalancer) Init(count int) {
	rrb.servers = make([]int, count)
}

func (rrb *RoundRobinBalancer) GiveStat() []int {
	rrb.mu.RLock()
	defer rrb.mu.RUnlock()

	return rrb.servers
}

func (rrb *RoundRobinBalancer) GiveNode() []int {
	rrb.mu.Lock()
	defer rrb.mu.Unlock()

	for i, server := range rrb.servers {

		if server == 0 {
			rrb.servers[i]++
			return rrb.servers
		}

		if i != 0 && server < rrb.servers[i-1] {
			rrb.servers[i]++
			return rrb.servers
		}

		if i == len(rrb.servers) - 1 {
			rrb.servers[0]++
			return rrb.servers
		}

	}
	return rrb.servers
}