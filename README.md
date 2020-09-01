# area go 
根据用户访问IP,返回用户所在区域信息 【go语言版本】
*ip库信息识别借用 github.com/ipipdotnet/ipdb-go*


目录结构及文件说明
*本项目是由原生goland自主进行了简单的MVC结构封装,没有使用框架（因为要实现的功能比较单一所以没有必要使用框架增加程序运行的负载）*

```
****controller // 控制器目录
****route // 路由控制
****service // 业务逻辑目录
****util //常用工具函数封装库
****view //前端视图目录
*build.bat //windows打包发布脚本 双击运行即可，详情里面有注释可以进行修改*
*build.sh //linux打包发布脚本 sh build.sh 执行即可（停止脚本执行ps aux |grep area[脚本定义的进程名可以自主修改的]），详情里面有注释可以进行修改*
city.free.ipdb //IP库不要管除非要替换更好高级的版本，离线看可以在[IPIP](https://www.ipip.net/product/ip.html)购买 
main.go //项目主程序
```

TIPS

1.  *注意首次加载包时间超时或者不成功请修改 镜像：Go 1.13 及以上（推荐）-> go env -w GOPROXY=https://goproxy.cn,direct*

2.  *windows 首次运行报包错误就重新执行 go mod init mini_go(mini_go这个名称修改的话要把其他引入包的mini_go字段一起修改)*
3.  *启动项目后访问 127.0.0.1:8800/area/check (这里注意本地访问不要用localhost访问会拿不到本地ip的值，可能封装的ip获取函数不够完善，后续会继续优化)*
4.  *接口响应示例：*
```
{
  "code": 0,
  "msg": "获取成功",
  "data": {
    "areaKey": 1003,
    "ip": "xx.xx.xx.82",
    "local": "area:中国广东广州",
    "localArr": [
      "中国",
      "广东",
      "广州"
    ]
  }
}
```
5.  *执行到这里项目就允许成功了~*
6.  *接口响应体等的修改在util/resp.go*
7.  *接口返回的areaKey是banlist定义的key值，用户方便前端接收到区域信息后做对应的判断逻辑*




