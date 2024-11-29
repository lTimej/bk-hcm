package main

import (
	"fmt"
	"reflect"
)

type HttpRequestPosition string

const (
	BODY   HttpRequestPosition = "Body"
	QUERY  HttpRequestPosition = "Query"
	PATH   HttpRequestPosition = "Path"
	HEADER HttpRequestPosition = "Header"
)

func ConvertRequest(request interface{}) {
	if request == nil {
		return
	}
	reqType := reflect.TypeOf(request)
	if reqType.Kind() == reflect.Ptr {
		reqType = reqType.Elem()
	}
	reqValue := reflect.ValueOf(request)
	if reqValue.Kind() == reflect.Ptr {
		reqValue = reqValue.Elem()
	}
	var flag = false
	for i := 0; i < reqType.NumField(); i++ {
		fieldType := reqType.Field(i)
		value := reqValue.FieldByName(fieldType.Name)
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue
			}
			value = value.Elem()
			fmt.Println("*******")
		}
		propertyType := fieldType.Type
		if propertyType.Kind() == reflect.Ptr {
			propertyType = propertyType.Elem()
		}
		fmt.Println(propertyType, "====")
		_, flag = propertyType.FieldByName(string(BODY))
		if flag {
			continue
		}
		_, flag = propertyType.FieldByName(string(HEADER))
		if flag {
			continue
		}
		_, flag = propertyType.FieldByName(string(QUERY))
		if flag {
			continue
		}
		_, flag = propertyType.FieldByName(string(PATH))
		if flag {
			continue
		}
	}
}

type CreateResellerUserRequest struct {
	// 经销商平台编码,测试账号申请后会提供
	DistributionChannel *string `json:"distributionChannel,omitempty"`
}

func main() {
	dc := "dfs"
	request := &CreateResellerUserRequest{
		DistributionChannel: &dc,
	}
	ConvertRequest(request)
}
