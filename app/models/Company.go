package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
	"time"
)

type Company struct {
	Id           int    `json:"id" form:"id" gorm:"id"`
	Name         string `json:"name" form:"name" gorm:"name"`
	Corporations string `json:"corporations" form:"corporations" gorm:"corporations"`
	Capital      string `json:"capital" form:"capital" gorm:"capital"`
	Capitaltype  string `json:"capitaltype" form:"capitaltype" gorm:"capitaltype"`
	Establish    string `json:"establish" form:"establish" gorm:"establish"`
	Checkflag    string `json:"checkflag" form:"checkflag" gorm:"checkflag"`
	Province     string `json:"province" form:"province" gorm:"province"`
	City         string `json:"city" form:"city" gorm:"city"`
	Area         string `json:"area" form:"area" gorm:"area"`
	Adcode       string `json:"adcode" form:"adcode" gorm:"adcode"`
	Street       string `json:"street" form:"street" gorm:"street"`
	Typename     string `json:"typename" form:"typename" gorm:"typename"`
	Creditcode   string `json:"creditcode" form:"creditcode" gorm:"creditcode"`
	Taxcode      string `json:"taxcode" form:"taxcode" gorm:"taxcode"`
	Registedcode string `json:"registedcode" form:"registedcode" gorm:"registedcode"`
	Orgcode      string `json:"orgcode" form:"orgcode" gorm:"orgcode"`
	Contact      string `json:"contact" form:"contact" gorm:"contact"`
	Others       string `json:"others" form:"others" gorm:"others"`
	Number       string `json:"number" form:"number" gorm:"number"`
	Industry     string `json:"industry" form:"industry" gorm:"industry"`
	Address      string `json:"address" form:"address" gorm:"address"`
	Urls         string `json:"urls" form:"urls" gorm:"urls"`
	Emails       string `json:"emails" form:"emails" gorm:"emails"`
	Tianyancha   string `json:"tianyancha" form:"tianyancha" gorm:"tianyancha"`
	Business     string `json:"business" form:"business" gorm:"business"`
	Content      string `json:"content" form:"content" gorm:"content"`
	Detail       string `json:"detail" form:"detail" gorm:"detail"`
	Addtime      int64  `json:"addtime" form:"addtime" gorm:"addtime"`
	Flag         int    `json:"flag" form:"flag" gorm:"flag"`
	State        int    `json:"state" form:"state" gorm:"state"`
	Status       int    `json:"status" form:"status" gorm:"status"`
}

var CompanyFields = []string{"id", "name", "corporations", "capital", "capitaltype", "establish", "checkflag", "province", "city", "area", "adcode", "street", "typename", "creditcode", "taxcode", "registedcode", "orgcode", "contact", "others", "number", "industry", "address", "urls", "emails", "tianyancha", "business", "content", "detail", "addtime", "flag", "state", "status"}

func FindCompanyById(id int) (*Company, error) {
	var company Company
	err := db.Model(&Company{}).Where("id = ? ", id).First(&company).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Company{}, err
	}
	return &company, nil
}

