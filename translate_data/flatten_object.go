package translatedata

import (
	"encoding/json"
)

func FlattenObject(jsonStr string, prefix string) (map[string]interface{}, error) {
	var obj map[string]interface{}
	// Unmarshal the JSON string into a map
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return nil, err // Return an error if the JSON cannot be parsed
	}

	// Call the recursive flatten function on the parsed object
	flatMap := make(map[string]interface{})
	flatten("", obj, flatMap)
	return flatMap, nil
}

func flatten(prefix string, nestedMap map[string]interface{}, flatMap map[string]interface{}) {
	for key, value := range nestedMap {
		if subMap, ok := value.(map[string]interface{}); ok {
			newPrefix := key
			if prefix != "" {
				newPrefix = prefix + "." + key
			}
			flatten(newPrefix, subMap, flatMap)
		} else {
			if prefix != "" {
				flatMap[prefix+"."+key] = value
			} else {
				flatMap[key] = value
			}
		}
	}
}
