package TrafficSimEng

const car_len = 5 // vehicle length

type Road struct {
	id     int
	source int
	target int
	weight int
	light  Light // traffic light
	next   []int // adjacent roads available for crossing
	cars   []Car
}

func (r *Road) RoadSim(list map[int]int) {
	// divide the road into cells
	cells := r.division()

	// there is a traffic light at the end of the road, so we simulate a transport of zero length in the last cell
	if r.light.red != 0 {
		var light Car
		light.l = 0
		cells[len(cells)-1] = true
		r.cars = append(r.cars, light)
	}

	// number of cars to be generated on the current road
	gen := list[r.id]

	added := false
	// generate a given number of vehicles
	for i := 0; i < gen; i++ {
		var car Car
		car.l = car_len
		// assign last id
		car.id = len(r.cars)

		for !added {
			if !cells[0] {
				car.CarIn(r)
				added = true
			}
			r.update()
		}
		// need to select the number of cells needed by the vehicle
		if !cells[0] {
			car.CarIn(r)
		}
		r.cars = append(r.cars, car)

	}

}

// Division returns cell array
func (r *Road) division() []bool {
	cells := make([]bool, r.weight)
	return cells
}

// Update updates the simulation
func (r *Road) update() {
	for _, v := range r.cars {
		v.IDM()
	}
}
