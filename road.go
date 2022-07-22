package TrafficSimEng

import (
	"time"
)

const (
	// simulation time in seconds
	sim_time = 5
	// update time in seconds
	upd_time = 1
)

type Road struct {
	id     int
	source int
	target int
	weight float64
	light  Light // traffic light
	cars   []Car
	next   []int
}

func (r *Road) RoadSim(list map[int]int, last_update time.Time) {
	// there is a traffic light at the end of the road, so we simulate a transport of zero length in the last cell
	if r.light.red != 0 {
		light := Car{length: 0, pos: float64(r.weight)}
		r.cars = append(r.cars, light)
	}
	r.generator(list, last_update)
	// update the road until a given amount of time has passed
	for i := 0; i < sim_time/upd_time; i++ {
		// stop the goroutine until update time has passed
		time.Sleep((upd_time - time.Since(last_update)) * time.Second)
		r.update()
		last_update = time.Now()

		// check set data
		/*if r.id == 1 || r.id == 3 {
			fmt.Printf("%d: дорога %d обновлена -> расстояние: %f\n", i, r.id, r.cars[1].pos)
		}*/
	}
}

// generator generates specified number of cars on a specific road
func (r *Road) generator(list map[int]int, last_update time.Time) {
	added := false
	for i := 0; i < list[r.id]; i++ {
		car := Car{length: car_len, vel: v0}
		for !added {
			// check if there are cars turning onto this road or cars waiting to enter

			// try to add a car
			car.CarIn(r)
			// check if the car is added by its id
			if car.id != 0 {
				added = true
			}
			// stop the goroutine until update time has passed
			time.Sleep((upd_time - time.Since(last_update)) * time.Second)
			r.update()
			last_update = time.Now()
		}
	}
}

// update updates the state of the simulation for a fixed amount of update time
func (r *Road) update() {
	for i := 0; i < len(r.cars); i++ {
		// there is a traffic light at the end of the road, so we start from the next vehicle
		if r.cars[i].length == 0 {
			continue
		}
		// when we read the first vehicle, use the distant virtual vehicle as the leader
		if i == 0 {
			car := Car{pos: r.weight}
			r.cars[i].IDM(&car)
		} else {
			r.cars[i].IDM(&r.cars[i-1])
		}
	}
}
