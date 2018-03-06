package utils

import (
	"crypto/tls"
	//"fmt"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

// 获取文件大小的接口
type Sizer interface {
	Size() int64
}

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

//var coder = base64.NewEncoding(base64Table)

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*********
reqString = Substr(reqString, 0, len(reqString)-1)
*********/
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {

		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

/*********
ltime : 秒数
dataformat : "2006/01/02 - 03:04:05"
**********/

func FormatTime(ltime int64, dataformat string) string {
	tm := time.Unix(ltime, 0)
	return tm.Format(dataformat)
}

func FormatTimeAll(ltime int64) string {
	tm := time.Unix(ltime, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func GetNetPath(uri string, path string) string {
	s := os.Args[0]
	fmt.Println("os.Args[0] : ", s)
	var index int = 0
	if runtime.GOOS == "windows" {
		index = strings.LastIndex(s, "\\")

	} else {
		index = strings.LastIndex(s, "/")
	}

	fmt.Println("index :", index)
	s = s[0 : index+1]

	if strings.LastIndex(uri, "/") != len(uri)-1 {
		uri += "/"
	}

	return uri + path[len(s):len(path)]

}

func ReadFile(path string) string {
	path = GetCurrPath() + path
	fmt.Println("path:  ", path)
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadFileFullPath(path string) string {

	fi, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func Base64Encode(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
	//return coder.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, bool) {
	str = strings.Replace(str, " ", "+", -1)
	data, err := b64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return string(data), true
}

func FormatImageBase64FormURL(data string) string {
	//head := "data:image/jpeg;base64,"
	//return Substr(data, len(head), len(data)-len(head))
	jpg_head := "data:image/jpeg;base64,"
	png_head := "data:image/png;base64,"
	imgData := ""

	dataArray := strings.Split(data, jpg_head)
	data = ""
	for _, value := range dataArray {
		data += value
	}
	dataArray = strings.Split(data, png_head)

	for _, value := range dataArray {
		imgData += value
	}

	//imgData = strings.Replace(imgData,"+"," ",-1)

	/*if strings.Index(data,jpg_head) == 0 {
		imgData = Substr(data, len(jpg_head), len(data)-len(jpg_head))
	} else if strings.Index(data,png_head) == 0 {
		imgData = Substr(data, len(png_head), len(data)-len(png_head))
	} else {
		imgData = data
	}*/
	return imgData
}

func WriteToFile(path string, data string) error {
	return ioutil.WriteFile(path, []byte(data), 0666) //写入文件(字节数组)
}

func Mkdir(path string) bool {
	err := os.MkdirAll(path, 0666)
	if err != nil {
		fmt.Printf("%s", err)
		return false
	} else {
		fmt.Print("Create Directory OK!")
	}
	return true
}

func ClientIp(addr string) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(addr)); err == nil {
		return ip
	}
	return ""
}

func SaveUploadFile(file multipart.File, filename string) {
	path := "./html/upload/" + filename
	nFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer nFile.Close()
	io.Copy(nFile, file)
}
func GetFileLenght(file multipart.File) (int64, bool) {
	if sizeInterface, ok := file.(Sizer); ok {
		if sizeInterface != nil {
			return sizeInterface.Size(), true
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
}

var config map[string]interface{} = make(map[string]interface{})

func ReadUtils() {
	data := ReadFile("utils.json")
	err := json.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Println("utils文件错误")
		fmt.Println(err)
		return
	}
}

func GetUtilsValue(key string) string {
	return config[key].(string)
}

func MD5(str string) string {

	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 sharejs.com
	return hex.EncodeToString(h.Sum(nil))

}

func SHA1(str string) string {

	h := sha1.New()
	h.Write([]byte(str)) // 需要加密的字符串为 sharejs.com
	return hex.EncodeToString(h.Sum(nil))

}

func GetValueInt(list map[string]interface{}, key string) (int, error) {
	if list[key] == nil {
		return 0, nil
	}
	switch list[key].(type) {
	case int:
		{
			return list[key].(int), nil
		}
	default:
		{
			return 0, errors.New("不是int类型")
		}
	}
	return 0, errors.New("未知类型")
}

func GetValueString(list map[string]interface{}, key string) (string, error) {
	if list[key] == nil {
		return "", nil
	}
	switch list[key].(type) {
	case string:
		{
			return list[key].(string), nil
		}
	default:
		{
			return "", errors.New("不是string类型")
		}
	}
	return "", errors.New("未知类型")
}

func GetValueFloat(list map[string]interface{}, key string) (float64, error) {
	if list[key] == nil {
		return 0, nil
	}
	switch list[key].(type) {
	case float64:
		{
			return list[key].(float64), nil
		}
	default:
		{
			return 0, errors.New("不是float类型")
		}
	}
	return 0, errors.New("未知类型")
}

/*
func GET(uri string) string {
	resp, err := http.Get(uri)

	if err != nil {
		panic(err)
		return "{}"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}*/

func GETAll(uri string, req *http.Request) (int, string) {

	var client = &http.Client{}
	//向服务端发送get请求

	request, _ := http.NewRequest("GET", uri, nil)
	request.Header = req.Header
	request.Header.Add("X-Forwarded-For", "cninct.com")
	request.Header.Add("X-Real-Ip", req.RemoteAddr)
	//request.Header.Add("REMOTE_ADDR",req.RemoteAddr)
	request.RemoteAddr = req.RemoteAddr

	for _, value := range req.Cookies() {
		request.AddCookie(value)
	}

	response, err := client.Do(request)

	if err != nil {

		fmt.Println(err)

		return 404, ""
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return response.StatusCode, string(body)
	//return 404,""
}

var POST_REMOTE_TIMEOUT time.Duration = 5

func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, time.Second*POST_REMOTE_TIMEOUT)
	if err != nil {
		return conn, err
	}

	tcp_conn := conn.(*net.TCPConn)
	tcp_conn.SetKeepAlive(false)

	return tcp_conn, err
}

func GETHTMLResult(uri string, req *http.Request) (int, string, bool, string) {

	var client = http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(network, addr, time.Second*POST_REMOTE_TIMEOUT)
				if err != nil {
					return conn, err
				}

				tcp_conn := conn.(*net.TCPConn)
				tcp_conn.SetKeepAlive(false)

				return tcp_conn, err
			},
		},
	}
	//向服务端发送get请求

	request, _ := http.NewRequest("GET", uri, nil)
	request.Header = req.Header
	request.Header.Add("X-Forwarded-For", "cninct.com")
	request.Header.Add("X-Real-Ip", req.RemoteAddr)
	//request.Header.Add("REMOTE_ADDR",req.RemoteAddr)
	request.RemoteAddr = req.RemoteAddr

	for _, value := range req.Cookies() {
		request.AddCookie(value)
	}

	response, err := client.Do(request)

	if err != nil {

		fmt.Println(err)

		return 404, "", false, ""
	}

	//request.Close()

	body, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close()
	//fmt.Println("response.Header:",response)
	//fmt.Println("type init : ",response.Header.Get("Content-Type"))
	if strings.Contains(response.Header.Get("Content-Type"), "text/plain") {
		return response.StatusCode, string(body), false, response.Header.Get("Content-Type") //判断返回的是否为文件，不是
	} else {
		return response.StatusCode, string(body), true, response.Header.Get("Content-Type") //是，且是HTML
	}

	//return 404,""
}

func GETS(uri string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//向服务端发送get请求
	response, err := client.Get(uri)
	//request, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {

		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(body)

	}
	response.Body.Close()
	return ""
}

