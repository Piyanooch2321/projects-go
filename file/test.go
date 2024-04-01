package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Header struct {
	id   int
	name string
}

func main() {
	// อ่านไฟล์ JSON
	data, err := os.ReadFile("startFileNew.json")
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการอ่านไฟล์:", err)
		return
	}

	// แปลง JSON เป็น slice ของ interface{}
	var mainArray []interface{}
	if err := json.Unmarshal(data, &mainArray); err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการแปลง JSON:", err)
		return
	}

	// เรียกใช้ฟังก์ชันเพื่อค้นหาคอลัมน์ที่มีข้อมูล "o" ที่ตำแหน่งที่ 0
	columns := findColumnsWithValue(mainArray[2].([]interface{}), "f")
	var headers []Header
	for _, column := range columns {
		colArray := column.([]interface{})
		num := int(colArray[1].(float64))

		for i := 0; i < len(colArray[2].([]interface{})); i++ {
			gid := num + i
			val := colArray[2].([]interface{})[i].(string)
			headers = append(headers, Header{id: gid, name: val})
		}

	}
	// ex {0 type}
	fmt.Println(headers)
}

// ฟังก์ชันเพื่อค้นหาคอลัมน์ที่มีข้อมูล "o" ที่ตำแหน่งที่ 0
func findColumnsWithValue(columnsArray []interface{}, value string) []interface{} {
	var result []interface{}
	for _, column := range columnsArray {
		if colArray, ok := column.([]interface{}); ok && len(colArray) > 0 && colArray[0] == value {
			result = append(result, column)
		}
	}
	return result
}
