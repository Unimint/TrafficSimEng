package TrafficSimEng

const (
	cell_len = 5   // cell length
	t        = 1   // update time
	v0       = 30  // free speed
	T        = 1.5 // advance time
	s0       = 2.0 // minimum space
	a        = 0.3 // minimum space
	b        = 3.0 // slowdown
	exp      = 4   // acceleration exponent
)

// yellow traffic light time
const wait_time = 2

type Car struct {
	id int
	x  float64 // position
	v  float64 // velocity
	a  float64 // acceleration
	l  int     //length
	d  Road    // destination
}

// CarIn adds vehicle to the road
func (c *Car) CarIn(r *Road) {

}

// CarOut deletes vehicle
func (c *Car) CarOut() {
	// if there are available adjacent roads and the vehicle has not reached the goal - regenerate the vehicle
	// if there are no available adjacent roads or the vehicle has reached the goal - delete the vehicle
}

// IDM returns the updated car data
func (c *Car) IDM() {
	// IDM
	// speed
	// pos
	// space
}

// Space returns the vehicle speed at the moment after the update
func (c *Car) speed() {

}

// Pos returns the vehicle position at the moment after the update
func (c *Car) pos() {

}

// Space returns the distance between lead vehicle and current
func (c *Car) space() {

}
