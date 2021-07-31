package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"smorty/ditto/move"
)

/*
TODO: support destination folder as argument
TODO: add flag for copy instead of move
TODO: add runmode for single file and optional filename
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Loops to copy is required. Please supply in the following format: <1-10,12,15> etc")
		os.Exit(1)
	}
	args := os.Args[1:]

	/*
		first argument is loops to move
		second argument is optional destination
			create subfolder with current timestamp and save loops as 01.wav, 50.wav, etc
		if env var not set, prompt for it
		if destination is not provided, check for env var and if not found prompt for it
	*/
	loops := args[0]

	target, targetFound := os.LookupEnv("DITTO_LOOP_TARGET")
	source, sourceFound := os.LookupEnv("DITTO_LOOP_SOURCE")

	if !targetFound || !sourceFound {
		log.Fatal("Required env variables not set")
	}

	// TODO: support destination folder as argument
	// if len(args) > 1 {
	// 	dest = args[1]
	// } else if !destFound {
	// 	dest = getDest()
	// }

	tracks, err := parseTracks(loops)

	if err != nil {
		log.Fatal(err)
	}

	move.MoveTracks(tracks, source, target)
	// TODO: validation for input
	// move.MoveTracks(parseTracks(loops))

}

func parseTracks(input string) ([]string, error) {
	split := strings.Split(input, ",")
	output := make([]string, 0)

	for _, v := range split {
		if strings.Contains(v, "-") {
			newRange, err := parseRange(v)
			if err != nil {
				return nil, err
			}
			output = append(output, newRange...)
		} else {
			output = append(output, fmt.Sprintf("%02s", v))
		}
	}

	return output, nil

}

func parseRange(input string) ([]string, error) {
	loopRange := strings.Split(input, "-")

	if len(loopRange) != 2 {
		return nil, errors.New("invalid range")
	}

	start, _ := strconv.Atoi(loopRange[0])
	end, _ := strconv.Atoi(loopRange[1])

	return makeRangeSlice(start, end), nil
}

func makeRangeSlice(start, finish int) []string {
	s := make([]string, finish-start+1)
	for i := range s {
		s[i] = fmt.Sprintf("%02d", start)
		start += 1
	}
	return s
}

// func getDest() string {

// }
