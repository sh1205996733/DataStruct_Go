package utils

func Hash(key any) int {
	return key.(int) % 10
}
