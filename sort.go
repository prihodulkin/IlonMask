package main

import "sort"

type By func(r1, r2 Route) bool

type routeSorter struct {
	routes []Route
	by     By
}

func (s *routeSorter) Len() int {
	return len(s.routes)
}

func (s *routeSorter) Swap(i, j int) {
	s.routes[i], s.routes[j] = s.routes[j], s.routes[i]
}

func (s *routeSorter) Less(i, j int) bool {
	return s.by(s.routes[i], s.routes[j])
}

func (by By) Sort(routes []Route) {
	rs := &routeSorter{
		routes: routes,
		by:     by,
	}
	sort.Sort(rs)
}

func FitnessCmp(r1 Route, r2 Route) bool {
	return fitness(r1) < fitness(r2)
}
