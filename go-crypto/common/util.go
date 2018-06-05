package main
import (
	"encoding/json"
	"strconv"
	"time"
	"fmt"
)

func GenDate() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func GenTimestamp() string {
	t := time.Now()
	nanos := t.UnixNano()
	millis := nanos / 1000000 //ms 13
	return strconv.FormatInt(millis,10)
}

func Serialize(dat map[string]interface{}) string {
	str, err := json.Marshal(dat)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func Deserialize(str string) map[string]interface{} {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(str), &dat)
        if  err != nil {
		panic(err)
	}
	return dat
}


func main() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
        dat := Deserialize(jsonStr)
        fmt.Println("===json 2 map===")
        fmt.Println(dat)
        fmt.Println(dat["host"])
        fmt.Println("===map 2 json===")
	str := Serialize(dat)
	fmt.Println(str)
	fmt.Println(GenTimestamp())
	fmt.Println(GenDate())
}

