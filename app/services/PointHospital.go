package services

import (
	"gin-data/app/models"
)

type PointHospital struct {
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

func GetPointHospitalTotal(maps interface{}) (count int,err error) {
	count,err = models.GetPointHospitalTotal(map[string]interface{}{})
	return count, err
}

func GetPointHospitalOne( query map[string]interface{},orderBy interface{}) (*PointHospital, error) {
	var nu *models.PointHospital
	nu,err := models.GetPointHospitalOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferPointHospitalModel(nu),nil
}

func GetPointHospitalPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (pointHospitals []*PointHospital, total int, err error) {
	total,err = models.GetPointHospitalTotal(query)
	if err != nil {
		return nil,0,err
	}
	us,err := models.GetPointHospitalPages(query,orderBy,pageNum,pageSize)
	pointHospitals = TransferPointHospitals(us)
	return pointHospitals,total,nil
}
func GetPointHospitalList( query map[string]interface{},orderBy interface{},limit int) ([]*PointHospital,[]error) {
	users, errors := models.GetPointHospitalList(query, orderBy, limit)
	pointHospitals := TransferPointHospitals(users)
	return pointHospitals,errors
}

func AddPointHospital( data map[string]interface{}) (string, error ){
	return models.AddPointHospital(data)
}

func ModifyPointHospital( id string,data map[string]interface{}) (err error) {
	err = models.ModifyPointHospitalById(id,data)
	return err
}

func DeletePointHospital(maps map[string]interface{}) (err error) {
	err = models.DeletePointHospitals(maps)
	return nil
}

func ClearAllPointHospital() (err error) {
	err = models.ClearAllPointHospital()
	return err
}

func TransferPointHospitalModel(u *models.PointHospital)(*PointHospital){
	pointHospital :=  &PointHospital{
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
	return pointHospital
}
func TransferPointHospitals(us []*models.PointHospital) ([]*PointHospital) {
	var pointHospitals []*PointHospital
	for _,value := range us {
		pointHospital := TransferPointHospitalModel(value)
		pointHospitals = append(pointHospitals, pointHospital)
	}
	return pointHospitals
}
