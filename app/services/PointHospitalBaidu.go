package services

import (
	"gin-data/app/models"
)

type PointHospitalBaidu struct {
	Id        string    `json:"id" form:"id" gorm:"id"`
	Name      string `json:"name" form:"name" gorm:"name"`
	Address   string `json:"address" form:"address" gorm:"address"`
	Adcode    string `json:"adcode" form:"adcode" gorm:"adcode"`
	Province  string `json:"province" form:"province" gorm:"province"`
	City      string `json:"city" form:"city" gorm:"city"`
	Area      string `json:"area" form:"area" gorm:"area"`
	Lat       string `json:"lat" form:"lat" gorm:"lat"`
	Lng       string `json:"lng" form:"lng" gorm:"lng"`
	Grade     string `json:"grade" form:"grade" gorm:"grade"`
	Natrue    string `json:"natrue" form:"nature" gorm:"natrue"`
	Type      string `json:"type" form:"type" gorm:"type"`
	Telephone string `json:"telephone" form:"telephone" gorm:"telephone"`
	Flag      int    `json:"flag" form:"flag" gorm:"flag"`
	State     int    `json:"state" form:"state" gorm:"state"`
	Status    int    `json:"status" form:"status" gorm:"status"`
}

func GetPointHospitalBaiduTotal(query map[string]interface{}) (count int,err error) {
	count,err = models.GetPointHospitalBaiduTotal(query)
	return count, err
}

func GetPointHospitalBaiduOne( query map[string]interface{},orderBy interface{}) (*PointHospitalBaidu, error) {
	var nu *models.PointHospitalBaidu
	nu,err := models.GetPointHospitalBaiduOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferPointHospitalBaiduModel(nu),nil
}

func GetPointHospitalBaiduPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (pointHospitalBaidus []*PointHospitalBaidu, total int, err error) {
	total,err = models.GetPointHospitalBaiduTotal(query)
	if err != nil {
		return nil,0,err
	}
	us,err := models.GetPointHospitalBaiduPages(query,orderBy,pageNum,pageSize)
	pointHospitalBaidus = TransferPointHospitalBaidus(us)
	return pointHospitalBaidus,total,nil
}
func GetPointHospitalBaiduList( query map[string]interface{},orderBy interface{},limit int) ([]*PointHospitalBaidu,[]error) {
	users, errors := models.GetPointHospitalBaiduList(query, orderBy, limit)
	pointHospitalBaidus := TransferPointHospitalBaidus(users)
	return pointHospitalBaidus,errors
}

func AddPointHospitalBaidu( data map[string]interface{}) (string, error ){
	return models.AddPointHospitalBaidu(data)
}

func ModifyPointHospitalBaidu( id string,data map[string]interface{}) (err error) {
	err = models.ModifyPointHospitalBaiduById(id,data)
	return err
}

func DeletePointHospitalBaidu(maps map[string]interface{}) (err error) {
	err = models.DeletePointHospitalBaidus(maps)
	return nil
}

func ClearAllPointHospitalBaidu() (err error) {
	err = models.ClearAllPointHospitalBaidu()
	return err
}

func TransferPointHospitalBaiduModel(u *models.PointHospitalBaidu)(*PointHospitalBaidu){
	pointHospitalBaidu :=  &PointHospitalBaidu{
		Id:u.Id,
		Name:u.Name,
		Address:u.Address,
		Adcode:u.Adcode,
		Province:u.Province,
		City:u.City,
		Area:u.Area,
		Lat:u.Lat,
		Lng:u.Lng,
		Grade:u.Grade,
		Natrue:u.Natrue,
		Type:u.Type,
		Telephone:u.Telephone,
		Flag:u.Flag,
		State:u.State,
		Status:u.Status,
	}
	return pointHospitalBaidu
}
func TransferPointHospitalBaidus(us []*models.PointHospitalBaidu) ([]*PointHospitalBaidu) {
	var pointHospitalBaidus []*PointHospitalBaidu
	for _,value := range us {
		pointHospitalBaidu := TransferPointHospitalBaiduModel(value)
		pointHospitalBaidus = append(pointHospitalBaidus, pointHospitalBaidu)
	}
	return pointHospitalBaidus
}
