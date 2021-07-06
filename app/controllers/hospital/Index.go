package hospital

import (
	"errors"
	"fmt"
	"gin-data/app/services"
	"gin-data/utils"
	"github.com/gin-gonic/gin"
	lichv "github.com/lichv/go"
)

func GetHospitalPage(c *gin.Context) {
	var page,size int
	input := utils.GetMapFromContext(c)
	_,ok := input["page"]
	if ok {
		page = lichv.IntVal(input["page"])	
	}else{
		page = 1
	}
	_,ok = input["size"]
	if ok{
		size = lichv.IntVal(input["size"])
	}else{
		size = 50
	}
	var order = ""
	_, ok = input["order"]
	if ok {
		order = lichv.Strval(input["order"])
		delete(input, "order")
	} else {
		order = "id"
	}
	var sort = ""
	inputSort, ok := input["sort"]
	if ok {
		delete(input, "sort")
		sort = lichv.Strval(inputSort)
		if sort == "2" {
			sort = "desc"
		}else if sort == "1"{
			sort = "asc"
		}else{
			sort = ""
		}
	} else {
		sort = "asc"
	}

	data, total, _ := services.GetHospitalPages(input, order+" "+sort, page, size)
	var last = 1
	if total%size == 0 {
		last = total / size
	} else {
		last = total/size + 1
	}
	c.JSON(200,gin.H{
		"state":2000,
		"message":"success",
		"data":data,
		"page":  map[string]interface{}{"page": page, "size": size, "total": total, "last": last},
	})
}

func SaveHospital(c *gin.Context) {
	var err error
	input := utils.GetMapFromContext(c)

	name:=c.DefaultPostForm("name","")
	fmt.Println(name)
	id ,ok := input["id"]
	isExist := false
	if ok {
		if lichv.IntVal(id) > 0{
			isExist = true
		}
	}
	if isExist {
		err = services.ModifyHospital(lichv.IntVal(id), input)
		if err != nil {
			c.JSON(200,gin.H{
				"state":4001,
				"message":err.Error(),
			})
			return
		}
	}else{
		id, err = services.AddHospital(input)
		if err != nil{
			c.JSON(200,gin.H{
				"state":4002,
				"message":err.Error(),
			})
			return
		}
	}

	c.JSON(200,gin.H{
		"state":2000,
		"message":"success",
		"data":id,
	})
}

func DelHospital(c *gin.Context) {
	var err error
	input := utils.GetMapFromContext(c)

	id, ok := input["id"]
	isExist := false
	if ok {
		if lichv.IntVal(id) > 0 {
			isExist = true
		}
	}
	if isExist {
		err = services.DeleteHospital(map[string]interface{}{"id": lichv.IntVal(id)})
		if err != nil {
			c.JSON(200, gin.H{
				"state":   4001,
				"message": err.Error(),
			})
			return
		}
	} else {
		length := len(input)
		if length > 0{
			err = services.DeleteHospital(input)
			if err != nil {
				c.JSON(200, gin.H{
					"state":   4002,
					"message": "删除错误",
					"data":err.Error(),
				})
				return
			}
		}else{
				c.JSON(200, gin.H{
					"state":   4003,
					"message": "参数错误",
					"data":errors.New("参数错误"),
				})
				return
		}
	}

	c.JSON(200, gin.H{
		"state":   2000,
		"message": "success",
		"data":    id,
	})
}
