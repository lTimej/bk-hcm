/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package rest

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"hcm/pkg/criteria/constant"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/kit"
	"hcm/pkg/logs"

	"github.com/emicklei/go-restful/v3"
)

// Contexts request context.
type Contexts struct {
	Kit            *kit.Kit
	Request        *restful.Request
	resp           *restful.Response
	respStatusCode int

	// request meta info
	bizID string
}

// RequestBody 返回拷贝的body内容
func (c *Contexts) RequestBody() ([]byte, error) {
	byt, err := ioutil.ReadAll(c.Request.Request.Body)
	if err != nil {
		return nil, err
	}

	c.Request.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byt))

	return byt, nil
}

// ReDecodeInto 解析Body到结构体中，且把body内容重新写入Body.
func (c *Contexts) ReDecodeInto(to interface{}) error {
	byt, err := ioutil.ReadAll(c.Request.Request.Body)
	if err != nil {
		return err
	}

	c.Request.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byt))

	err = json.Unmarshal(byt, to)
	if err != nil {
		logs.ErrorDepthf(1, "decode request body failed, err: %s, rid: %s", err.Error(), c.Kit.Rid)
		return errf.NewFromErr(errf.InvalidParameter, err)
	}

	return nil
}

// DecodeInto decode request body to a struct, if failed, then return the
// response with an error
func (c *Contexts) DecodeInto(to interface{}) error {
	err := json.NewDecoder(c.Request.Request.Body).Decode(to)
	if err != nil {
		logs.ErrorDepthf(1, "decode request body failed, err: %s, rid: %s", err.Error(), c.Kit.Rid)
		return errf.NewFromErr(errf.InvalidParameter, err)
	}

	return nil
}

// PathParameter get path parameter value by its name.
func (c *Contexts) PathParameter(name string) PathParam {
	return PathParam(c.Request.PathParameter(name))
}

// DecodePathParamInto decode path parameter value into its actual type.
func (c *Contexts) DecodePathParamInto(name string, to interface{}) error {
	param := c.PathParameter(name)
	err := json.Unmarshal([]byte(param), &to)
	if err != nil {
		logs.ErrorDepthf(1, "decode path parameter %s failed, err: %v, rid: %s", param, err, c.Kit.Rid)
		return errf.NewFromErr(errf.InvalidParameter, err)
	}
	return nil
}

// PathParam defines path parameter value type.
type PathParam string

// String convert path parameter to string.
func (p PathParam) String() string {
	return string(p)
}

// Uint64 convert path parameter to uint64.
func (p PathParam) Uint64() (uint64, error) {
	value, err := strconv.ParseUint(string(p), 10, 64)
	if err != nil {
		logs.ErrorDepthf(1, "decode path parameter %s failed, err: %v", p, err)
		return 0, errf.NewFromErr(errf.InvalidParameter, err)
	}
	return value, nil
}

// Int64 convert path parameter to int64.
func (p PathParam) Int64() (int64, error) {
	value, err := strconv.ParseInt(string(p), 10, 64)
	if err != nil {
		logs.ErrorDepthf(1, "decode path parameter %s failed, err: %v", p, err)
		return 0, errf.NewFromErr(errf.InvalidParameter, err)
	}
	return value, nil
}

// WithStatusCode set the response status header code
func (c *Contexts) WithStatusCode(statusCode int) *Contexts {
	c.respStatusCode = statusCode
	return c
}

