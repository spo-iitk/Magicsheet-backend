package sync

import "strings"

func getProgramDepartment(id uint) (program, department string) {
	value, ok := ProgramDepartmentMap[id]

	if !ok {
		return "Unknown", "Unknown"
	}

	parts := strings.SplitN(value, "-", 2)

	if len(parts) != 2 {
		return value, ""
	}

	return parts[0], parts[1]
}
