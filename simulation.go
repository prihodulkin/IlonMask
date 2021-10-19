package main

type ShuttleState struct {
	hSpeed float64
	vSpeed float64
	fuel   int
	power  int
	rotate int
}

func (state *ShuttleState) Init(data *ShuttleData) {
	state.hSpeed = data.hSpeed
	state.vSpeed = data.vSpeed
	state.fuel = int(data.fuel)
	state.power = data.power
	state.rotate = data.rotate
}

func (state *ShuttleState) ChangeRotate(dRotate int) {
	rotate := state.rotate + dRotate
	if rotate > 90 {
		state.rotate = 90
	} else if rotate < -90 {
		state.rotate = -90
	} else{
		state.rotate=rotate
	}
}

func (state *ShuttleState) ChangePower(dPower int) {
	power:= state.power + dPower
	if power > 4 {
		state.power = 4
	} else if power < 0 {
		state.power = 0
	} else{
		state.power = power
	}
}

func Move(state *ShuttleState,  point Point, time int) Point {
	if state.fuel < state.power*time {
		time = state.fuel / state.power
		if time == 0 {
			state.power = state.fuel
		}
	}
	floatTime := float64(time)
	vA := float64(state.power)*anglesSin[state.rotate+90] - g
	hA := float64(state.power) * anglesCos[state.rotate+90]
	hSpeed := state.hSpeed + hA*floatTime
	vSpeed := state.vSpeed + vA*floatTime
	result := Point{}
	result.X = point.X + (hSpeed+state.hSpeed)/2*floatTime
	result.Y = point.Y + (vSpeed+state.vSpeed)/2*floatTime
	state.hSpeed = hSpeed
	state.vSpeed = vSpeed
	state.fuel = state.fuel - state.power*time
	return result
}

func ApplyChromosome(chromosome Chromosome, shuttleData *ShuttleData, ground Ground, time int) ChromosomeData {
	var shuttleState ShuttleState
	shuttleState.Init(shuttleData)
	p1 := Point{shuttleData.x, shuttleData.y}
	p2 := Move(&shuttleState,  p1, time)
	path := make(Path, 2, chromosomeSize)
	path[0] = p1
	path[1] = p2
	p, res := IsLandedOrCrashed(ground, p1, p2)
	i := 1
	for ; !res && p2.X <= width && p2.X >= 0 && p2.Y <= height; i++ {
		p1 = p2
		shuttleState.ChangeRotate(chromosome[i-1].dRotate)
		shuttleState.ChangePower(chromosome[i-1].dPower)
		p2 = Move(&shuttleState, p1, time)
		path = append(path, p2)
		p, res = IsLandedOrCrashed(ground, p1, p2)
	}
	fitnessData := FitnessData{x: p.X,
		hSpeed:       shuttleState.hSpeed,
		vSpeed:       shuttleState.vSpeed,
		lastRotation: shuttleState.rotate,
		predRotation: shuttleState.rotate-chromosome[i-2].dRotate,
	}
	return ChromosomeData{Path: path,
		Chromosome:   chromosome,
		LandingPoint: p,
		FitnessData:  fitnessData}
}

func ApplyPopulation(shuttleData *ShuttleData, ground Ground) {
	for i := 0; i < populationSize; i++ {
		populationData[i] = ApplyChromosome(populationData[i].Chromosome, shuttleData, ground, dTime)
	}
}
