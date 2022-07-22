package TrafficSimEng

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Graph struct {
	roads []Road
}

// graphFromCSV initializes the Graph structure using the data extracted from the csv file
// all roads are unidirectional by default, so there is no check for directionality
func (g *Graph) graphFromCSV(fname string) (*map[int]int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	reader.Comma = ';'

	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// road id
	i := 1

	// define a list of generations
	list := make(map[int]int)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		// source vertex ID
		sv_id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}

		// target vertex ID
		tv_id, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		// the number of generated cars at the roads source vertex
		n, err := strconv.Atoi(line[2])
		if err != nil {
			return nil, err
		}
		list[i] = n

		// the weight of the road
		w, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			return nil, err
		}

		var light Light

		// traffic light at the end of the road
		l, err := strconv.ParseBool(line[4])
		if err != nil {
			return nil, err
		}

		if l {
			// the green traffic light time in seconds
			gl, err := strconv.Atoi(line[5])
			if err != nil {
				return nil, err
			}
			//the red traffic light time in seconds
			rl, err := strconv.Atoi(line[6])
			if err != nil {
				return nil, err
			}

			light = Light{green: gl, red: rl}
		}

		// array of available neighbors
		neighbors := []int{}
		//start from the eighth value
		j := 7

		for line[j] != "" {
			neighbor, err := strconv.Atoi(line[j])
			if err != nil {
				return nil, err
			}
			neighbors = append(neighbors, neighbor)
			j++
		}
		g.roads = append(g.roads, Road{id: i, source: sv_id, target: tv_id, weight: w, light: light, next: neighbors})
		i++
	}
	return &list, file.Close()
}
