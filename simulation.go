package main

type ShuttleState struct {
	hSpeed float64
	vSpeed float64
	fuel   int
}



func (state *ShuttleState) Init(data *ShuttleData) {
	state.hSpeed = data.hSpeed
	state.vSpeed = data.vSpeed
	state.fuel = int(data.fuel)
}

func Move(state *ShuttleState, gene Gene, point Point, time int) Point {
	if state.fuel < gene.power*time {
		time = state.fuel / gene.power
		if time == 0 {
			gene.power = state.fuel
		}
	}
	floatTime := float64(time)
	vA := float64(gene.power)*anglesSin[gene.rotate+90] - g
	hA := float64(gene.power) * anglesCos[gene.rotate+90]
	hSpeed := state.hSpeed + hA*floatTime
	vSpeed := state.vSpeed + vA*floatTime
	result := Point{}
	result.X = point.X + (hSpeed+state.hSpeed)/2*floatTime
	result.Y = point.Y + (vSpeed+state.vSpeed)/2*floatTime
	state.hSpeed = hSpeed
	state.vSpeed = vSpeed
	state.fuel = state.fuel - gene.power*time
	return result
}

func ApplyChromosome(chromosome Chromosome, shuttleData *ShuttleData, ground Ground, time int) ChromosomeData {
	var shuttleState ShuttleState
	shuttleState.Init(shuttleData)
	p1 := Point{shuttleData.x, shuttleData.y}
	p2 := Move(&shuttleState, chromosome[0], p1, time)
	path := make(Path, 2, chromosomeSize)
	path[0] =  p1
	path[1] =  p2
	p, res := IsLandedOrCrashed(ground, p1, p2)
	i:=1
	for ; !res && p2.X <= width && p2.X >= 0 && p2.Y <= height; i++ {
		p1 = p2
		p2 = Move(&shuttleState, chromosome[i], p1, time)
		path = append(path, p2)
		p, res = IsLandedOrCrashed(ground, p1, p2)
	}
	fitnessData := FitnessData{x: p.X,
		hSpeed: shuttleState.hSpeed,
		vSpeed: shuttleState.vSpeed,
		lastRotation: chromosome[i].rotate,
		predRotation: chromosome[i-1].rotate,
	}
	return ChromosomeData{Path: path,
		Chromosome: chromosome,
		LandingPoint: p,
		FitnessData:fitnessData}
}

func ApplyPopulation(shuttleData *ShuttleData, ground Ground) {
	for i:=0;i< populationSize;i++{
		populationData[i]=ApplyChromosome(population[i],shuttleData,ground, deltaTime)
	}
}
