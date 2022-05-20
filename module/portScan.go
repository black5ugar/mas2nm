package module

import (
	"fmt"
	"mas2nm/model"
	"strconv"
)

func PortScan(input string) error {
	// mas := model.NewMasscan()
	// mas.Scan(input)
	masOutput, err := model.GetMasscanResult()
	if err != nil {
		return err
	}

	nmapInput := make(map[string][]string)
	for _, result := range masOutput {
		nmapInput[result.IP] = append(nmapInput[result.IP], strconv.Itoa(result.Ports[0].Port))
	}
	fmt.Println(nmapInput)

	return nil
}
