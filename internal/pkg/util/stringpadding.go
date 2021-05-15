package util

import (
	"strconv"
	"strings"
)

const totalLength = 10

func PaddingLeadingZero(num uint) string {
	numString := strconv.FormatUint(uint64(num), 10)
	sb := strings.Builder{}
	sb.Grow(totalLength)
	for i := 0; i < totalLength-len(numString); i++ {
		sb.WriteRune('0')
	}
	sb.WriteString(numString)
	return sb.String()
}

func RemoveLeadingZero(str string) (uint, error) {
	str = strings.TrimLeft(str, "0")
	num, err := strconv.ParseUint(str, 10, 32)
	return uint(num), err
}
