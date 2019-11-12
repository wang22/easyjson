package easyjson

import "encoding/json"

type EasyJSON struct {
	JSONMap map[string]interface{}
}

func (e *EasyJSON) ContainsKey(key string) bool {
	return e.JSONMap[key] != nil
}

func (e *EasyJSON) GetObject(key string) *EasyJSON {
	return &EasyJSON{e.JSONMap[key].(map[string]interface{})}
}

func (e *EasyJSON) GetArrayObject(key string) []*EasyJSON {
	array := e.JSONMap[key].([]interface{})
	rst := []*EasyJSON{}
	for _, v := range array {
		rst = append(rst, &EasyJSON{v.(map[string]interface{})})
	}
	return rst
}

func (e *EasyJSON) GetString(key string) string {
	return e.JSONMap[key].(string)
}

func (e *EasyJSON) GetInt(key string) int {
	return int(e.GetFloat64(key))
}

func (e *EasyJSON) GetFloat64(key string) float64 {
	return e.JSONMap[key].(float64)
}

func ParseJSON(jsonData []byte) *EasyJSON {
	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)
	return &EasyJSON{data}
}
