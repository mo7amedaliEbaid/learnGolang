package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
	// b.tip = tip
}




func getbillInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getbillInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getbillInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getbillInput("Item name: ", reader)
		price, _ := getbillInput("Item price: ", reader)

		fmt.Println(name, price)
	case "t":
		tip, _ := getbillInput("Enter tip amount ($): ", reader)

		fmt.Println(tip)
	case "s":
		fmt.Println("you chose to save the bill")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}


func main() {
	
	mybill := createBill()
	promptOptions(mybill)

}



