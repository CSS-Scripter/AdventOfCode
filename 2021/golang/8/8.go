package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"sort"
	"bytes"
	"strconv"
)

type Signal struct {
	value []byte
}

func (s Signal) String() string {
	return string(s.value)
}

type Entry struct {
	signals [10]Signal
	outputs [4]Signal
}

type Decipher struct {
	a byte
	b byte
	c byte
	d byte
	e byte
	f byte
	g byte
}

type SegmentCount struct {
	a int
	b int
	c int
	d int
	e int
	f int
	g int
}

func (d Decipher) String() string {
	return fmt.Sprintf(
		" %s%s%s%s \n" +
		"%s    %s\n" +
		"%s    %s\n" +
		" %s%s%s%s \n" +
		"%s    %s\n" +
		"%s    %s\n" +
		" %s%s%s%s \n",
		string(d.a), string(d.a), string(d.a), string(d.a),
		string(d.b), string(d.c), string(d.b), string(d.c), 
		string(d.d), string(d.d), string(d.d), string(d.d),
		string(d.e), string(d.f), string(d.e), string(d.f), 
		string(d.g), string(d.g), string(d.g), string(d.g),
	)
}

var (
	entries []Entry
)

func main() {
	// one()
	two()
}

func one() {
	prepareInput("input.txt")
	count1478()
}

func count1478() {
	count1 := 0
	count4 := 0
	count7 := 0
	count8 := 0
	for _, entry := range entries {
		for _, output := range entry.outputs {
			switch len(output.value) {
				case 2:
					count1 += 1
					break
				case 3:
					count7 += 1
					break
				case 4:
					count4 += 1
					break
				case 7:
					count8 += 1
					break
			}
		}
	}
	fmt.Printf("1: %d\n4: %d\n7: %d\n8: %d\nsum: %d", count1, count4, count7, count8, (count1+count4+count7+count8))
}

func prepareInput(file string) {
	input, _ := ioutil.ReadFile(file)
	inputStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, entry := range inputStrings {
		entryStrings := strings.Split(entry, " | ")
		signalString := entryStrings[0]
		outputString := entryStrings[1]
		entry := Entry{}
		for i, signal := range strings.Split(signalString, " ") {
			entry.signals[i] = Signal{value: []byte(strings.TrimSpace(signal))}
		}
		for i, output := range strings.Split(outputString, " ") {
			entry.outputs[i] = Signal{value: []byte(strings.TrimSpace(output))}
		}
		entries = append(entries, entry);
	} 
}

func two() {
	prepareInput("input.txt")
	sum := 0
	for _, entry := range entries {
		decipher := getDecipher(entry.signals)
		sum += decipherOutputs(decipher, entry.outputs)
	}
	fmt.Println(sum)
}

func getDecipher(signals [10]Signal) Decipher {
	count := getCountOfSegments(signals)
	decipher := Decipher{}
	decipher.a = decipherA(signals)
	dg := decipherDG(signals, count)
	decipher.d = dg.d
	decipher.g = dg.g
	decipher.b = getSegmentWithOccurances(signals, 6, count)[0]
	decipher.c = getSegmentWithOccurancesNot(signals, 8, count, decipher.a)
	decipher.e = getSegmentWithOccurances(signals, 4, count)[0]
	decipher.f = getSegmentWithOccurances(signals, 9, count)[0]
	return decipher
}

type DG struct {
	d byte
	g byte
}

func decipherDG(signals [10]Signal, count SegmentCount) DG {
	signal4 := find4(signals)
	segmentsDG := getSegmentWithOccurances(signals, 7, count)
	if (doesSignalContainSegment(signal4, segmentsDG[0])) {
		return DG{d: segmentsDG[0], g: segmentsDG[1]}
	}
	return DG{d: segmentsDG[1], g: segmentsDG[0]}
}

func decipherA(signals [10]Signal) byte {
	signal1 := find1(signals)
	signal7 := find7(signals)
	for _, segment := range signal7.value {
		if !doesSignalContainSegment(signal1, segment) {
			return segment
		}
	}
	panic("Segment not found!")
}

func find1(signals [10]Signal) Signal {
	return findSignalWithLength(signals, 2)
}

func find4(signals [10]Signal) Signal {
	return findSignalWithLength(signals, 4)
}

func find7(signals [10]Signal) Signal {
	return findSignalWithLength(signals, 3)
}

func findSignalWithLength(signals [10]Signal, length int) Signal {
	for _, signal := range signals {
		if (len(signal.value) == length) {
			return signal
		}
	}
	panic("Signal not found with this length!");
}

