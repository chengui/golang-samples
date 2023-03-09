package strconv

import (
	"fmt"
	"strings"
)

type Currency int64

func FormatCurrency(c Currency) string {
	s := fmt.Sprintf("%d", c)
	var sb strings.Builder
	i, n := 0, len(s)
	if s[0] == '+' || s[0] == '-' {
		sb.WriteByte(s[0])
		i++
	}
	if (n-i)%3 == 0 {
		sb.WriteString(s[i : i+3])
		i += 3
	} else {
		sb.WriteString(s[i : i+(n-i)%3])
		i += (n - i) % 3
	}
	for ; i < n; i += 3 {
		sb.WriteByte(',')
		sb.WriteString(s[i : i+3])
	}
	return sb.String()
}

func ParseCurrency(s string) (Currency, error) {
	return 0, nil
}
