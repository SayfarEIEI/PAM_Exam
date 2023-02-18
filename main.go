package main

import (
	"log"
	"strconv"
)

func main() {
	log.Print("------------------------------")
	log.Print(timeConversion("03:04:05PM"))
	log.Print("------------------------------")
}
func timeConversion(input string) string {
	result := ""
	clockType := input[len(input)-2:]
	hhStr := input[0:2]
	mmStr := input[3:5]
	ssStr := input[6:8]
	if clockType == "PM" {
		hh, err := strconv.Atoi(hhStr)
		if err != nil {
			return err.Error()
		}
		if hhStr == "12" {
			return hhStr + ":" + mmStr + ":" + ssStr
		} else {
			hhStr = strconv.Itoa(hh + 12)
		}

	}
	if clockType == "AM" {
		_, err := strconv.Atoi(hhStr)
		if err != nil {
			return err.Error()
		}
		if hhStr == "12" {
			hhStr = "00"
		} else {

			return hhStr + ":" + mmStr + ":" + ssStr
		}
	}

	result = hhStr + ":" + mmStr + ":" + ssStr
	return result
}
