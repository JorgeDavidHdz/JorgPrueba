package postFix

import "unicode"

func SplitInfix(infix string) ([]string, error) {
	var listCharacter []string
	var element string

	for _, char := range infix {

		if unicode.IsDigit(char) || string(char) == "." {
			element += string(char)
		} else {

			if element != "" {
				listCharacter = append(listCharacter, element)
				element = ""
			}

			listCharacter = append(listCharacter, string(char))
		}

	}

	if element != "" {
		listCharacter = append(listCharacter, element)
	}

	return listCharacter, nil
}
