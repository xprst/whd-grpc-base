package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xprst/whd-grpc-base/util"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type ConfigEngine struct {
	data map[string]interface{}
}

func (c *ConfigEngine) Load(path string) error {
	ext := c.guessFileType(path)
	if ext == "" {
		return errors.New("load config " + path + " fail. Please insure your config file is *.js or *.json.")
	}

	bExists, err := util.PathExists(path)
	if bExists {
		return c.loadFromJsonFile(path)
	} else {
		return errors.New(fmt.Sprintf("cannot find config file in path: %s. error: %v", path,err))
	}


}

// guessFileType 判断配置文件名是否为js 或 json格式
func (c *ConfigEngine) guessFileType(path string) string {
	s := strings.Split(path, ".")
	ext := s[len(s)-1]
	switch ext {
	case "json", "js":
		return "json"
	}

	return ""
}

// loadFromJsonFile 加载配置文件
func (c *ConfigEngine) loadFromJsonFile(path string) error {
	yamlS, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	err := json.Unmarshal(yamlS, &c.data)
	if err != nil {
		return errors.New(fmt.Sprintf("load json fail. %v", err))
	}

	return nil
}

// Get 从配置文件中获取值,name 可以是层级的，以 . 分隔
func (c *ConfigEngine) Get(name string) interface{} {
	path := strings.Split(name, ".")
	data := c.data
	for key, value := range path {
		v, ok := data[value]
		if !ok {
			break
		}
		if (key + 1) == len(path) {
			return v
		}
		if reflect.TypeOf(v).String() == "map[string]interface {}" {
			data = v.(map[string]interface{})
		}
	}

	return nil
}

// GetString 从配置文件中获取string类型的值
func (c *ConfigEngine) GetString(name string) string {
	value := c.Get(name)
	switch value := value.(type) {
	case string:
		return value
	case bool, float64, int:
		return fmt.Sprint(value)
	default:
		return ""
	}
}

// GetInt 从配置文件中获取int类型的值
func (c *ConfigEngine) GetInt(name string) int {
	value := c.Get(name)
	switch value := value.(type) {
	case string:
		i, _ := strconv.Atoi(value)
		return i
	case int:
		return value
	case bool:
		if value {
			return 1
		}
		return 0
	case float64:
		return int(value)
	default:
		return 0
	}
}

// GetBool 从配置文件中获取bool类型的值
func (c *ConfigEngine) GetBool(name string) bool {
	value := c.Get(name)
	switch value := value.(type) {
	case string:
		b, _ := strconv.ParseBool(value)
		return b
	case int:
		if value != 0 {
			return true
		}
		return false
	case bool:
		return value
	case float64:
		if value != 0.0 {
			return true
		}
		return false
	default:
		return false
	}
}

// GetFloat64 从配置文件中获取Float64类型的值
func (c *ConfigEngine) GetFloat64(name string) float64 {
	value := c.Get(name)
	switch value := value.(type) {
	case string:
		f, _ := strconv.ParseFloat(value, 64)
		return f
	case int:
		return float64(value)
	case bool:
		if value {
			return float64(1)
		}
		return float64(0)
	case float64:
		return value
	default:
		return float64(0)
	}
}

// GetStruct 从配置文件中获取Struct类型的值,这里的struct是自定义的
func (c *ConfigEngine) GetStruct(name string, s interface{}) interface{} {
	value := c.Get(name)
	switch value.(type) {
	case string:
		c.setField(s, name, value)
	case map[string]interface{}:
		c.mapToStruct(value.(map[string]interface{}), s)
	}

	return s
}

func (c *ConfigEngine) mapToStruct(m map[string]interface{}, s interface{}) interface{} {
	for key, value := range m {
		c.setField(s, key, value)
	}

	return s
}

// setField 字段反射赋值
func (c *ConfigEngine) setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.Indirect(reflect.ValueOf(obj))
	if structValue.Kind() == reflect.Map {
		mapType := reflect.TypeOf(obj)
		elemType := mapType.Elem()
		newValue := reflect.New(elemType).Elem()

		switch value := value.(type) {
		case string:
			newValue.SetString(value)
		case map[string]interface{}:
			c.mapToStruct(value, newValue.Addr().Interface())
		}

		structValue.SetMapIndex(reflect.ValueOf(name), newValue)

		return nil
	}

	// 获取Json标签,如果有Json标签，则将要赋值的字段换成类型真正的字段名
	structType := reflect.TypeOf(obj).Elem()
	for i := 0; i < structType.NumField(); i++ {
		if name == structType.Field(i).Tag.Get("json") {
			name = structType.Field(i).Name
		}
	}

	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)

	if structFieldValue.Kind() == reflect.Struct && val.Kind() == reflect.Map {
		vInterface := val.Interface()
		switch vInterface.(type) {
		case map[string]interface{}:
			for key, value := range vInterface.(map[string]interface{}) {
				c.setField(structFieldValue.Addr().Interface(), key, value)
			}
		}
	} else {
		if structFieldType != val.Type() {
			return errors.New("provided value type didn't match obj field type")
		}

		structFieldValue.Set(val)
	}

	return nil
}
