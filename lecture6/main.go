package main

import (
	"encoding/json"
	"fmt"
)

func generateGoCodeFromJSON(jsonStr string) string {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return err.Error()
	}

	return generateCodeRecursive(data, "result")
}

func generateCodeRecursive(data interface{}, variableName string) string {
	switch v := data.(type) {
	case map[string]interface{}:
		code := variableName + " := map[string]interface{}{\n"
		for key, value := range v {
			code += fmt.Sprintf(`"%s": %s,`, key, generateCodeRecursive(value, key))
		}
		code += "}\n"
		return code
	case []interface{}:
		code := variableName + " := []interface{}{\n"
		for index, item := range v {
			code += generateCodeRecursive(item, fmt.Sprintf("%s[%d]", variableName, index))
		}
		code += "}\n"
		return code
	default:
		return fmt.Sprintf(`"%s"`, v)
	}
}

func main() {
	jsonStr := `{
		"key1": "value1",
		"key2": {
			"subkey1": "subvalue1",
			"subkey2": [1, 2, 3]
		}
	}`

	goCode := generateGoCodeFromJSON(jsonStr)
	fmt.Println(goCode)
}
