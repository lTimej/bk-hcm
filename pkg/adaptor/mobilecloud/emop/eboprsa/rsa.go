package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"gitlab.ecloud.com/ecloud/ecloudsdkcore/consts"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/errs"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type Credentials struct {
	AppId      string
	Method     string
	Format     string
	Status     string
	PublicKey  string
	PrivateKey string
}

type CredentialsBuilder struct {
	Credentials *Credentials
}

func NewCredentialsBuilder(c *Credentials) *CredentialsBuilder {
	return &CredentialsBuilder{Credentials: c}
}

func (builder *CredentialsBuilder) WithAppId(appId string) *CredentialsBuilder {
	builder.Credentials.AppId = appId
	return builder
}

func (builder *CredentialsBuilder) WithMethod(method string) *CredentialsBuilder {
	builder.Credentials.Method = method
	return builder
}

func (builder *CredentialsBuilder) WithFormat(format string) *CredentialsBuilder {
	builder.Credentials.Format = format
	return builder
}

func (builder *CredentialsBuilder) WithStatus(status string) *CredentialsBuilder {
	builder.Credentials.Status = status
	return builder
}

func (builder *CredentialsBuilder) WithPublicKey(publicKey string) *CredentialsBuilder {
	builder.Credentials.PublicKey = publicKey
	return builder
}

func (builder *CredentialsBuilder) WithPrivateKey(privateKey string) *CredentialsBuilder {
	builder.Credentials.PrivateKey = privateKey
	return builder
}

func (builder *CredentialsBuilder) Build() *Credentials {
	if builder.Credentials.AppId == "" || builder.Credentials.Method == "" || builder.Credentials.PublicKey == "" || builder.Credentials.PrivateKey == "" {
		if builder.Credentials.AppId == "" {
			panic("AppId is required")
		}
		if builder.Credentials.Method == "" {
			panic("Method is required")
		}
		if builder.Credentials.PublicKey == "" {
			panic("PublicKey is required")
		}
		if builder.Credentials.PrivateKey == "" {
			panic("PrivateKey is required")
		}
	}
	if builder.Credentials.Format == "" {
		builder.Credentials.Format = "json"
	}
	if builder.Credentials.Status == "" {
		builder.Credentials.Status = "0"
	}
	return builder.Credentials
}

func (c *Credentials) Sign(publicMap map[string]string) (string, error) {
	// 1. 提取公共参数的键并排序
	keys := make([]string, 0, len(publicMap))
	for k := range publicMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 2. 组装参数字符串（排除 "sign" 和 "flowdId"）
	var paramList []string
	for _, key := range keys {
		if strings.EqualFold(key, "sign") || strings.EqualFold(key, "flowdId") {
			continue
		}
		value := fmt.Sprintf("%v", publicMap[key])
		paramList = append(paramList, fmt.Sprintf("%s=%s", key, value))
	}
	publicReqStr := strings.Join(paramList, "&")

	// 3. 解析私钥
	// block, _ := pem.Decode([]byte(privateKey))
	block, _ := base64.StdEncoding.DecodeString(c.PrivateKey)
	if block == nil {
		return "", errors.New("invalid private key")
	}
	privateKeyBytes, err := x509.ParsePKCS8PrivateKey(block)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}
	privateKeyRSA, ok := privateKeyBytes.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("private key is not of type RSA")
	}

	// 4. 使用 SHA-256 哈希生成签名
	hashed := sha256.Sum256([]byte(publicReqStr))
	signature, err := rsa.SignPKCS1v15(nil, privateKeyRSA, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}
	// 5. 返回 Base64 编码的签名
	return base64.StdEncoding.EncodeToString(signature), nil
}

func (c *Credentials) Verify(publicMap map[string]string) (bool, error) {
	// 1. 提取参数键并排序
	keys := make([]string, 0, len(publicMap))
	for k := range publicMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 2. 组装待验签字符串（排除 "sign" 和 "flowdId"）
	var paramList []string
	for _, key := range keys {
		if strings.EqualFold(key, "sign") || strings.EqualFold(key, "flowdId") {
			continue
		}
		value := fmt.Sprintf("%v", publicMap[key])
		paramList = append(paramList, fmt.Sprintf("%s=%s", key, value))
	}
	publicReqStr := strings.Join(paramList, "&")

	// 3. 获取签名
	signEncoded, ok := publicMap["sign"]
	if !ok {
		return false, errors.New("sign field is missing or invalid")
	}
	signEncoded, err := url.QueryUnescape(signEncoded)
	if err != nil {
		return false, fmt.Errorf("failed to decode sign: %v", err)
	}
	signature, err := base64.StdEncoding.DecodeString(signEncoded)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64 signature: %v", err)
	}

	// 4. 解析公钥
	// block, _ := pem.Decode([]byte(publicKey))
	block, _ := base64.StdEncoding.DecodeString(c.PublicKey)
	if block == nil {
		return false, errors.New("invalid public key format")
	}
	publicKeyParsed, err := x509.ParsePKIXPublicKey(block)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}
	publicKeyRSA, ok := publicKeyParsed.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("public key is not of type RSA")
	}

	// 5. 验签
	hashed := sha256.Sum256([]byte(publicReqStr))
	err = rsa.VerifyPKCS1v15(publicKeyRSA, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false, nil // 验签失败
	}
	return true, nil // 验签成功
}

