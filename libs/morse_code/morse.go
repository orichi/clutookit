package morse_code

import (
	"errors"
	"strings"
	"time"
	"regexp"
)



func Encoder(data string) *MorseResult{
	encodedData := make([]string, len(data))
	bytesData := []byte(data)
	for i:=0;i<len(data);i++ {
		charI := string(bytesData[i])
		if alphaNumToMorse[charI] == ""{
			return &MorseResult{err: errors.New("字符不合法: "+charI), originData: data, timestamp:time.Now()}
			break
		}else{
			encodedData[i] = alphaNumToMorse[charI]
		}
	}
	return &MorseResult{data: strings.Join(encodedData, " "), originData: data, timestamp: time.Now()}
}

func Decoder(data string) *MorseResult{
	re, _ := regexp.Compile(".-[.\\-\\ ]*");
	data1 := re.FindString(data) //获取正则匹配的字符串
	splitedData := strings.Split(data1, " ")
	decodedData := make([]string, len(splitedData))
	for i := 0; i < len(splitedData); i++ {
		if morseToAlphaNum[splitedData[i]] == ""{
			return &MorseResult{err: errors.New("字符不合法: "+splitedData[i]), originData: data1, timestamp:time.Now()}
			break
		}else{
			decodedData[i] = morseToAlphaNum[splitedData[i]]
		}
	}
	return &MorseResult{originData: data1, data: strings.Join(decodedData, ""),timestamp:time.Now()}
}
