package clc

func Min(a int, b int ) int{
	if a > b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}