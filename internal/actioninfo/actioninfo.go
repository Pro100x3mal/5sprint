package actioninfo

import "fmt"

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for i := range dataset {
		err := dp.Parse(dataset[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		str, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(str)
	}

}
