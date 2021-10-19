package main

import "sort"

type By func(r1, r2 ChromosomeData) bool

type chromosomeSorter struct {
	populationData PopulationData
	by             By
}

func (s *chromosomeSorter) Len() int {
	return len(s.populationData)
}

func (s *chromosomeSorter) Swap(i, j int) {
	s.populationData[i], s.populationData[j] = s.populationData[j], s.populationData[i]
}

func (s *chromosomeSorter) Less(i, j int) bool {
	return s.by(s.populationData[i], s.populationData[j])
}

func (by By) Sort(routes PopulationData) {
	rs := &chromosomeSorter{
		populationData: routes,
		by:             by,
	}
	sort.Sort(rs)
}

func FitnessCmp(d1 ChromosomeData, d2 ChromosomeData) bool {
	return Fitness(d1.FitnessData)<Fitness(d2.FitnessData)
}
