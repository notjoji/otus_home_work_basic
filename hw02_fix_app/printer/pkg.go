package printer

import (
	"fmt"

	"github.com/notjoji/otus_home_work_basic/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		str := fmt.Sprintf("%d %d %s %d", staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(str)
	}
}
