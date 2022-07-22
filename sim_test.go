package TrafficSimEng

import (
	"testing"
	"time"
)

// TestSmallGraph performs a simulation on a graph given in csv format
func TestGraph(t *testing.T) {
	g := Graph{}
	list, err := g.graphFromCSV("./data/abstract_small_graph.csv")
	if err != nil {
		t.Error(err)
		return
	}

	// check set data
	/*for i := 0; i < len(g.roads); i++ {
		fmt.Printf("%v\n", g.roads[i])
	}*/

	//фиксируем текущее время
	start_time := time.Now()

	// allocate a goroutine for each road
	for i := range g.roads {
		go g.roads[i].RoadSim(*list, start_time)
	}

	time.Sleep(sim_time * time.Second)
}