func getSegmentWithOccurancesNot(signals [10]Signal, occurances int, count SegmentCount, not byte) byte {
	segments := getSegmentWithOccurances(signals, occurances, count)
	for _, segment := range(segments) {
		if (segment != not) {
			return segment
		}
	}
	panic("Not found!")
}

func getSegmentWithOccurances(signals [10]Signal, occurances int, count SegmentCount) []byte {
	returnArray := []byte{}
	if (count.a == occurances) { returnArray = append(returnArray, 'a')}
	if (count.b == occurances) { returnArray = append(returnArray, 'b')}
	if (count.c == occurances) { returnArray = append(returnArray, 'c')}
	if (count.d == occurances) { returnArray = append(returnArray, 'd')}
	if (count.e == occurances) { returnArray = append(returnArray, 'e')}
	if (count.f == occurances) { returnArray = append(returnArray, 'f')}
	if (count.g == occurances) { returnArray = append(returnArray, 'g')}
	return returnArray
}

func doesSignalContainSegment(signal Signal, segment byte) bool {
	for _, sigSegment := range signal.value {
		if (sigSegment == segment) {
			return true
		}
	}
	return false
}

func getCountOfSegments(signals [10]Signal) SegmentCount {
	count := SegmentCount{}
	for _, signal := range signals {
		for _, segment := range signal.value {
			switch string(segment) {
				case "a":
					count.a += 1
					break
				case "b":
					count.b += 1
					break
				case "c":
					count.c += 1
					break
				case "d":
					count.d += 1
					break
				case "e":
					count.e += 1
					break
				case "f":
					count.f += 1
					break
				case "g":
					count.g += 1
					break
			}
		}
	}
	return count
}

func decipherOutputs(decipher Decipher, outputs [4]Signal) int {
	sum := "";
	for _, outputSignal := range outputs {
		outputSignal.value = sortBytes(outputSignal.value)
		if (isZero(outputSignal, decipher))  { sum += "0"; continue }
		if (isOne(outputSignal, decipher))   { sum += "1"; continue }
		if (isTwo(outputSignal, decipher))   { sum += "2"; continue }
		if (isThree(outputSignal, decipher)) { sum += "3"; continue }
		if (isFour(outputSignal, decipher))  { sum += "4"; continue }
		if (isFive(outputSignal, decipher))  { sum += "5"; continue }
		if (isSix(outputSignal, decipher))   { sum += "6"; continue }
		if (isSeven(outputSignal, decipher)) { sum += "7"; continue }
		if (isEight(outputSignal, decipher)) { sum += "8"; continue }
		if (isNine(outputSignal, decipher))  { sum += "9"; continue }
		panic("Couldn't decipher!")
	}
	returnVal, _ := strconv.Atoi(sum)
	return returnVal
}

func isZero(signal Signal, decipher Decipher) bool {
	zero := []byte{
		decipher.a,
		decipher.b,
		decipher.c,
		decipher.e,
		decipher.f,
		decipher.g,
	}
	return bytes.Compare(signal.value, sortBytes(zero)) == 0
}

func isOne(signal Signal, decipher Decipher) bool {
	return len(signal.value) == 2
}

func isTwo(signal Signal, decipher Decipher) bool {
	two := []byte{
		decipher.a,
		decipher.c,
		decipher.d,
		decipher.e,
		decipher.g,
	}
	return bytes.Compare(signal.value, sortBytes(two)) == 0
}

func isThree(signal Signal, decipher Decipher) bool {
	three := []byte{
		decipher.a,
		decipher.c,
		decipher.d,
		decipher.f,
		decipher.g,
	}
	return bytes.Compare(signal.value, sortBytes(three)) == 0
}

func isFour(signal Signal, decipher Decipher) bool {
	return len(signal.value) == 4
}

func isFive(signal Signal, decipher Decipher) bool {
	five := []byte{
		decipher.a,
		decipher.b,
		decipher.d,
		decipher.f,
		decipher.g,
	}
	return bytes.Compare(signal.value, sortBytes(five)) == 0
}

func isSix(signal Signal, decipher Decipher) bool {
	six := []byte{
		decipher.a,
		decipher.b,
		decipher.d,
		decipher.e,
		decipher.f,
		decipher.g,
	}
	return bytes.Compare(signal.value, sortBytes(six)) == 0
}

func isSeven(signal Signal, decipher Decipher) bool {
	return len(signal.value) == 3
}

func isEight(signal Signal, decipher Decipher) bool {
	return len(signal.value) == 7
}

func isNine(signal Signal, decipher Decipher) bool {
	nine := []byte{
		decipher.a,
		decipher.b,
		decipher.c,
		decipher.d,
		decipher.f,
		decipher.g,
	}
	return bytes.Compare(signal.value, sortBytes(nine)) == 0
}


func sortBytes(bytes []byte) []byte {
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return bytes
}