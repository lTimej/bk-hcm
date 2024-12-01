package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
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

	url := "https://36.133.25.49:31015/api/access/admin/subsystem/admin/create/decrypt?AccessKey=49c0ea9b0f2045d9b61cd95e459ecd60&SignatureMethod=HmacSHA1&SignatureNonce=c29b2492ecb94a978a94f2a6b6c5ac64&SignatureVersion=V2.0&Timestamp=2024-12-01T14%3A21%3A04Z&Version=2016-12-05&Signature=0a634e770d826662fc39b1a2f2e636b1122c99aa"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("userId", "CIDC-U-8b63b06d89bd4a9d95250f0ffa9d80b2")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println(req.Header)
	fmt.Println("=======")
	fmt.Println(req.Body)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func main1() {
	dc := "dfs"
	request := &CreateResellerUserRequest{
		DistributionChannel: &dc,
	}
	ConvertRequest(request)
}
