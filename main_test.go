package main

import (
	"reflect"
	"testing"
    "runtime"
    "strings"
)

func GetTestFunctionName() string {
    function, _, _, _ := runtime.Caller(1)
    fullName := runtime.FuncForPC(function).Name()

    parts := strings.Split(fullName, ".")
    return parts[len(parts) - 1]
}

func TestDoOverridesSimple(t *testing.T) {
    base := map[string]interface{}{
        "a": 1,
        "b": 2,
    }
    overrides := map[string]interface{}{
        "a": 2,
    }
    expected := map[string]interface{}{
        "a": 2,
        "b": 2,
    }

    DoOverrides(base, overrides, nil)
    if !reflect.DeepEqual(base, expected) {
        t.Errorf("[Test: %s] Expected value: %v, got: %v", GetTestFunctionName(), expected, base)
    }
}
