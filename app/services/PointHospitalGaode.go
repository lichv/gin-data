package services

import (
	"gin-data/app/models"
)

type PointHospitalGaode struct {
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

func GetPointHospitalGaodeTotal(maps interface{}) (count int,err error) {
	count,err = models.GetPointHospitalGaodeTotal(map[string]interface{}{})
	return count, err
}

func GetPointHospitalGaodeOne( query map[string]interface{},orderBy interface{}) (*PointHospitalGaode, error) {
	var nu *models.PointHospitalGaode
	nu,err := models.GetPointHospitalGaodeOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferPointHospitalGaodeModel(nu),nil
}

func GetPointHospitalGaodePages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (pointHospitalGaodes []*PointHospitalGaode, total int, err error) {
	total,err = models.GetPointHospitalGaodeTotal(query)
	if err != nil {
		return nil,0,err
	}
	us,err := models.GetPointHospitalGaodePages(query,orderBy,pageNum,pageSize)
	pointHospitalGaodes = TransferPointHospitalGaodes(us)
	return pointHospitalGaodes,total,nil
}
func GetPointHospitalGaodeList( query map[string]interface{},orderBy interface{},limit int) ([]*PointHospitalGaode,[]error) {
	users, errors := models.GetPointHospitalGaodeList(query, orderBy, limit)
	pointHospitalGaodes := TransferPointHospitalGaodes(users)
	return pointHospitalGaodes,errors
}

func AddPointHospitalGaode( data map[string]interface{}) (string, error ){
	return models.AddPointHospitalGaode(data)
}

func ModifyPointHospitalGaode( id string,data map[string]interface{}) (err error) {
	err = models.ModifyPointHospitalGaodeById(id,data)
	return err
}

func DeletePointHospitalGaode(maps map[string]interface{}) (err error) {
	err = models.DeletePointHospitalGaodes(maps)
	return nil
}

func ClearAllPointHospitalGaode() (err error) {
	err = models.ClearAllPointHospitalGaode()
	return err
}

func TransferPointHospitalGaodeModel(u *models.PointHospitalGaode)(*PointHospitalGaode){
	pointHospitalGaode :=  &PointHospitalGaode{
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
	return pointHospitalGaode
}
func TransferPointHospitalGaodes(us []*models.PointHospitalGaode) ([]*PointHospitalGaode) {
	var pointHospitalGaodes []*PointHospitalGaode
	for _,value := range us {
		pointHospitalGaode := TransferPointHospitalGaodeModel(value)
		pointHospitalGaodes = append(pointHospitalGaodes, pointHospitalGaode)
	}
	return pointHospitalGaodes
}