func (c *Contexts) respFile(resp FileDownloadResp) {
	c.resp.AddHeader("Content-Type", resp.ContentType())
	c.resp.AddHeader("Content-Disposition", resp.ContentDisposition())

	filepath := resp.Filepath()
	file, err := os.Open(filepath)
	defer func() {
		// 把文件流返回给前端后，将业务中创建的临时文件删除，业务如果需要保留文件，应当将相关的文件另外保存一份
		err := os.Remove(filepath)
		if err != nil {
			logs.ErrorDepthf(1, "remove file failed, filepath: %s, err: %s, rid: %s",
				filepath, err.Error(), c.Kit.Rid)
			return
		}
	}()
	defer file.Close()
	if err != nil {
		logs.ErrorDepthf(1, "open file failed, err: %s, rid: %s", err.Error(), c.Kit.Rid)
		return
	}
	// 使用bufio.NewReader创建一个新的Reader，用于流式读取
	reader := bufio.NewReader(file)

	var buffer [4096]byte // 定义一个4096字节的缓冲区，大小可以根据实际情况调整
	for {
		// 从reader中读取数据到缓冲区
		n, err := reader.Read(buffer[:cap(buffer)])
		if err != nil {
			if err == io.EOF {
				// 到达文件末尾，退出循环
				break
			}
			// 读取数据发生错误，记录日志并返回
			logs.ErrorDepthf(1, "read file failed, err: %s, rid: %s", err.Error(), c.Kit.Rid)
			return
		}

		// 将缓冲区中的数据写入HTTP响应
		_, writeErr := c.resp.ResponseWriter.Write(buffer[:n])
		if writeErr != nil {
			// 写入响应时发生错误，记录日志并返回
			logs.ErrorDepthf(1, "write response failed, err: %s, rid: %s", writeErr.Error(), c.Kit.Rid)
			return
		}
		if f, ok := c.resp.ResponseWriter.(http.Flusher); ok {
			f.Flush()
		}
	}

	// 如果使用HTTP/1.1协议并且没有发送Content-Length头，可能需要调用Flush以确保数据被发送
	if f, ok := c.resp.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// respEntity response request with a success response.
func (c *Contexts) respEntity(data interface{}) {
	if c.respStatusCode != 0 {
		c.resp.WriteHeader(c.respStatusCode)
	}

	c.resp.Header().Set(constant.RidKey, c.Kit.Rid)

	if fileResp, ok := data.(FileDownloadResp); ok {
		c.respFile(fileResp)
		return
	}

	c.resp.AddHeader(restful.HEADER_ContentType, restful.MIME_JSON)
	resp := &Response{
		Code:    errf.OK,
		Message: "",
		Data:    data,
	}

	if err := json.NewEncoder(c.resp.ResponseWriter).Encode(resp); err != nil {
		logs.ErrorDepthf(1, "do response failed, err: %s, rid: %s", err.Error(), c.Kit.Rid)
		return
	}

	return
}

// respError response request with error response.
func (c *Contexts) respError(err error) {
	if c.respStatusCode > 0 {
		c.resp.WriteHeader(c.respStatusCode)
	}

	if c.Kit != nil {
		c.resp.Header().Set(constant.RidKey, c.Kit.Rid)
	}
	c.resp.AddHeader(restful.HEADER_ContentType, restful.MIME_JSON)
	resp := errf.Error(err).Resp()

	encodeErr := json.NewEncoder(c.resp.ResponseWriter).Encode(resp)
	if encodeErr != nil {
		logs.ErrorDepthf(1, "response with error failed, err: %v, rid: %s", encodeErr, c.Kit.Rid)
		return
	}

	return
}

// respErrorWithEntity response request with error response.
func (c *Contexts) respErrorWithEntity(data interface{}, err error) {
	if c.respStatusCode != 0 {
		c.resp.WriteHeader(c.respStatusCode)
	}

	c.resp.Header().Set(constant.RidKey, c.Kit.Rid)
	c.resp.AddHeader(restful.HEADER_ContentType, restful.MIME_JSON)

	parsedErr := errf.Error(err)
	resp := &Response{
		Code:    parsedErr.Code,
		Message: parsedErr.Message,
		Data:    data,
	}

	if err := json.NewEncoder(c.resp.ResponseWriter).Encode(resp); err != nil {
		logs.ErrorDepthf(1, "do response failed, err: %s, rid: %s", err.Error(), c.Kit.Rid)
		return
	}

	return
}

// DecodeHeader decode request header to a struct, if failed, then return the
// response with an error
func (c *Contexts) DecodeHeader(to interface{}) error {
	val := reflect.ValueOf(to).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		headerName := field.Name
		if headerName == "" {
			continue
		}

		headerValue := c.Request.Request.Header.Get(headerName)
		if headerValue == "" {
			continue
		}

		fieldVal := val.Field(i)
		if !fieldVal.CanSet() {
			continue
		}

		switch fieldVal.Kind() {
		case reflect.String:
			fieldVal.SetString(headerValue)
		default:
			return fmt.Errorf("unsupported field type: %s", fieldVal.Kind())
		}
	}

	return nil
}

// DecodeHeader decode request header to a struct, if failed, then return the
// response with an error
func (c *Contexts) DecodeQuery(to interface{}) error {
	st := reflect.TypeOf(to).Elem()  // 获取type
	sv := reflect.ValueOf(to).Elem() // 获取value
	m := c.Request.Request.URL.Query()
	// 只需要遍历全部的go struct field ，复杂度O(N), url.Values操作是0(1)
	for i := 0; i < st.NumField(); i++ {
		// 从type里面获取tag
		name := strings.Split(st.Field(i).Tag.Get("json"), ",")[0]
		v := m.Get(name)
		if len(v) == 0 {
			//fmt.Println("not found value in urlValues: key=", name)
			continue
		}
		val := reflect.ValueOf(v)
		// 类型匹配
		if st.Field(i).Type == val.Type() {
			sv.Field(i).Set(val)
		} else if sv.Field(i).Kind() == reflect.Ptr {
			sv.Field(i).Set(val.Elem())
		} else {
			return fmt.Errorf("Provided value type didn't match obj field type: %v != %v", st.Field(i).Type, val.Type())
		}
		// struct value 赋值

	}
	return nil
}
