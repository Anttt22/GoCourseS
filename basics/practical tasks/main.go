package main

import (
	"fmt"
)

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	SrviceID int
	status   string
}

func GenerateCheck() []HealthCheck {
	var hc []HealthCheck
	temp := HealthCheck{}
	for i := 0; i < 5; i++ {
		if i%2 != 0 {
			temp = HealthCheck{i, FailStatus}
			hc = append(hc, temp)
		} else {
			temp = HealthCheck{i, PassStatus}
			hc = append(hc, temp)
		}
	}
	return hc
}

func PassID(Gc []HealthCheck) []int {
	var result []int
	for _, HCheck := range Gc {
		switch HCheck.status {
		case "pass":
			result = append(result, HCheck.SrviceID)
		}
	}
	return result
}

func main() {
	Gc := GenerateCheck()
	fmt.Println(Gc)
	fmt.Println("Тут будет выведен идентификатор")
	ID := PassID(Gc)
	fmt.Println(ID)
}
