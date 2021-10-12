package main

type LandingPositionHeap []ShuttleState



func (h LandingPositionHeap) Len() int { return len(h) }

func (h LandingPositionHeap) Less(i, j int) bool {
	return landingFitness(h[i].x, xFlatMin, xFlatMax) > landingFitness(h[j].x, xFlatMin, xFlatMax)
}
func (h LandingPositionHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *LandingPositionHeap) Push(x interface{}) {
	*h = append(*h, x.(ShuttleState))
}

func (h *LandingPositionHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type HSpeedHeap []ShuttleState

func (h HSpeedHeap) Len() int { return len(h) }

func (h HSpeedHeap) Less(i, j int) bool {
	return hSpeedFitness(h[i].hSpeed) > hSpeedFitness(h[i].vSpeed)
}
func (h HSpeedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *HSpeedHeap) Push(x interface{}) {
	*h = append(*h, x.(ShuttleState))
}

func (h *HSpeedHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type VSpeedHeap []ShuttleState

func (h VSpeedHeap) Len() int { return len(h) }

func (h VSpeedHeap) Less(i, j int) bool {
	return vSpeedFitness(h[i].vSpeed) > vSpeedFitness(h[j].vSpeed)
}
func (h VSpeedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *VSpeedHeap) Push(x interface{}) {
	*h = append(*h, x.(ShuttleState))
}

func (h *VSpeedHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
