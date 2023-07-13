package header

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//获取上传的文件的文件流
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "上传文件失败:", err)
		log.Println("获取不到文件夹:", err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println("关闭上传的文件时失败:", err)
		}
	}(file)
	if header == nil {
		return
	}

	//读取文件，将文件内容转化为字节流
	data, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "读取上传文件时失败:", err)
		log.Println("读取上传的文件时失败:", err)
	}
	//将字节流写入到文件中
	err = os.WriteFile("./upload/"+header.Filename, data, os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "写入上传的文件时失败:", err)
		log.Println("写入上传的文件时失败:", err)
		return
	}
	//返回上传成功的信息
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//定义返回的json数据
	resp := map[string]string{
		"code":    "0",
		"message": "上传成功",
		"data":    header.Filename,
	}
	//转化为json数据
	jsonresp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "转化json数据时失败:", err)
		log.Println("转化json数据时失败:", err)
		return
	}
	//写入返回数据
	_, err = w.Write(jsonresp)
	if err != nil {
		return
	}
	log.Println("上传成功")
	return
}
func Download(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//通过 url(/v1/book/download/{name}) 中的{name},去获取到文件名
	filename, ok := pathParams["name"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "获取文件名失败")
		return
	}
	//读取文件内容
	data, err := os.ReadFile("./download/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "文件不存在:", err)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.WriteHeader(http.StatusOK)
	//将文件内容写入到响应中
	_, err = w.Write(data)
	if err != nil {
		return
	}
	log.Println("下载成功")
	return
}
