package morse

type Alphabet map[string]string

type Code []string

func Translate(code Code) (string, error) {
	alphabet := getMorseAlphabet()
	result := ""

	for _, currentSymbol := range code {
		result += alphabet[currentSymbol]
	}

	return result, nil
}

func GetCode(inputStr string) Code {
	var code Code

	buffer := ""
	for i := 0; i < len(inputStr); i++ {

		if inputStr[i:i+1] == "." || inputStr[i:i+1] == "-" {

			buffer += inputStr[i : i+1]

		} else {

			code = append(code, buffer)
			buffer = ""

			for i < len(inputStr) {

				if inputStr[i:i+1] != "." && inputStr[i:i+1] != "-" {
					code = append(code, inputStr[i:i+1])
				} else {
					buffer += inputStr[i : i+1]
					break
				}
				i++

			}
		}
	}
	if len(buffer) != 0 {
		code = append(code, buffer)
	}

	return code
}

func getMorseAlphabet() Alphabet {
	return Alphabet{
		".-":   "A",
		"-...": "B",
		"-.-.": "C",
		"-..":  "D",
		".":    "E",
		"..-.": "F",
		"--.":  "G",
		"....": "H",
		"..":   "I",
		".---": "J",
		"-.-":  "K",
		".-..": "L",
		"--":   "M",
		"-.":   "N",
		"---":  "O",
		".--.": "P",
		"--.-": "Q",
		".-.":  "R",
		"...":  "S",
		"-":    "T",
		"..-":  "U",
		"...-": "V",
		".--":  "W",
		"-..-": "X",
		"-.--": "Y",
		"--..": "Z",
		" ":    "",
		"/":    " ",
	}
}
