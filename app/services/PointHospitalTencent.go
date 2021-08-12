package services

import (
	"gin-data/app/models"
)

type PointHospitalTencent struct {
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

func GetPointHospitalTencentTotal(maps interface{}) (count int,err error) {
	count,err = models.GetPointHospitalTencentTotal(map[string]interface{}{})
	return count, err
}

func GetPointHospitalTencentOne( query map[string]interface{},orderBy interface{}) (*PointHospitalTencent, error) {
	var nu *models.PointHospitalTencent
	nu,err := models.GetPointHospitalTencentOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferPointHospitalTencentModel(nu),nil
}

func GetPointHospitalTencentPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (pointHospitalTencents []*PointHospitalTencent, total int, err error) {
	total,err = models.GetPointHospitalTencentTotal(query)
	if err != nil {
		return nil,0,err
	}
	us,err := models.GetPointHospitalTencentPages(query,orderBy,pageNum,pageSize)
	pointHospitalTencents = TransferPointHospitalTencents(us)
	return pointHospitalTencents,total,nil
}
func GetPointHospitalTencentList( query map[string]interface{},orderBy interface{},limit int) ([]*PointHospitalTencent,[]error) {
	users, errors := models.GetPointHospitalTencentList(query, orderBy, limit)
	pointHospitalTencents := TransferPointHospitalTencents(users)
	return pointHospitalTencents,errors
}

func AddPointHospitalTencent( data map[string]interface{}) (string, error ){
	return models.AddPointHospitalTencent(data)
}

func ModifyPointHospitalTencent( id string,data map[string]interface{}) (err error) {
	err = models.ModifyPointHospitalTencentById(id,data)
	return err
}

func DeletePointHospitalTencent(maps map[string]interface{}) (err error) {
	err = models.DeletePointHospitalTencents(maps)
	return nil
}

func ClearAllPointHospitalTencent() (err error) {
	err = models.ClearAllPointHospitalTencent()
	return err
}

func TransferPointHospitalTencentModel(u *models.PointHospitalTencent)(*PointHospitalTencent){
	pointHospitalTencent :=  &PointHospitalTencent{
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
	return pointHospitalTencent
}
func TransferPointHospitalTencents(us []*models.PointHospitalTencent) ([]*PointHospitalTencent) {
	var pointHospitalTencents []*PointHospitalTencent
	for _,value := range us {
		pointHospitalTencent := TransferPointHospitalTencentModel(value)
		pointHospitalTencents = append(pointHospitalTencents, pointHospitalTencent)
	}
	return pointHospitalTencents
}