func (builder *CredentialsBuilder) DecryptByPrivateKey(encrypted string, privateKey string) (string, error) {
	// 解码加密数据和私钥
	encryptedData, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted data: %v", err)
	}

	keyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %v", err)
	}

	// 解析私钥
	privateKeyObj, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	rsaPrivateKey, ok := privateKeyObj.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("invalid private key type")
	}

	// 分段解密
	inputLen := len(encryptedData)
	var buffer bytes.Buffer
	offset := 0

	for offset < inputLen {
		end := offset + MAX_DECRYPT_BLOCK
		if end > inputLen {
			end = inputLen
		}

		decryptedChunk, err := rsa.DecryptPKCS1v15(nil, rsaPrivateKey, encryptedData[offset:end])
		if err != nil {
			return "", fmt.Errorf("failed to decrypt chunk: %v", err)
		}

		buffer.Write(decryptedChunk)
		offset = end
	}

	return buffer.String(), nil
}

const (
	KEY_ALGORITHM     = "RSA"
	MAX_DECRYPT_BLOCK = 128 // 根据实际密钥长度调整，例如 1024 位密钥对应 128 字节
)

func (builder *CredentialsBuilder) DecryptByPublicKey(encrypted string, publicKey string) (string, error) {
	// 解码加密数据和公钥
	encryptedData, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted data: %v", err)
	}

	keyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key: %v", err)
	}

	// 解析公钥
	publicKeyObj, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPublicKey, ok := publicKeyObj.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("invalid public key type")
	}

	// 分段解密
	inputLen := len(encryptedData)
	var buffer bytes.Buffer
	offset := 0

	for offset < inputLen {
		end := offset + MAX_DECRYPT_BLOCK
		if end > inputLen {
			end = inputLen
		}

		decryptedChunk, err := rsa.DecryptPKCS1v15(nil, rsaPublicKey, encryptedData[offset:end])
		if err != nil {
			return "", fmt.Errorf("failed to decrypt chunk: %v", err)
		}

		buffer.Write(decryptedChunk)
		offset = end
	}

	return buffer.String(), nil
}

const (
	MAX_ENCRYPT_BLOCK = 10000 // 根据密钥长度调整，1024 位密钥对应 117 字节
)

func (builder *CredentialsBuilder) EncryptByPublicKey(source []byte, publicKey string) (string, error) {
	// 将源数据和公钥解码
	data := source
	keyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key: %v", err)
	}

	// 解析公钥
	publicKeyObj, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPublicKey, ok := publicKeyObj.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("invalid public key type")
	}

	// 分段加密
	// inputLen := len(data)
	// var buffer bytes.Buffer
	// offset := 0

	// for offset < inputLen {
	// 	end := offset + MAX_ENCRYPT_BLOCK
	// 	if end > inputLen {
	// 		end = inputLen
	// 	}
	// 	encryptedChunk, err := rsa.EncryptPKCS1v15(nil, rsaPublicKey, data[offset:end])
	// 	fmt.Println("哈哈哈哈")
	// 	if err != nil {
	// 		return "", fmt.Errorf("failed to encrypt chunk: %v", err)
	// 	}

	// 	buffer.Write(encryptedChunk)
	// 	offset = end
	// }

	// 将加密数据编码为 BASE64
	// encryptedData := buffer.Bytes()

	var encryptedData []byte
	for len(data) > 0 {
		var chunk []byte
		if len(data) > 64 {
			chunk, err = rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, data[:64])
			data = data[64:]
		} else {
			chunk, err = rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, data)
			data = nil
		}
		if err != nil {
			return "", errs.NewCredentialError("RSAPublicKeyEncrypt error", err)
		}
		encryptedData = append(encryptedData, chunk...)
	}

	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func (builder *CredentialsBuilder) EncryptByPrivateKey(source string, privateKey string) (string, error) {
	// 将源数据和私钥解码
	data := []byte(source)
	keyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %v", err)
	}

	// 解析私钥
	privateKeyObj, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	rsaPrivateKey, ok := privateKeyObj.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("invalid private key type")
	}

	// 分段加密
	inputLen := len(data)
	var buffer bytes.Buffer
	offset := 0

	for offset < inputLen {
		end := offset + MAX_ENCRYPT_BLOCK
		if end > inputLen {
			end = inputLen
		}

		encryptedChunk, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, 0, data[offset:end])
		if err != nil {
			return "", fmt.Errorf("failed to encrypt chunk: %v", err)
		}

		buffer.Write(encryptedChunk)
		offset = end
	}

	// 将加密数据编码为 BASE64
	encryptedData := buffer.Bytes()
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// func testEncryptByPrivateKey(source, privateKey string) {
// 	// 示例测试
// 	// source := "待加密的原始数据"
// 	// privateKey := "私钥的BASE64编码字符串"

