package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	hexToBin = map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	versionSum int
)

func main() {
	versionSum = 0
	value, _ := readPacket(prepareInput("input.txt"))
	fmt.Printf("One: %d\n", versionSum)
	fmt.Printf("Two: %d\n", value)
}

func prepareInput(file string) string {
	input, _ := ioutil.ReadFile(file)
	hexString := strings.TrimSpace(string(input))
	return hexToBinFunc(hexString)
}

func hexToBinFunc(hex string) string {
	bin := ""
	for _, hexChar := range hex {
		bin += hexToBin[hexChar]
	}
	return bin
}

func getVersionAndTypeId(pckt string) (int, int) {
	version := binToInt(pckt[0:3])
	typeId := binToInt(pckt[3:6])
	return version, typeId
}

func binToInt(bin string) int {
	if i, err := strconv.ParseInt(bin, 2, 64); err != nil {
		panic(err)
	} else {
		return int(i)
	}
}

func readPacket(pckt string) (int, int) {
	version, typeId := getVersionAndTypeId(pckt)
	versionSum += version
	switch typeId {
	case 4:
		return literalPacket(pckt)
	default:
		return operator(pckt)
	}
}

func literalPacket(pcktBin string) (int, int) {
	pcktBin = pcktBin[6:]
	lastPacket := false
	contentBin := ""
	length := 6
	for i := 0; !lastPacket; i += 5 {
		length += 5
		if pcktBin[i] == '0' {
			lastPacket = true
		}
		contentBin += pcktBin[i+1 : i+5]
	}
	return binToInt(contentBin), length
}

func operator(pcktBin string) (int, int) {
	_, typeId := getVersionAndTypeId(pcktBin)
	lengthTypeId := pcktBin[6]
	pcktBin = pcktBin[7:]
	if lengthTypeId == '0' {
		subPacketsBinLength := binToInt(pcktBin[:15])
		pcktBin = pcktBin[15:]
		pcktsBin := pcktBin[:subPacketsBinLength]
		totalLength := 0
		subPacketValues := []int{}
		for totalLength < subPacketsBinLength {
			value, length := readPacket(pcktsBin[totalLength:])
			totalLength += length
			subPacketValues = append(subPacketValues, value)
		}
		return determinePacketValue(typeId, subPacketValues), 22 + subPacketsBinLength //header + packetlength + subpackets
	}
	if lengthTypeId == '1' {
		amountOfPackets := binToInt(pcktBin[:11])
		pcktBin = pcktBin[11:]
		subPacketValues := make([]int, amountOfPackets)
		lengthSum := 18
		for i := 0; i < amountOfPackets; i++ {
			value, length := readPacket(pcktBin)
			pcktBin = pcktBin[length:]
			lengthSum += length
			subPacketValues[i] = value
		}

		return determinePacketValue(typeId, subPacketValues), lengthSum
	}
	panic("unsupported operation")
}

func determinePacketValue(typeId int, values []int) int {
	switch typeId {
	case 0:
		return sum(values)
	case 1:
		return multiply(values)
	case 2:
		return minimum(values)
	case 3:
		return maximum(values)
	case 5:
		return greater(values)
	case 6:
		return less(values)
	case 7:
		return equal(values)
	}
	panic("Unsupported Type ID: " + fmt.Sprint(typeId))
}

func sum(values []int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func multiply(values []int) int {
	total := 1
	for _, val := range values {
		total *= val
	}
	return total
}

func minimum(values []int) int {
	min := -1
	for _, val := range values {
		if min == -1 || val < min {
			min = val
		}
	}
	return min
}

func maximum(values []int) int {
	max := 0
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

func greater(values []int) int {
	if values[0] > values[1] {
		return 1
	}
	return 0
}

func less(values []int) int {
	if values[0] < values[1] {
		return 1
	}
	return 0
}

func equal(values []int) int {
	if values[0] == values[1] {
		return 1
	}
	return 0
}
