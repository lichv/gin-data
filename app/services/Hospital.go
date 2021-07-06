package services

import (
	"gin-data/app/models"
	"time"
)

type Hospital struct {
	Id           int    `json:"id" form:"id" gorm:"id"`
	Code         string `json:"code" form:"code" gorm:"code"`
	Name         string `json:"name" form:"name" gorm:"name"`
	Alias        string `json:"alias" form:"alias" gorm:"alias"`
	Logo         string `json:"logo" form:"logo" gorm:"logo"`
	Grade        string `json:"grade" form:"grade" gorm:"grade"`
	Nature       string `json:"nature" form:"nature" gorm:"nature"`
	Type         string `json:"type" form:"type" gorm:"type"`
	Regyear      string `json:"regyear" form:"regyear" gorm:"regyear"`
	Size         string `json:"size" form:"size" gorm:"size"`
	Province     string `json:"province" form:"province" gorm:"province"`
	City         string `json:"city" form:"city" gorm:"city"`
	Town         string `json:"town" form:"town" gorm:"town"`
	Address      string `json:"address" form:"address" gorm:"address"`
	Location     string `json:"location" form:"location" gorm:"location"`
	Members      string `json:"members" form:"members" gorm:"members"`
	Beds         string `json:"beds" form:"beds" gorm:"beds"`
	Emergency    string `json:"emergency" form:"emergency" gorm:"emergency"`
	Inpatients   string `json:"inpatients" form:"inpatients" gorm:"inpatients"`
	Creditcode   string `json:"creditcode" form:"creditcode" gorm:"creditcode"`
	Corporations string `json:"corporations" form:"corporations" gorm:"corporations"`
	Typename     string `json:"typename" form:"typename" gorm:"typename"`
	Introduction string `json:"introduction" form:"introduction" gorm:"introduction"`
	Others       string `json:"others" form:"others" gorm:"others"`
	Source       string `json:"source" form:"source" gorm:"source"`
	Tianyancha   string `json:"tianyancha" form:"tianyancha" gorm:"tianyancha"`
	Yixue        string `json:"yixue" form:"yixue" gorm:"yixue"`
	Url          string `json:"url" form:"url" gorm:"url"`
	Phone        string `json:"phone" form:"phone" gorm:"phone"`
	Email        string `json:"email" form:"email" gorm:"email"`
	Addtime      string `json:"addtime" form:"addtime" gorm:"addtime"`
	Flag         int    `json:"flag" form:"flag" gorm:"flag"`
	State        int    `json:"state" form:"state" gorm:"state"`
	Status       int    `json:"status" form:"status" gorm:"status"`
}

func GetHospitalTotal(maps interface{}) (count int,err error) {
	count,err = models.GetHospitalTotal(map[string]interface{}{})
	return count, err
}

func GetHospitalOne( query map[string]interface{},orderBy interface{}) (*Hospital, error) {
	var nu *models.Hospital
	nu,err := models.GetHospitalOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferHospitalModel(nu),nil
}

func GetHospitalPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (hospitals []*Hospital, total int, err error) {
	total,err = models.GetHospitalTotal(query)
	if err != nil {
		return nil,0,err
	}
	us,err := models.GetHospitalPages(query,orderBy,pageNum,pageSize)
	hospitals = TransferHospitals(us)
	return hospitals,total,nil
}
func GetHospitalList( query map[string]interface{},orderBy interface{},limit int) ([]*Hospital,[]error) {
	users, errors := models.GetHospitalList(query, orderBy, limit)
	hospitals := TransferHospitals(users)
	return hospitals,errors
}

func AddHospital( data map[string]interface{}) (int, error ){
	return models.AddHospital(data)
}

func ModifyHospital( id int,data map[string]interface{}) (err error) {
	err = models.ModifyHospitalById(id,data)
	return err
}

func DeleteHospital(maps map[string]interface{}) (err error) {
	err = models.DeleteHospitals(maps)
	return nil
}

func ClearAllHospital() (err error) {
	err = models.ClearAllHospital()
	return err
}

func TransferHospitalModel(u *models.Hospital)(*Hospital){
	hospital :=  &Hospital{
		Id:u.Id,
		Code:u.Code,
		Name:u.Name,
		Alias:u.Alias,
		Logo:u.Logo,
		Grade:u.Grade,
		Nature:u.Nature,
		Type:u.Type,
		Regyear:u.Regyear,
		Size:u.Size,
		Province:u.Province,
		City:u.City,
		Town:u.Town,
		Address:u.Address,
		Location:u.Location,
		Members:u.Members,
		Beds:u.Beds,
		Emergency:u.Emergency,
		Inpatients:u.Inpatients,
		Creditcode:u.Creditcode,
		Corporations:u.Corporations,
		Typename:u.Typename,
		Introduction:u.Introduction,
		Others:u.Others,
		Source:u.Source,
		Tianyancha:u.Tianyancha,
		Yixue:u.Yixue,
		Url:u.Url,
		Phone:u.Phone,
		Email:u.Email,
		Addtime:time.Unix(u.Addtime,u.Addtime*100000).Format("2006-01-02 15:04:05"),
		Flag:u.Flag,
		State:u.State,
		Status:u.Status,
	}
	return hospital
}
func TransferHospitals(us []*models.Hospital) ([]*Hospital) {
	var hospitals []*Hospital
	for _,value := range us {
		hospital := TransferHospitalModel(value)
		hospitals = append(hospitals, hospital)
	}
	return hospitals
}
