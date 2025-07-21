package valueobjects

import "fmt"

type Language string

const (
	Tseltal  Language = "tseltal"
	Zapoteco Language = "zapoteco"
	Maya     Language = "maya"
)

func NewLanguage(value string) (Language, error) {
	lang := Language(value)

	if !lang.IsValid() {
		return "", fmt.Errorf("lengua no soportada: %s. Lenguas Válidas: tseltal, zapoteco, maya", value)
	}
	return lang, nil
}

func (l Language) IsValid() bool {
	validLanguges := []Language{
		Tseltal, Zapoteco, Maya,
	}

	for _, valid := range validLanguges {
		if l == valid {
			return true
		}
	}

	return false
}

func (l Language) String() string {
	return string(l)
}
