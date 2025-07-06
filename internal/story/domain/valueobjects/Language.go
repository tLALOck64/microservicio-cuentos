package valueobjects

import "fmt"

type Language string

const (
	Tzeltal  Language = "tzeltal"
	Zapoteco Language = "zapoteco"
	Maya     Language = "maya"
)

func NewLanguage(value string) (Language, error) {
	lang := Language(value)

	if !lang.IsValid() {
		return "", fmt.Errorf("lengua no soportada: %s. Lenguas Válidas: tzeltal, zapoteco, maya", value)
	}
	return lang, nil
}

func (l Language) IsValid() bool {
	validLanguges := []Language{
		Tzeltal, Zapoteco, Maya,
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
