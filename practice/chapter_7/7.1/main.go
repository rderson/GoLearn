package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w += WordCounter(1)
	}

	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l += LineCounter(1)
	}

	return len(p), nil
}

func main() {
	var w WordCounter
	w.Write([]byte("Ya sosu hui."))
	fmt.Println(w)

	w = 0
	var name = "Yasuo"
	fmt.Fprintf(&w, "Ya ebal tvoi rot, %s", name)
	fmt.Println(w)

	input := "Beneath a sky of scattered clouds, a lone squirrel stares at an unclaimed acorn atop\na mossy rock. Nearby, a tiny robot whirs to life, projecting a hologram of a sunlit\nforest in full bloom, though it's the dead of winter. Somewhere, a piano begins to play—\nsoftly at first, then louder—echoing through empty hallways in an abandoned mansion\nfilled with dusty paintings of people no one remembers. A crow caws, unimpressed.\nThe squirrel, still undecided about the acorn, finally shrugs (as much as a squirrel can)\nand skitters away, leaving the robot to hum softly into the night.\n"
	w = 0
	var l LineCounter
	fmt.Fprintf(&w, "%s", input)
	fmt.Fprintf(&l, "%s", input)
	fmt.Println("Words: ", w)
	fmt.Println("Lines: ", l)
}
