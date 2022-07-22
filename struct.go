package rain

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

// JsonFileToMap
// JSON文件转换成Map
func JsonFileToMap(jsonFile string) (result interface{}, err error) {
	b, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Println("JsonFileToMap: ", err)
		return result, err
	}

	s := string(b)

	// 转换成Map
	result, err = JsonToMap(s)
	if err != nil {
		log.Println("JsonFileToMap: ", err)
	}

	return result, err
}

// JsonToMap
// JSON字符串转换成Map(或Map切片)
//   返回值为 Map或Map切片
//   参数为空字符串时直接返回空字符串
func JsonToMap(sJson string) (result interface{}, err error) {
	if len(sJson) == 0 || len(strings.TrimSpace(sJson)) == 0 {
		return "", nil
	}

	// 尝试转换成单个JSON对象
	obj, err := jsonObjectToMap(sJson)
	if err == nil {
		return obj, nil
	}

	// 尝试转换成JSON数组
	arr, err := jsonArrayToMap(sJson)
	if err != nil {
		log.Println("JsonToMap: ", err)
		return arr, err
	}

	return arr, nil
}

// JSON(对象)字符串转换成Map
func jsonObjectToMap(sJson string) (result map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(sJson), &result)
	if err != nil {
		log.Println("jsonObjectToMap: ", err)
		return result, err
	}

	return result, nil
}

// JSON(数组)字符串转换成Map
func jsonArrayToMap(sJson string) (result []map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(sJson), &result)
	if err != nil {
		log.Println("jsonArrayToMap: ", err)
		return result, err
	}

	return result, nil
}

// StructToMap
// 将struct转化为map
// 》使用json
func StructToMap(s interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	j, err := json.Marshal(s)
	if err != nil {
		log.Print("StructToMap() -> json.Marshal()")
		return m, err
	}
	err = json.Unmarshal(j, &m)
	if err != nil {
		log.Print("StructToMap() -> json.Unmarshal()")
		return m, err
	}

	return m, nil
}

// ToIfKeyMap
// 将string类型Key的Map转化为interface{}类型Key的Map
// 》interface{}类型Key的Map用于go-flutter插件
func ToIfKeyMap(m map[string]interface{}) (result map[interface{}]interface{}, err error) {
	result = make(map[interface{}]interface{})
	for k, v := range m {
		kindTmp := reflect.ValueOf(v).Kind()

		if kindTmp == reflect.Slice {
			vSlice, ok := v.([]interface{})
			if !ok {
				return result, errors.New("ToIfKeyMap: reflect slice.")
			}

			// 用于存放Map的切片
			var valSlice []interface{}
			for _, subVal := range vSlice {
				// 切片元素为Map类型
				if reflect.ValueOf(subVal).Kind() == reflect.Map {
					itemMap, itemOk := subVal.(map[string]interface{})
					if itemOk {
						ifKeyMap, err := ToIfKeyMap(itemMap)

						if err != nil {
							return nil, err
						}

						valSlice = append(valSlice, ifKeyMap)
					}
				} else { // 非Map类型（认为是普通类型）
					valSlice = append(valSlice, subVal)
				}
			}

			result[k] = valSlice
		} else if kindTmp == reflect.Map {
			tmpMap := make(map[interface{}]interface{})

			for k2, v2 := range v.(map[string]interface{}) {
				tmpMap[k2] = v2
			}

			result[k] = tmpMap
		} else {
			result[k] = v
		}
	}

	return result, nil
}

// ToStringKeyMap
// 将interface{}类型Key的Map转化为string类型Key的Map
// 》interface{}类型Key的Map用于go-flutter插件
func ToStringKeyMap(m map[interface{}]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	for k, v := range m {
		kindTmp := reflect.ValueOf(v).Kind()

		if kindTmp == reflect.Slice {
			vSlice, ok := v.([]interface{})

			if !ok {
				return result, errors.New("ToStringKeyMap: reflect slice.")
			}

			// 用于存放Map的切片
			var valSlice []interface{}
			for _, subVal := range vSlice {
				// 切片元素为Map类型
				if reflect.ValueOf(subVal).Kind() == reflect.Map {
					itemMap, itemOk := subVal.(map[interface{}]interface{})
					if itemOk {
						ifKeyMap, err := ToStringKeyMap(itemMap)

						if err != nil {
							return nil, err
						}

						valSlice = append(valSlice, ifKeyMap)
					}

				} else { // 非Map类型（认为是普通类型）
					valSlice = append(valSlice, subVal)
				}
			}

			result[k.(string)] = valSlice
		} else if kindTmp == reflect.Map {
			tmpMap := make(map[string]interface{})

			for k2, v2 := range v.(map[interface{}]interface{}) {
				tmpMap[k2.(string)] = v2
			}

			result[k.(string)] = tmpMap
		} else {
			result[k.(string)] = v
		}
	}

	return result, nil
}

// StructToIfKeyMap
// 将struct转化为interface{}类型Key的Map
// 》interface{}类型Key的Map用于go-flutter插件
func StructToIfKeyMap(s interface{}) (result map[interface{}]interface{}, err error) {
	// 先转换成string类型Key的Map
	m, err := StructToMap(s)
	if err != nil {
		log.Print("StrctToIfKeyMap: StructToMap()")
		return result, err
	}

	// 再转换成interface类型Key的Map
	result, err = ToIfKeyMap(m)
	if err != nil {
		log.Print("StrctToIfKeyMap: ToIfKeyMap()")
		return result, err
	}

	return result, nil
}

// JsonFileToStruct
// JSON文件内容转struct
//  file: json文件
//  s   : 定义的struct的地址(调用处需要加&)
//
// 使用例：参照README.md
func JsonFileToStruct(file string, s interface{}) error {

	// 读取文件
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Print(err)
		return err
	}

	// 转换成Struct
	err = json.Unmarshal(b, s)
	if err != nil {
		log.Printf("JsonFileToStruct: %v\n", err.Error())
	}

	return nil
}

// MapToStruct
//  map转struct(map->_json->struct)
//    file: json文件
//    s   : 定义的struct的地址(调用处需要加&)
func MapToStruct(m map[string]interface{}, s interface{}) error {
	// map 转换为 JSON 字节
	bJson, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return err
	}

	// 转换成Struct
	err = json.Unmarshal(bJson, s)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

// func ToMap(tests []interface) {

// 	l := make([]map[string]interface{}, 0)
// 	for _, t := range tests {
// 		elem := reflect.ValueOf(&t).Elem()
// 		relType := elem.Type()

// 		m := make(map[string]interface{}, 1)
// 		for i := 0; i < relType.NumField(); i++ {
// 			m[relType.Field(i).Name] = elem.Field(i).Interface()
// 		}

// 		l = append(l, m)
// 	}
// }