// 	encrypted, err := EncryptByPrivateKey(source, privateKey)
// 	if err != nil {
// 		fmt.Println("加密失败:", err)
// 	} else {
// 		fmt.Println("加密成功:", encrypted)
// 	}
// }

// func testEncryptByPublicKey(source, publicKey string) {
// 	// 示例测试
// 	// source := "待加密的原始数据"
// 	// publicKey := "公钥的BASE64编码字符串"

// 	encrypted, err := EncryptByPublicKey(source, publicKey)
// 	if err != nil {
// 		fmt.Println("加密失败:", err)
// 	} else {
// 		fmt.Println("加密成功:", encrypted)
// 	}
// }

// func testDecryptByPublicKey(encrypted, publicKey string) {
// 	// 示例测试
// 	// encrypted := "密文的BASE64编码字符串"
// 	// publicKey := "公钥的BASE64编码字符串"

// 	decrypted, err := DecryptByPublicKey(encrypted, publicKey)
// 	if err != nil {
// 		fmt.Println("解密失败:", err)
// 	} else {
// 		fmt.Println("解密成功:", decrypted)
// 	}
// }

// func testDecryptByPrivateKey(encrypted, privateKey string) {
// 	// 示例测试
// 	// encrypted := "密文的BASE64编码字符串"
// 	// privateKey := "私钥的BASE64编码字符串"

// 	decrypted, err := DecryptByPrivateKey(encrypted, privateKey)
// 	if err != nil {
// 		fmt.Println("解密失败:", err)
// 	} else {
// 		fmt.Println("解密成功:", decrypted)
// 	}
// }

