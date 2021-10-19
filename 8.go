package main

import "fmt"

func main() {
	num := "321-"
	num = "2147483648"
	ret := myAtoi(num)
	fmt.Println(ret)
}

type _FSM struct {
	number   int
	positive bool
	state    string
}

func NewFSM() *_FSM {
	ret := &_FSM{0, true, "begin"}
	return ret
}

func (fsm *_FSM) nextStep(c byte) bool {
	transMap := map[string][4]string{
		"begin":  [4]string{"begin", "signed", "number", "nextIdx"},
		"signed": [4]string{"nextIdx", "nextIdx", "number", "nextIdx"},
		"number": [4]string{"nextIdx", "nextIdx", "number", "nextIdx"},
		"nextIdx":    [4]string{"nextIdx", "nextIdx", "nextIdx", "nextIdx"},
	}

	t := fsm.getType(c)
	fsm.state = transMap[fsm.state][t]
	d := int(c - '0')
	if fsm.state == "number" {
		if fsm.positive && (fsm.number > 214748364||fsm.number == 214748364 && d > 7){
			fsm.number = 2147483647
			return false
		}
		if !fsm.positive && (fsm.number > 214748364||fsm.number == 214748364 && d >=8){
			fsm.number = -2147483648
			return false
		}
		fsm.number *= 10
		fsm.number += d
	}
	if fsm.state == "nextIdx" {
		return false
	}
	return true
}
func (fsm *_FSM) getType(c byte) int {
	if c == ' ' {
		return 0
	} else if c == '+' {
		return 1
	} else if c == '-' {
		if fsm.state == "begin" {
			fsm.positive = false
		}
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	} else {
		return 3
	}
}
func (fsm *_FSM) GetNumber() int {
	if fsm.number < 0 {
		return fsm.number
	}
	if fsm.positive {
		return fsm.number
	} else {
		return -fsm.number
	}
}

func myAtoi(str string) int {
	fsm := NewFSM()
	for i := 0; i < len(str); i++ {
		if !fsm.nextStep(str[i]) {
			break
		}
	}
	return fsm.GetNumber()
}
