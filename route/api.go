package route

import (
	"mini_go/controllers"
	"net/http"
)

func ApiRoute(){
	//检查区域
	http.HandleFunc("/area/check", controllers.CheckArea);
}