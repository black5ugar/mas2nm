package module

import "mas2nm/model"

func PortScan(input string) error {
	mas := model.NewMasscan()
	mas.Scan(input)

	return nil
}