// RSAUtil 加密函数
func EncryptByPublicKey(msg string, publicKeyPEM string) (string, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("key type is not RSA")
	}

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(msg))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// doPostWithJson 发送POST请求
func DoPostWithJson(url, publicKey, encryptedReq string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(encryptedReq))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("user_id", "234324535235")
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func test1() {
	// 示例使用
	publicMap := map[string]string{
		"appId":  "50381602",
		"method": "SYAN_UNHT_createResellerUser",
		"format": "json",
		"status": "0",
	}
	publicKey := `MGcwDQYJKoZIhvcNAQEBBQADVgAwUwJMAItG02NDbB7FzQxEUS1yQckJovnVP3m5X+asOM2Tfv3HNd5kbQPnWEmrcPhOkIlAnV+ZWFmM+f5XBmGsH+YDSLuAyzJMTJMAw4D5awIDAQAB`
	privateKey := `MIIBhQIBADANBgkqhkiG9w0BAQEFAASCAW8wggFrAgEAAkwAi0bTY0NsHsXNDERRLXJByQmi+dU/eblf5qw4zZN+/cc13mRtA+dYSatw+E6QiUCdX5lYWYz5/lcGYawf5gNIu4DLMkxMkwDDgPlrAgMBAAECSzTZwntnaU7gFmgyQG+zbL1B9+NABZ9GNdsNvVxdPRJGFu32Q9vvVU82DoVDxZ2R5fBZXvmOph4EpdLvxiOVMtndYaSL/8W4PVYz6QImDWjaZKhJSh7+dcoKili0TdyNXrlyBwaM71SLGyLrRFGeifsM5rcCJgpi6A/TzD0kzr9iOu3rvN1d+jiw1TGRwKRNI2n/A9NioMtzOa7tAiYFGEWB4P6XjtcXYcBHeBRpUNbVmpfcW3zIodKIaOgCeRBHVH8+WQImAMQe3dv/epsWbONv+VCkE6f05u2ULB3WGchezlizDYp+1cLgBFkCJgY3ORoLB5TMckLlKGPl3+zw1GgiFEJ6ii4Yns6BvZwm3EflcZ7/`
	credential := Credentials{
		AppId:      "50381602",
		Method:     "SYAN_UNHT_createResellerUser",
		Format:     "json",
		Status:     "0",
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
	signature, err := credential.Sign(publicMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Signature:", signature)
	signature = url.QueryEscape(signature)
	publicMap["sign"] = signature
	fmt.Println(signature, "###")
	ok, err := credential.Verify(publicMap)
	if err != nil {
		fmt.Println(err, "?????")
	} else {
		fmt.Println(ok)
	}
	reqStr := map[string]string{
		"distributionChannel":     "CIDC-A-61b9a0d762cc425eb8a988009196489a",
		"distributionChannelName": "SZSL",
		"phone":                   "18752719177",
		"icType":                  "2",
		"icNo":                    "321320198710085632",
		"province":                "311",
		"city":                    "3160",
		"county":                  "316007",
		"address":                 "浙江省温州市苍南县",
	}
	data, _ := json.Marshal(reqStr)
	builder := NewCredentialsBuilder(&credential)
	r, err := builder.EncryptByPublicKey(data, publicKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-------------------------")
	str := ""
	for k, v := range publicMap {
		str += k
		str += "="
		str += v
		str += "&"
	}
	s := str[:len(str)-1]
	fmt.Println(s)
	d, err := DoPostWithJson("https://36.133.25.49:31015/api/query/emop?"+s, publicKey, r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
}

func Sign(params map[string]string) error {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var now time.Time
	if loc != nil {
		now = time.Now().In(loc)
	} else {
		now = time.Now()
	}
	params[consts.Timetamp] = now.Format(consts.TimestampFormat)
	params[consts.SignatureMethod] = consts.SignatureMethodValue
	params[consts.SignatureVersion] = consts.SignatureVersionValue
	params[consts.SignatureNonce] = utils.Nonce()
	keys := make([]string, len(params))
	index := 0
	for key := range params {
		keys[index] = key
		index++
	}
	sort.Strings(keys)
	builder := ""
	pos := 0
	paramsLen := len(keys)
	for _, key := range keys {
		value := params[key]
		builder += utils.PercentEncode(key)
		builder += consts.QuerySeparator
		builder += utils.PercentEncode(value)
		if pos != paramsLen-1 {
			builder += consts.ParameterSeparator
			pos++
		}
	}
	hashString := utils.ConvertToHexString(utils.Sha256Encode(builder))

	fmt.Println(hashString)
	return nil
}

func main() {
	// 示例使用
	// publicMap := map[string]string{
	// 	"appId":  "50381602",
	// 	"method": "SYAN_UNHT_createResellerUser",
	// 	"format": "json",
	// 	"status": "0",
	// }
	params := map[string]string{
		"Version":   "2016-12-05",
		"AccessKey": "49c0ea9b0f2045d9b61cd95e459ecd60",
	}
	// publicKey := `MGcwDQYJKoZIhvcNAQEBBQADVgAwUwJMAItG02NDbB7FzQxEUS1yQckJovnVP3m5X+asOM2Tfv3HNd5kbQPnWEmrcPhOkIlAnV+ZWFmM+f5XBmGsH+YDSLuAyzJMTJMAw4D5awIDAQAB`
	// privateKey := `MIIBhQIBADANBgkqhkiG9w0BAQEFAASCAW8wggFrAgEAAkwAi0bTY0NsHsXNDERRLXJByQmi+dU/eblf5qw4zZN+/cc13mRtA+dYSatw+E6QiUCdX5lYWYz5/lcGYawf5gNIu4DLMkxMkwDDgPlrAgMBAAECSzTZwntnaU7gFmgyQG+zbL1B9+NABZ9GNdsNvVxdPRJGFu32Q9vvVU82DoVDxZ2R5fBZXvmOph4EpdLvxiOVMtndYaSL/8W4PVYz6QImDWjaZKhJSh7+dcoKili0TdyNXrlyBwaM71SLGyLrRFGeifsM5rcCJgpi6A/TzD0kzr9iOu3rvN1d+jiw1TGRwKRNI2n/A9NioMtzOa7tAiYFGEWB4P6XjtcXYcBHeBRpUNbVmpfcW3zIodKIaOgCeRBHVH8+WQImAMQe3dv/epsWbONv+VCkE6f05u2ULB3WGchezlizDYp+1cLgBFkCJgY3ORoLB5TMckLlKGPl3+zw1GgiFEJ6ii4Yns6BvZwm3EflcZ7/`
	// credential := Credentials{
	// 	AppId:      "50381602",
	// 	Method:     "SYAN_UNHT_createResellerUser",
	// 	Format:     "json",
	// 	Status:     "0",
	// 	PublicKey:  publicKey,
	// 	PrivateKey: privateKey,
	// }
	// fmt.Println(publicMap)
	// fmt.Println(credential)
	Sign(params)
}
