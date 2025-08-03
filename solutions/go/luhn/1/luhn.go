package luhn

import (
    "strings"
    "strconv"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	sm := 0
	isEven := len(id)%2 == 0 
	for i := 0; i < len(id); i++ {
		d, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		if (i%2 == 0) == isEven { 
			if d*2 > 9 {
				sm += (d * 2) - 9
			} else {
				sm += d * 2
			}
		} else {
			sm += d
		}
	}

	return sm%10 == 0
}
