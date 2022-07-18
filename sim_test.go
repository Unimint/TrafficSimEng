package TrafficSimEng

import (
	"fmt"
	"testing"
	"time"
)

// Simulation time in seconds
const simtime = 2

// TestSmallGraph performs a simulation on a graph given in csv format
func TestGraph(t *testing.T) {
	g := Graph{}
	err := g.graphFromCSV("./data/abstract_small_graph.csv")
	if err != nil {
		t.Error(err)
		return
	}

	// set data check
	for i := 0; i < len(g.roads); i++ {
		fmt.Printf("%v\n", g.roads[i])
	}

	// define a list of generations
	list := make(map[int]int)
	list[1] = 1
	list[3] = 1

	// allocate a goroutine for each road
	for i := range g.roads {
		go g.roads[i].RoadSim(list)
	}

	time.Sleep(simtime * time.Second)
}
