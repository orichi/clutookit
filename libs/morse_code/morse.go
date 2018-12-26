package morse_code

import (
	"errors"
	"strings"
	"time"
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
