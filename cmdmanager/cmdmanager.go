package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmdm *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm your price with ENTER.")
	var prices []string

	for {
		var price string

		fmt.Print("Print: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmdm *CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