func GET(uri string) string {
	var client = &http.Client{}
	//向服务端发送get请求

	request, _ := http.NewRequest("GET", uri, nil)

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	response.Header.Get("Content-Type")
	if response.StatusCode == 200 {

		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(body)

	}

	response.Body.Close()
	return ""
}

func GETForward(uri string) (string, string) {
	var client = &http.Client{}
	//向服务端发送get请求

	request, _ := http.NewRequest("GET", uri, nil)

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	type_ := response.Header.Get("Content-Type")
	if response.StatusCode == 200 {

		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(body), type_

	}

	response.Body.Close()
	return "", type_
}

func POSTAll(uri string, req1 *http.Request, data1 string) (int, string, http.Header) {
	fmt.Println("Content-Type : ", req1.Header["Content-Type"])

	var req *http.Request

	if len(req1.PostForm) > 0 {
		fmt.Println("req1.PostForm.Encode():", req1.PostForm.Encode())
		body := ioutil.NopCloser(strings.NewReader(req1.PostForm.Encode())) //把form数据编下码

		req, _ = http.NewRequest("POST", uri, body)
		if len(req1.Header["Content-Type"]) > 0 {
			req.Header.Set("Content-Type", req1.Header["Content-Type"][0])
		}
		//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	}
	fmt.Println("Content-Type BODY : ", data1)

	if len(req1.PostForm) <= 0 && req1.MultipartForm != nil && req1.MultipartForm.File != nil && len(req1.MultipartForm.File) > 0 {
		body := new(bytes.Buffer) // caveat IMO dont use this for large files, \
		// create a tmpfile and assemble your multipart from there (not tested)
		w := multipart.NewWriter(body)
		for key, value := range req1.MultipartForm.File {
			for _, value2 := range value {
				fw, err := w.CreateFormFile(key, value2.Filename)
				if err != nil {
					fmt.Println(err)
				}
				fd, err := value2.Open()
				if err != nil {
					fmt.Println(err)
				}
				defer fd.Close()

				_, err = io.Copy(fw, fd)
				if err != nil {
					fmt.Println(err)

				}
			}
		}
		w.Close()
		fmt.Println("POST File", body)
		req, _ = http.NewRequest("POST", uri, body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		list := strings.Split(data1, "\r\n")
		var contenttype string
		for _, value := range list {
			fmt.Println("value:", value)
			if strings.HasPrefix(value, "Content-Type: ") {
				l := strings.Split(value, ":")
				if len(l) == 2 {
					contenttype = l[1]
				}
				break
			}
		}
		req.Header.Add("Content-Type", contenttype)
	}

	//body := ioutil.NopCloser(strings.NewReader(data1)) //把form数据编下码

	//req, _ := http.NewRequest("POST", uri, body)

	if req == nil {

		req, _ = http.NewRequest("POST", uri, req1.Body)
		for _, value := range req1.Header["Content-Type"] {
			req.Header.Add("Content-Type", value)
		}
		//fmt.Println("2 req == nil")

	}

	var client = &http.Client{}
	//向服务端发送get请求

	req.Header.Add("X-Forwarded-For", "cninct.com")
	req.Header.Add("X-Real-Ip", req1.RemoteAddr)
	//request.Header.Add("REMOTE_ADDR",req.RemoteAddr)
	req.RemoteAddr = req1.RemoteAddr

	for _, value := range req1.Cookies() {
		req.AddCookie(value)
	}

	fmt.Println("req:", req)

	resp, err := client.Do(req)

	if err != nil {
		// handle error
		fmt.Println(err)
		return 404, "{}", nil
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	header := resp.Header
	resp.Body.Close()
	//fmt.Println("result : ",string(data))
	return resp.StatusCode, string(data), header

}

func POST(uri string, datas map[string]string) string {
	postValues := url.Values{}
	for key, value := range datas {
		postValues.Add(key, value)
	}
	body := ioutil.NopCloser(strings.NewReader(postValues.Encode())) //把form数据编下码
	var client = &http.Client{}
	//向服务端发送get请求

	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.AddCookie(&http.Cookie{Name: "_gat", Value: "1", Domain: ".sslforfree.com", Path: "/"})
	resp, err := client.Do(req)

	if err != nil {
		// handle error
		fmt.Println(err)
		return "{}"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	resp.Body.Close()
	return string(data)

}

func POSTS(uri string, datas map[string]string) string {
	postValues := url.Values{}
	for key, value := range datas {
		postValues.Add(key, value)
	}
	//body := ioutil.NopCloser(strings.NewReader(postValues.Encode())) //把form数据编下码
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	/*resp,err := client.Post(
	uri,
	"application/x-www-form-urlencoded",
	body)*/
	resp, err := client.PostForm(uri, postValues)

	if err != nil {
		// handle error
		fmt.Println(err)
		return "{}"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	resp.Body.Close()
	return string(data)

}

func PostData(uri string, datas string) string {

	body := ioutil.NopCloser(strings.NewReader(datas)) //把form数据编下码
	var client = &http.Client{}
	//向服务端发送get请求
	//fmt.Println(uri,"body",datas)
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		// handle error
		fmt.Println(err)
		return "{}"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		//panic(err)

		// handle error
	}
	resp.Body.Close()
	return string(data)

}

func PostDataForWSDL(uri string, datas string) string {

	body := ioutil.NopCloser(strings.NewReader(datas)) //把form数据编下码
	var client = &http.Client{}
	//向服务端发送get请求

	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)

	if err != nil {
		// handle error
		panic(err)
		return "{}"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		// handle error
	}
	resp.Body.Close()
	return string(data)

}

func UrlDecode(v string) (string, error) {
	l3, err3 := url.QueryUnescape(v)

	return l3, err3
}

func UrlEncode(v string) string {
	l3 := url.QueryEscape(v)

	return l3
}

//新建一个map，里面存放的数据是数组
func Map_A(name string, data []map[string]interface{}) map[string]interface{} {
	data_new := make(map[string]interface{})
	data_new[name] = data
	return data_new
}

//新建一个map，里面存放的数据是map
func Map_M(name string, data map[string]interface{}) map[string]interface{} {
	data_new := make(map[string]interface{})
	data_new[name] = data
	return data_new
}
func GetMax(first int64, args ...int64) int64 {
	for _, v := range args {
		if first < v {
			first = v
		}
	}
	return first
}

func GetPostData(r *http.Request) interface{} {
	requestbody, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	bf := bytes.NewBuffer(requestbody)
	r.Body = ioutil.NopCloser(bf)
	//text ,_:= utils.Base64Decode(string(requestbody))
	return string(requestbody)

}
