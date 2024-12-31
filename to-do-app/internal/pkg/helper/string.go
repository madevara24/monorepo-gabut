package helper

import "encoding/json"

func ConvertMapStringToString(data map[string][]string) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonData)
}
