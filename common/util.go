package main
import (
	"encoding/json"
	"fmt"
)


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
}

