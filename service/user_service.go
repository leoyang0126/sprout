package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sprout/service/model"
	"strings"
	"sync"
)

const HOST = "http://127.0.0.1:8888/boffer"

type UserSaveReq struct {
	Name   string
	UserId string
}

func SaveUserInfo(ctx *gin.Context) {
	user := new(UserSaveReq)
	ctx.ShouldBindBodyWithJSON(user)
	fmt.Println("---user---", user)
	// 字典表接口
	dictUrl := HOST + "/hrcore/batchGetDictByCode"
	// 人员信息接口
	userUrl := HOST + "/hrcore/getUserBasicByKeyWord?keyword=${keyword}&userStatus=0"
	group := sync.WaitGroup{}
	group.Add(2)

	userChan := make(chan []model.UserDetailDto)
	go func() {
		defer group.Done()
		dictType := "ZZGZDD"
		post := "[\"" + dictType +
			"\"]"
		response, error := http.Post(dictUrl, "application/json", bytes.NewBufferString(post))
		if error != nil {
			fmt.Println("===调用地点接口没错误===")
		}
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("字典接口返回", string(body))
	}()
	go func() {
		defer group.Done()
		//utf8.DecodeRuneInString()
		userUrl := strings.ReplaceAll(userUrl, "${keyword}", url.QueryEscape(user.Name))
		response, _ := http.Get(userUrl)
		body, _ := ioutil.ReadAll(response.Body)
		bodyStr := string(body)
		fmt.Println("人员接口返回", bodyStr)
		data := gjson.Get(bodyStr, "data")
		var userList []model.UserDetailDto
		json.Unmarshal([]byte(data.String()), &userList)
		userChan <- userList
	}()
	log.Println("人员数据", <-userChan)
	group.Wait()
	QueryResumeBaseList()
	close(userChan)
}
