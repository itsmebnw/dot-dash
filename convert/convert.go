package convert

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type load struct {
	entries []entry
}

type entry struct {
	value  string
	status string
}

type slot struct {
	data []string
}

func Run(path *string) error {
	p := *path

	dt := defineTranslations()

	load, err := process(p, dt)
	if err != nil {
		log.Print(err)
	}
	printResults(load)

	return nil
}

func process(path string, dt map[string]int) (*load, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error when opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	runecount, rowcount, slotcount := 0, 0, 0

	var symbols []string

	slots := make([]slot, 9)

	entry := &entry{value: "", status: ""}

	load := &load{entries: nil}

	for scanner.Scan() {
		s := scanner.Text()

		symbols = append(symbols, s)

		if rowcount == 3 {
			t, err := translate(slots, dt)

			enter(entry, t, load, err)

			runecount, slotcount, symbols, slots, entry = resetEntry()
			rowcount = 0

			continue
		}

		if runecount == 27 {
			rowcount++

			runecount, slotcount, symbols = resetRow()

			continue
		}

		if (runecount+1)%3 == 0 {
			slotcount = endSlotSegment(symbols, slots, slotcount)
			symbols = nil
		}

		runecount++
	}

	if len(slots) == 9 {
		t, err := translate(slots, dt)
		enter(entry, t, load, err)
	}

	return load, nil
}

func endSlotSegment(symbols []string, slots []slot, slotcount int) int {
	ss := strings.Join(symbols, "")
	slots[slotcount].data = append(slots[slotcount].data, ss)
	slotcount++

	return slotcount
}

func enter(entry *entry, t []string, load *load, err error) {
	entry.value = strings.Join(t, "")
	if err != nil {
		entry.status = "ILL"
	} else if !checksum(t) {
		entry.status = "ERR"
	}

	load.entries = append(load.entries, *entry)
}

func resetRow() (runecount, slotcount int, symbols []string) {
	return 0, 0, nil
}

func resetEntry() (runecount, slotcount int, symbols []string, slots []slot, ent *entry) {
	x, y, z := resetRow()
	e := &entry{value: "", status: ""}

	return x, y, z, make([]slot, 9), e
}

func translate(slots []slot, translations map[string]int) ([]string, error) {
	var translation []string

	var err error

	for _, slot := range slots {
		key := strings.Join(slot.data, "")
		if v, ok := translations[key]; ok {
			translation = append(translation, strconv.Itoa(v))
		} else {
			translation = append(translation, "?")
			err = fmt.Errorf("key not found %v with length: %v", key, len(key))
		}
	}

	return translation, err
}

func defineTranslations() map[string]int {
	// making these show in the code with line breaks is difficult with go's gofmt enforcing tabs
	return map[string]int{
		"     |  |": 1,
		" _  _||_ ": 2,
		" _  _| _|": 3,
		"   |_|  |": 4,
		" _ |_  _|": 5,
		" _ |_ |_|": 6,
		" _   |  |": 7,
		" _ |_||_|": 8,
		" _ |_| _|": 9,
		" _ | ||_|": 0,
	}
}

func checksum(t []string) bool {
	var sum int

	for i, x := range t {
		xi, _ := strconv.Atoi(x)

		val := xi * (9 - i)
		sum += val
	}

	return sum%11 == 0
}

func printResults(l *load) {
	for _, e := range l.entries {
		log.Printf("%v %v \n", e.value, e.status)
	}
}

// story 5
// use keys from definitions
// compare symbols with loop
// if more then 1 miss skip
// try each 1 miss vs checksum

// compare input vs each pattern
func findAlts(input string, translations map[string]int) {
	for key := range translations {
		offs := 0
		for _, symbol := range key {

			for _, v := range input {
				if v != symbol {
					offs++
				}
			}

			if offs > 1 {
				continue
			}
		}

		if offs == 1 {
			// this is a possible alt

		}
	}
}
