package main

import (
	"flag"
	"fmt"
	"small_tool/libs/pwd_generator"
	"os"
	"log"
	"small_tool/libs/morse_code"
)

func main() {
	funcPtr := flag.String("f", "pwd1", "pwd1: easy password\npwd2: normal password\npwd3: hard password\nmorse")
	dataPtr := flag.String("s", "", "data be morse encode")
	pwdLenPtr := flag.Int("len", 8, "default is 8, max is 62")
	flag.Parse()
	switch *funcPtr {
	case "pwd1":
		{
			result1 := pwd_generator.EasyGenerate(*pwdLenPtr)
			showResult(result1)
		}
	case "pwd2":
		{
			result1 := pwd_generator.NormalGenerate(*pwdLenPtr)
			showResult(result1)
		}
	case "pwd3":
		{
			result1 := pwd_generator.ComplexGenerate(*pwdLenPtr)
			showResult(result1)
		}
	case "morse":
		{
			if *dataPtr == ""{
				fmt.Printf("%v\n", *morse_code.MorseTable())
				fmt.Printf("%v\n", *morse_code.ReverseMorseTable())
			}else{
				result := morse_code.Encoder(*dataPtr)
				showResult(result)
				result2 := morse_code.Decoder(result.Data())
				showResult(result2)
			}
		}
	}
}

type ToolResult interface{
	Data() string
	Timestamp() string
}

//os.O_APPEND 追加写入
func showResult(result ToolResult){
	homePath  := []byte(os.Getenv("HOME"))
	homePath = append(homePath, []byte("/.small_tool.log")...)
	f, err := os.OpenFile(string(homePath), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f,"%v: %v\n", result.Timestamp(), result.Data())
	fmt.Printf("time:\t%s\ndata:\t%s\n", result.Timestamp(), result.Data())
}
