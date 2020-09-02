package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chenhg5/collection"
	"github.com/go-redis/redis/v8"
	"log"
	"mini_go/service"
	"mini_go/util"
	"net/http"
	"time"
)

var (
	ctx = context.Background()
	areaService service.CheckAreaService
)

type areaData struct {
	Local string `json:"local"`
	LocalArr interface{} `json:"localArr"`
	AreaKey int `json:"areaKey"`
	Ip string `json:"ip"`
}

//区域检查
func CheckArea(writer http.ResponseWriter, request *http.Request) {
	// 获取用户真实ip
	ip := util.ClientPublicIP(request)
	if ip == ""{
		ip = util.ClientIP(request)
	}
		//实例化redis
		client := redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "123456", // no password set
			DB:       2,  // use default DB
		})
		_, redisErr := client.Ping(ctx).Result()
		if redisErr != nil {
			log.Fatal(redisErr.Error());
		}
		cacheData, cacheErr := client.HGet(ctx,"shield_area_ip",fmt.Sprintf("%s%s","ip_",ip)).Result()
		//判断是否有存在该ip的缓存值
		if cacheErr == redis.Nil {
			// ip库识别ip区域
			m, err := areaService.CheckAreaByIP(ip, "CN")
			if err == nil {
				tmp := areaData{
					Ip:ip,
					Local:fmt.Sprintf("area:%s%s%s",m[0],m[1],m[2]),
					LocalArr: m,
					AreaKey: 0,
				}
				//加载黑名单列表
				blackList := map[int]string{1001: "深圳", 1002: "唐山",1003: "广州"}
				for k, v := range blackList {
					flag := collection.Collect(m).Contains(v);
					if flag {
						fmt.Println(k, v)
						fmt.Println(fmt.Sprintf("属于屏蔽区域%s",v))
						tmp.AreaKey = k
						break;
					}
				}
				ret,err := json.Marshal(tmp)
				if err!=nil{
					log.Println(err.Error()) //打印日志
				}
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"ot cache",ret)
				client.HSet(ctx,"shield_area_ip",fmt.Sprintf("%s%s","ip_",ip),ret)
				util.RespOk(writer,tmp,"获取成功")
			}
		}else if cacheErr != nil { //异常
			panic(cacheErr)
		} else { //存在缓存直接转码返回
			byt := []byte(cacheData)
			tmp := make(map[string]interface{})
			jsonErr := json.Unmarshal(byt,&tmp)
			if jsonErr != nil {
				panic(jsonErr)
			}
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"cache json",tmp)
			util.RespOk(writer,tmp,"获取成功")
		}
}
