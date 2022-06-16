package main

import (
	"bufio"
	"fmt"
	"os"
	"parkingLot/car"
	"parkingLot/parking"
	"strconv"
	"strings"
)

func main() {
	//read input from file (default)
	if len(os.Args) > 1 && os.Args[1] != "" {
		cmdFile, _ := os.Open(os.Args[1])
		defer cmdFile.Close()
		cmdScanner := bufio.NewScanner(cmdFile)
		for cmdScanner.Scan() {
			cmdInput := cmdScanner.Text()
			cmdInput = strings.TrimRight(cmdInput, "\n")
			if cmdInput != "" {
				parseInput(cmdInput)
			}
		}
	}
	//read input from cmd
	reader := bufio.NewReader(os.Stdin)
	for {
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput, "\n")
		if cmdInput != "" {
			parseInput(cmdInput)
		}
	}
}

func parseInput(cmdInput string) {

	cmds := strings.SplitN(cmdInput, " ", 3)
	menu := cmds[0]

	if menu == "create_parking_lot" {
		val, _ := strconv.ParseUint(cmds[1], 0, 64)
		parkingLotObj := parking.NewParking(uint(val))
		fmt.Printf("Created a parking lot with %d slots\n", parkingLotObj.Capacity)
	} else if menu == "park" {
		var carObj *car.Car = car.NewCar(cmds[1], cmds[2])
		sl, err := parking.Get().AddCar(*carObj)
		if err == nil {
			fmt.Printf("Allocated slot number: %d \n", sl.Index)
		} else {
			fmt.Printf("Sorry, parking lot is full")
		}
	} else if menu == "leave" {
		val, _ := strconv.ParseUint(cmds[1], 0, 64)
		err := parking.Get().RemoveCarBySlot(uint(val))
		if err == nil {
			fmt.Printf("Slot number %d is free\n", uint(val))
		}
	} else if menu == "status" {
		var list = []string{fmt.Sprintf("%-12s%-20s%-10s", "Slot No.", "Registration No", "Colour")}
		filledSlots := parking.Get().GetFilledSlots()
		for _, sl := range filledSlots {
			cr := sl.Car
			list = append(list, fmt.Sprintf("%-12v%-20v%-10v", sl.Index, cr.Number, cr.Color))
		}
		output := strings.Join(list, "\n")
		fmt.Println(output)
	} else if menu == "registration_numbers_for_cars_with_colour" {
		var output string
		var list []string
		slots := parking.Get().GetSlotsByCarColor(cmds[1])
		if slots == nil {
			output = "Not found"
		} else {
			for _, sl := range slots {
				list = append(list, sl.GetCarNumber())
			}
			output = strings.Join(list, ", ")
		}
		fmt.Println(output)
	} else if menu == "slot_numbers_for_cars_with_colour" {
		var output string
		var list []string
		slots := parking.Get().GetSlotsByCarColor(cmds[1])
		if slots == nil {
			output = "Not found"
		} else {
			for _, sl := range slots {
				list = append(list, fmt.Sprint(sl.Index))
			}
			output = strings.Join(list, ", ")
		}
		fmt.Println(output)
	} else if menu == "slot_number_for_registration_number" {
		var output string
		slot := parking.Get().GetSlotByCarNumber(cmds[1])
		if slot == nil {
			fmt.Println("Not Found")
		} else {
			output = fmt.Sprint(slot.Index)
			fmt.Println(output)
		}
	}
}
