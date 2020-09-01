package service;

import (
	"github.com/ipipdotnet/ipdb-go"
	"log"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/xorm"
	//"log"
)

type CheckAreaService struct {

}

var db *ipdb.City;

func (s *CheckAreaService)CheckAreaByIP(
	ip,//IP
	area string,//地区编码
)(data []string,err error){
	//str := fmt.Sprintf("%d",time.Now().Unix());
	db, err := ipdb.NewCity("city.free.ipdb")
	if err != nil {
		log.Fatal(err.Error());
	}

	tmp, err := db.Find(ip, area)
	if err != nil {
		//fmt.Println(tmp)
		return tmp,err
		//for k, v := range m {
		//	fmt.Println(k, v)
		//}
	}
	////返回数据
	//fmt.Println("test3")
	return tmp,nil;
}