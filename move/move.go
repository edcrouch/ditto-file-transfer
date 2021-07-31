package move

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func MoveTracks(tracks []string, source string, target string) error {
	if _, err := os.Stat(source); err != nil {
		return err
	}

	if _, err := os.Stat(target); err != nil {
		return err
	}

	if !strings.HasSuffix(source, "/") {
		source += "/"
	}

	if !strings.HasSuffix(target, "/") {
		target += "/"
	}

	target = target + time.Now().Format("2006-01-02 15_04_05") + "/"

	err := os.Mkdir(strings.TrimSuffix(target, "/"), os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	for _, track := range tracks {
		// TODO: implement with goroutines
		err := CopyTrack(track, source, target, true)

		if err != nil {
			// log.Fatal(err)
			fmt.Println("Skipping " + track)
		}
	}

	fmt.Println("All done")
	return nil
}

func CopyTrack(track, source, target string, remove bool) error {
	sourcePath := source + track + "track/LOOP.wav"
	sourceFile, err := os.Open(sourcePath)

	if err != nil {
		// log.Fatal(err)
		return err
	}

	defer sourceFile.Close()

	targetFile, err := os.Create(target + track + ".wav")

	if err != nil {
		log.Fatal(err)
	}

	defer targetFile.Close()

	_, err = io.Copy(targetFile, sourceFile)

	if err != nil {
		log.Fatal(err)
	}

	if remove {
		err = os.Remove(sourcePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Copied " + track)

	return nil
}