func GetCompanyOne(query map[string]interface{}, orderBy interface{}) (*Company, error) {
	var company Company
	model := db.Model(&Company{})
	for key, value := range query {
		if lichv.In(PointHospitalTencentFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	err := model.First(&company).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Company{}, nil
	}
	return &company, nil
}

func GetCompanyPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*Company, error) {
	var companys []*Company
	var errs []error
	var err error

	model := db.Model(&Company{})
	for key, value := range query {
		if lichv.In(PointHospitalTencentFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	beginNum := (pageNum - 1) * pageSize
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&companys)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return companys, err
}

func GetCompanyList(query map[string]interface{}, orderBy interface{}, limit int) ([]*Company, []error) {
	var Companys []*Company
	var errs []error
	model := db.Model(&Company{})
	for key, value := range query {
		if lichv.In(PointHospitalTencentFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	if limit > 0 {
		model = model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Companys).GetErrors()

	return Companys, errs
}

func GetCompanyTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&Company{})
	for key, value := range query {
		if lichv.In(PointHospitalTencentFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	err = model.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
func AddCompany(data map[string]interface{}) (int, error) {
	company := Company{}
	_, ok := data["id"]
	if ok {
		company.Id = lichv.IntVal(data["id"])
	}
	_, ok = data["name"]
	if ok {
		company.Name = lichv.Strval(data["name"])
	}
	_, ok = data["corporations"]
	if ok {
		company.Corporations = lichv.Strval(data["corporations"])
	}
	_, ok = data["capital"]
	if ok {
		company.Capital = lichv.Strval(data["capital"])
	}
	_, ok = data["capitaltype"]
	if ok {
		company.Capitaltype = lichv.Strval(data["capitaltype"])
	}
	_, ok = data["establish"]
	if ok {
		company.Establish = lichv.Strval(data["establish"])
	}
	_, ok = data["checkflag"]
	if ok {
		company.Checkflag = lichv.Strval(data["checkflag"])
	}
	_, ok = data["province"]
	if ok {
		company.Province = lichv.Strval(data["province"])
	}
	_, ok = data["city"]
	if ok {
		company.City = lichv.Strval(data["city"])
	}
	_, ok = data["area"]
	if ok {
		company.Area = lichv.Strval(data["area"])
	}
	_, ok = data["adcode"]
	if ok {
		company.Adcode = lichv.Strval(data["adcode"])
	}
	_, ok = data["street"]
	if ok {
		company.Street = lichv.Strval(data["street"])
	}
	_, ok = data["typename"]
	if ok {
		company.Typename = lichv.Strval(data["typename"])
	}
	_, ok = data["creditcode"]
	if ok {
		company.Creditcode = lichv.Strval(data["creditcode"])
	}
	_, ok = data["taxcode"]
	if ok {
		company.Taxcode = lichv.Strval(data["taxcode"])
	}
	_, ok = data["registedcode"]
	if ok {
		company.Registedcode = lichv.Strval(data["registedcode"])
	}
	_, ok = data["orgcode"]
	if ok {
		company.Orgcode = lichv.Strval(data["orgcode"])
	}
	_, ok = data["contact"]
	if ok {
		company.Contact = lichv.Strval(data["contact"])
	}
	_, ok = data["others"]
	if ok {
		company.Others = lichv.Strval(data["others"])
	}
	_, ok = data["number"]
	if ok {
		company.Number = lichv.Strval(data["number"])
	}
	_, ok = data["industry"]
	if ok {
		company.Industry = lichv.Strval(data["industry"])
	}
	_, ok = data["address"]
	if ok {
		company.Address = lichv.Strval(data["address"])
	}
	_, ok = data["urls"]
	if ok {
		company.Urls = lichv.Strval(data["urls"])
	}
	_, ok = data["emails"]
	if ok {
		company.Emails = lichv.Strval(data["emails"])
	}
	_, ok = data["tianyancha"]
	if ok {
		company.Tianyancha = lichv.Strval(data["tianyancha"])
	}
	_, ok = data["business"]
	if ok {
		company.Business = lichv.Strval(data["business"])
	}
	_, ok = data["content"]
	if ok {
		company.Content = lichv.Strval(data["content"])
	}
	_, ok = data["Detail"]
	if ok {
		company.Detail = lichv.Strval(data["detail"])
	}

	company.Addtime = time.Now().Unix()

	_, ok = data["flag"]
	if ok {
		company.Flag = lichv.IntVal(data["flag"])
	} else {
		company.Flag = 0
	}
	_, ok = data["state"]
	if ok {
		company.State = lichv.IntVal(data["state"])
	} else {
		company.State = 0
	}
	_, ok = data["status"]
	if ok {
		company.Status = lichv.IntVal(data["status"])
	} else {
		company.Status = 0
	}
	fmt.Println(company)
	if company.Name == "" {
		return 0, errors.New("参数为空")
	}

	if err := db.Create(&company).Error; err != nil {
		return 0, err
	}
	return company.Id, nil
}

func ModifyCompanyById(id int, data map[string]interface{}) error {
	if err := db.Model(&Company{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCompanyById(id int) error {
	if err := db.Where("id=?", id).Delete(Company{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCompanys(maps map[string]interface{}) error {
	model := db
	for key, value := range maps {
		if lichv.In(PointHospitalTencentFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	if err := model.Delete(&Company{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllCompany() error {
	if err := db.Unscoped().Delete(&Company{}).Error; err != nil {
		return err
	}
	return nil
}
