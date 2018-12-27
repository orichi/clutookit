package morse_code

import (
	"errors"
	"time"
)

// alphaNumToMorse is a mapping of Alpha numeric characters to Morse code
var alphaNumToMorse = map[string]string{
	"A":  ".-",
	"B":  "-...",
	"C":  "-.-.",
	"D":  "-..",
	"E":  ".",
	"F":  "..-.",
	"G":  "--.",
	"H":  "....",
	"I":  "..",
	"J":  ".---",
	"K":  "-.-",
	"L":  ".-..",
	"M":  "--",
	"N":  "-.",
	"O":  "---",
	"P":  ".--.",
	"Q":  "--.-",
	"R":  ".-.",
	"S":  "...",
	"T":  "-",
	"U":  "..-",
	"V":  "...-",
	"W":  ".--",
	"X":  "-..-",
	"Y":  "-.--",
	"Z":  "--..",
	"1":  ".----",
	"2":  "..---",
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	"0":  "-----",
	".":  ".-.-.-",  // period
	":":  "---...",  // colon
	",":  "--..--",  // comma
	";":  "-.-.-",   // semicolon
	"?":  "..--..",  // question
	"=":  "-...-",   // equals
	"'":  ".----.",  // apostrophe
	"/":  "-..-.",   // slash
	"!":  "-.-.--",  // exclamation
	"-":  "-....-",  // dash
	"_":  "..--.-",  // underline
	"\"": ".-..-.",  // quotation marks
	"(":  "-.--.",   // parenthesis (open)
	")":  "-.--.-",  // parenthesis (close)
	"()": "-.--.-",  // parentheses
	"$":  "...-..-", // dollar
	"&":  ".-...",   // ampersand
	"@":  ".--.-.",  // at
	"+":  ".-.-.",   // plus
	"Á":  ".--.-",   // A with acute accent
	"Ä":  ".-.-",    // A with diaeresis
	"É":  "..-..",   // E with acute accent
	"Ñ":  "--.--",   // N with tilde
	"Ö":  "---.",    // O with diaeresis
	"Ü":  "..--",    // U with diaeresis
	" ":  ".......", // word interval
}

// morseToAlphaNum is a mapping of Alpha numeric characters to Morse code
var morseToAlphaNum map[string]string
func MorseTable() *map[string]string{
	return &alphaNumToMorse
}
func ReverseMorseTable() *map[string]string{
	return &morseToAlphaNum
}

func init() {
	morseToAlphaNum = map[string]string{}
	for k, v := range alphaNumToMorse {
		morseToAlphaNum[v] = k
	}
}

type MorseResult struct {
	data string
	originData string
	timestamp time.Time
	err error
}

func (result *MorseResult) Data() string{
	if result.err != nil{
		return result.originData + "=>\t"+result.err.Error()
	}else{
		return result.originData + "=>\t"+result.data
	}
}

func (result *MorseResult) Timestamp() string{
	return result.timestamp.Format("2006-01-02 15:04:05")
}

func checkInput(inputData string) (bool, error){
	bytesData := []byte(inputData)
	for i := 0; i < len(inputData); i++ {
		charI := string(bytesData[i])
		if alphaNumToMorse[charI] == ""{
			return false, errors.New("字符串不合法: "+charI)
		}
	}
	return true, nil
}

