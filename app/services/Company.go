package services

import (
	"gin-data/app/models"
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

func GetCompanyTotal(maps interface{}) (count int,err error) {
	count,err = models.GetCompanyTotal(map[string]interface{}{})
	return count, err
}

func GetCompanyOne( query map[string]interface{},orderBy interface{}) (*Company, error) {
	var nu *models.Company
	nu,err := models.GetCompanyOne(query,orderBy)
	if err != nil {
		return nil,err
	}
	return TransferCompanyModel(nu),nil
}

func GetCompanyPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) (companys []*Company, total int, err error) {
	total,err = models.GetCompanyTotal(query)
	if err != nil {
		return nil,0,err
	}
	us,err := models.GetCompanyPages(query,orderBy,pageNum,pageSize)
	companys = TransferCompanys(us)
	return companys,total,nil
}
func GetCompanyList( query map[string]interface{},orderBy interface{},limit int) ([]*Company,[]error) {
	users, errors := models.GetCompanyList(query, orderBy, limit)
	companys := TransferCompanys(users)
	return companys,errors
}

func AddCompany( data map[string]interface{}) (int, error ){
	return models.AddCompany(data)
}

func ModifyCompany( id int,data map[string]interface{}) (err error) {
	err = models.ModifyCompanyById(id,data)
	return err
}

func DeleteCompany(maps map[string]interface{}) (err error) {
	err = models.DeleteCompanys(maps)
	return nil
}

func ClearAllCompany() (err error) {
	err = models.ClearAllCompany()
	return err
}

func TransferCompanyModel(u *models.Company)(*Company){
	company :=  &Company{
		Id:u.Id,
		Name:u.Name,
		Corporations:u.Corporations,
		Capital:u.Capital,
		Capitaltype:u.Capitaltype,
		Establish:u.Establish,
		Checkflag:u.Checkflag,
		Province:u.Province,
		City:u.City,
		Area:u.Area,
		Adcode:u.Adcode,
		Street:u.Street,
		Typename:u.Typename,
		Creditcode:u.Creditcode,
		Taxcode:u.Taxcode,
		Registedcode:u.Registedcode,
		Orgcode:u.Orgcode,
		Contact:u.Contact,
		Others:u.Others,
		Number:u.Number,
		Industry:u.Industry,
		Address:u.Address,
		Urls:u.Urls,
		Emails:u.Emails,
		Tianyancha:u.Tianyancha,
		Business:u.Business,
		Content:u.Content,
		Detail:u.Detail,
		Addtime:time.Unix(u.Addtime,u.Addtime*100000).Unix(),
		Flag:u.Flag,
		State:u.State,
		Status:u.Status,
	}
	return company
}
func TransferCompanys(us []*models.Company) ([]*Company) {
	var companys []*Company
	for _,value := range us {
		company := TransferCompanyModel(value)
		companys = append(companys, company)
	}
	return companys
}
