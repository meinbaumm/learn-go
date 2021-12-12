package gross

import (
	b "my_module/my_module/basic"
)

func GrossSalary() int {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}
