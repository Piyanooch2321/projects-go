package main

import (
	"encoding/json"
	"fmt"
	"os"

	Linq "github.com/ahmetb/go-linq/v3"
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

	f := findColumnsWithValue(mainArray[2].([]interface{}), "f")
	header := setHeaderName(f)
	// fmt.Println("--------", header[0].id)
	o := findColumnsWithValue(mainArray[2].([]interface{}), "o")
	valueo := getValue(o, header)
	fmt.Println("---V----", valueo)
	m := findColumnsWithValue(mainArray[2].([]interface{}), "m")
	valuem := getValue(m, header)
	fmt.Println("----M----", valuem)

	// เรียกใช้ฟังก์ชันเพื่อค้นหาคอลัมน์ที่มีข้อมูล "o" ที่ตำแหน่งที่ 0

}

func getValue(columns []interface{}, header []Header) []interface{} {
	resultMap := make(map[string]interface{})
	result := make([]interface{}, 0, len(resultMap))

	for _, column := range columns {
		column := column.([]interface{})
		for i := 0; i < len(column)-1; i += 2 {
			if i%2 == 0 {

				h := Linq.From(header).Where(func(c interface{}) bool {
					return c.(Header).id == int(column[i].(float64))
				}).First().(Header)

				// fmt.Println("--------", h.name, column[i], column[i+1])
				resultMap[h.name] = column[i+1]
			}
		}

		result = append(result, resultMap)
	}

	return result
}

func setHeaderName(columns []interface{}) []Header {
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

	return headers

}

// ฟังก์ชันเพื่อค้นหาคอลัมน์ที่มีข้อมูล "o" ที่ตำแหน่งที่ 0
func findColumnsWithValue(columnsArray []interface{}, value string) []interface{} {
	var result []interface{}
	for _, column := range columnsArray {
		colArray, ok := column.([]interface{})

		if ok && len(colArray) > 0 && colArray[0] == value {
			result = append(result, column)
		} else if ok && len(colArray) > 0 && colArray[1] == value {
			result = append(result, column)

		}

	}
	return result
}
