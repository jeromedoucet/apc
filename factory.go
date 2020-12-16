package apc

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Configuration of a Factory
type Configuration struct {
	machinesCapacities []int
	nbChipset          int
}

// Factory contains machines that produce shipset
type Factory struct {
	conf Configuration
}

// NewConf build a factory configuration using an input file
func NewConf(inputFile string) (conf Configuration, err error) {
	var content []byte

	// we already know the number of line. No needs for a bufio.ReadString() here.
	content, err = ioutil.ReadFile(inputFile)
	if err != nil {
		return
	}

	sc := bufio.NewScanner(strings.NewReader(string(content)))
	i := 0
	for sc.Scan() {
		switch i {
		case 1:
			var capacities []int
			capacities, err = readMachineCaps(sc.Text())
			if err != nil {
				return
			}
			conf.machinesCapacities = capacities
		case 2:
			var nbChipset int
			nbChipset, err = strconv.Atoi(sc.Text())
			if err != nil {
				return
			}
			conf.nbChipset = nbChipset
		}
		i++
	}

	if i != 2 {
		err = errors.New(fmt.Sprintf("Wrong configuration file format : got %d lines instead of 3", i+1))
	}

	return
}

func readMachineCaps(line string) (capacities []int, err error) {
	machines := strings.Split(line, " ")
	capacities = make([]int, len(machines), len(machines))

	for i, machine := range machines {
		var machineCap int
		machineCap, err = strconv.Atoi(machine)

		if err != nil {
			return
		}

		capacities[i] = machineCap
	}

	return
}

// NewFactory build a new factory using the given configuration
func NewFactory(conf Configuration) Factory {
	return Factory{conf}
}

// Optimums do print Optimums configuration of a Factory
func (f Factory) Optimums(outputFile string) {

}

func computePossibilites(sum, remainingCapacity, res map[int][]int) {
	for i := 0; i < len(remainingCapacity); i++ {
		consumption = sum + remainingCapacity[i]
		if _, ok := m[consumption]; !ok {
			m[consumption] = make([]int)
		}
		m[consumption] = append(m[consumption])
	}
}
