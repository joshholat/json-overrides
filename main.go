package main

import (
    "os"
    "reflect"

    "io/ioutil"
    "encoding/json"

    "github.com/icza/dyno"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    argsWithoutSelf := os.Args[1:]
    if len(argsWithoutSelf) < 2 {
        panic("Two parameters must be provided, the base JSON and the override JSON")
    }

    // Get the JSON from our files
    baseJson, err := ioutil.ReadFile(argsWithoutSelf[0])
    check(err)
    overridesJson, err := ioutil.ReadFile(argsWithoutSelf[1])
    check(err)

    // Unmarshal the JSON to editable types
    var base interface{}
    err = json.Unmarshal([]byte(baseJson), &base);
    check(err)
    var overrides interface{}
    err = json.Unmarshal([]byte(overridesJson), &overrides)
    check(err)

    // Actually apply the overrides
    DoOverrides(base, overrides, nil);

    // Write our output with the new values
    overridenJson, err := json.MarshalIndent(base, "", "\t")
    check(err)

    outputFile := "output.json"
    if 3 == len(argsWithoutSelf) {
        outputFile = argsWithoutSelf[2]
    }
    err = ioutil.WriteFile(outputFile, []byte(overridenJson), 0644)
    check(err)
}

func DoOverrides(base interface{}, overrides interface{}, path []interface{}) {
    switch root := overrides.(type) {
    	case []interface{}:
    		for key, _ := range root {
               newPath := append(path, key)
               DoOverrides(base, root[key], newPath)
    		}

    	case map[string]interface{}:
    		for key, value := range root {
                typeName := reflect.TypeOf(value).Name()
                if "string" == typeName || "float64" == typeName || "int" == typeName {
                    newPath := append(path, key)
                    err := dyno.Set(base, value, newPath...)
                    check(err)
                } else {
                    newPath := append(path, key)
                    DoOverrides(base, root[key], newPath)
                }
    		}
    }
}
