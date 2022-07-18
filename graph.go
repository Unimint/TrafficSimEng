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

func (g *Graph) graphFromCSV(fname string) error {

	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(file))

	reader.Comma = ';'

	_, err = reader.Read()
	if err != nil {
		return err
	}

	//  road id
	i := 1

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		// source vertex ID
		fv_id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}
		// target vertex ID
		tv_id, err := strconv.Atoi(line[1])
		if err != nil {
			return err
		}

		// line[2] contains direction
		// all roads are unidirectional by default, so there is no check for directionality

		// the weight of the road as a number of cells
		w, err := strconv.Atoi(line[3])
		if err != nil {
			return err
		}

		var light Light

		// traffic light at the end of the road
		l, err := strconv.ParseBool(line[4])
		if err != nil {
			return err
		}

		if l {
			// the green traffic light time in seconds
			gl, err := strconv.Atoi(line[5])
			if err != nil {
				return err
			}
			//the red traffic light time in seconds
			rl, err := strconv.Atoi(line[6])
			if err != nil {
				return err
			}

			light = Light{green: gl, red: rl}
		}

		// array of available neighbors
		neighbors := []int{}
		j := 7

		for line[j] != "" {
			neighbor, err := strconv.Atoi(line[j])
			if err != nil {
				return err
			}
			neighbors = append(neighbors, neighbor)
			j++
		}
		g.roads = append(g.roads, Road{id: i, source: fv_id, target: tv_id, weight: w, next: neighbors, light: light})
		i++
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
