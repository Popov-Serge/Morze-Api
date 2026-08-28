package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertText(s string) string {
	str := strings.ReplaceAll(s, ".", "")
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, " ", "")

	if str == "" {
		return morse.ToText(s)
	}

	return morse.ToMorse(s)
}
