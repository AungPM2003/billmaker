package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	value, err := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(value)), err
}

func promptOptions(b *bill) {
	value, _ := getInput("a - (add item) | t - (add tip) | s - (save bill): ")
	switch value {
	case "a":
		fmt.Println("adding item")
		item, _ := getInput("Item: ")
		price, _ := getInput("Price: ")
		parsePrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addingItems(item, parsePrice)
		promptOptions(b)
	case "t":
		fmt.Println("adding tip")
		tip, _ := getInput("Tip: ")
		parseTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updatingTip(parseTip)
		promptOptions(b)
	case "s":
		fmt.Println("saving bill")
		data := []byte(b.format())
		err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
		if err != nil {
			fmt.Println("oops something went wrong")
			promptOptions(b)
		}
		fmt.Println("File was saved")
	default:
		fmt.Println("This is not a valid option")
		promptOptions(b)
	}
}

func main() {
	billName, _ := getInput("Create a new bill: ")
	myBill := newBill(billName)
	promptOptions(&myBill)
}
