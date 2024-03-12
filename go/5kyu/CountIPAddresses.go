package kyu

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*

* With input "10.0.0.0", "10.0.0.50"  => return   50
ambil 2 data,dan setiap titik key di check apakah sama
  - data1,2key,key1 dan key2 = 10.0
  - data2,2key,key1 dan key2 = 10.0

kalo sama silahkan lanjut dan tidak di hitung
  - data1,2key,key3 dan key4 = 0.0
  - data2,2key,key3 dan key4 = 0.50

kalo berbeda pada akhir silahkan di hitung
0+50 = 50

* With input "10.0.0.0", "10.0.1.0"   => return  256
ambil 2 data,dan setiap titik key di check apakah sama
  - data1,2key,key1 dan key2 = 10.0
  - data2,2key,key1 dan key2 = 10.0

kalo sama silahkan lanjut dan tidak di hitung
  - data1,2key,key3 dan key4 = 0.0
  - data2,2key,key3 dan key4 = 1.0

kalo berbeda pada akhir silahkan di hitung
key 3 = (0+1)*256 = 256
key 4 = 0+0 = 0

* With input "20.0.0.10", "20.0.1.0"  => return  246
ambil 2 data,dan setiap titik key di check apakah sama
  - data1,2key,key1 dan key2 = 20.0
  - data2,2key,key1 dan key2 = 20.0

kalo sama silahkan lanjut dan tidak di hitung
  - data1,2key,key3 dan key4 = 0.10
  - data2,2key,key3 dan key4 = 1.0

kalo berbeda pada akhir silahkan di hitung
key 3 = (0+1)*256 = 256
key 4 = 10+0 = 10

key3-key4
256-10=246
*/
func IpsBetween(start, end string) int {
	var resultCountBoolData = false
	var noteCountBoolData string
	var resultIncrementData int
	data1 := strings.Split(start, ".")
	data2 := strings.Split(end, ".")
	result := map[string]int{}

	for i := 0; i < len(data1); i++ {
		tempConvertToIntData1, _ := strconv.Atoi(data1[i])
		tempConvertToIntData2, _ := strconv.Atoi(data2[i])
		if tempConvertToIntData1 > tempConvertToIntData2 {
			result["result"+strconv.Itoa(i)] = tempConvertToIntData1 - tempConvertToIntData2
		} else {
			result["result"+strconv.Itoa(i)] = tempConvertToIntData2 - tempConvertToIntData1
		}
	}

	result["result0"] *= 16777216
	result["result1"] *= 65536
	result["result2"] *= 256
	for i := 0; i < len(result); i++ {
		if !resultCountBoolData && result["result"+strconv.Itoa(i)] != 0 {
			resultCountBoolData = !resultCountBoolData
			noteCountBoolData = "result" + strconv.Itoa(i)
			resultIncrementData = i
			break
		}
	}
	if len(result)-1 == resultIncrementData {
		return result[noteCountBoolData]
	}
	fmt.Println(result)
	fmt.Println("result" + strconv.Itoa(resultIncrementData))
	fmt.Println(result["result"+strconv.Itoa(resultIncrementData)])
	for i := resultIncrementData; i < len(result)-1; i++ {
		fmt.Println(i + 1)
		result["result"+strconv.Itoa(resultIncrementData)] -= result["result"+strconv.Itoa(i+1)]
		fmt.Println(result["result"+strconv.Itoa(resultIncrementData)])
	}
	return result["result"+strconv.Itoa(resultIncrementData)]
}
