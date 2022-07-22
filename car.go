package TrafficSimEng

import (
	"fmt"
	"math"
)

const (
	car_len = 5.0  // vehicle length
	a       = 0.3  // acceleration
	b       = 3.0  // slowdown
	s0      = 2.0  // minimum space
	v0      = 30.0 // free speed
	T       = 1.5  // advance time
	exp     = 4.0  // acceleration exponent
	yellow  = 2    // yellow traffic light time
)

type Car struct {
	id     int
	length float64
	pos    float64 //position
	vel    float64 //velocity
	sp     float64 //space
}

// CarIn adds vehicle to the road
func (c *Car) CarIn(r *Road) {
	c.space(&r.cars[len(r.cars)-1])
	// there are no vehicles and traffic light on the road yet or the last vehicle has already driven to a safe distance or there is a traffic light
	if len(r.cars) == 0 || c.sp >= s0 {
		fmt.Printf("Расстояние: %v\n", c.sp)
		c.id = len(r.cars)
		r.cars = append(r.cars, *c)
	}
}

// IDM returns the updated car data
func (c *Car) IDM(lead *Car) {
	// s*
	s := s0 + c.vel*T + c.vel*v0/(2*math.Sqrt(a*b))
	// IDM acceleration
	idm := a * (1 - c.vel/v0*exp - math.Pow(s/c.sp, 2))
	c.velocity(idm)
	c.position(idm)
	c.space(lead)
}

// space returns the distance between lead vehicle and current
func (c *Car) space(lead *Car) {
	c.sp = lead.pos - lead.length - c.pos
}

// velocity counts the vehicle position at the moment after the update
func (c *Car) velocity(idm float64) {
	c.vel += idm * upd_time
}

// position counts the vehicle position at the moment after the update
func (c *Car) position(idm float64) {
	c.pos += c.vel*upd_time + 1/2*idm*math.Pow(upd_time, 2)
}

// CarOut deletes vehicle
func (c *Car) CarOut() {
	//the path for the car determines randomly
}
