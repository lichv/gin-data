package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"regexp"
	"text/template"
)


func GetMapFromContext(context *gin.Context) (map[string]interface{}) {
	result := make(map[string]interface{})
	json := make(map[string]interface{})
	for i,_ := range context.Request.URL.Query(){
		temp := context.DefaultQuery(i,"")
		if temp != "" {
			result[i] = temp
		}
	}
	for j,_ := range context.Request.Form{
		temp := context.DefaultPostForm(j,"")
		if temp != ""{
			result[j] = temp
		}
	}
	for k,_ := range context.Request.PostForm{
		temp := context.DefaultPostForm(k,"")
		if temp!= ""{
			result[k] = temp
		}
	}
	_ = context.ShouldBindJSON(&json)
	for index,value:=range json{
		result[index] = value
	}

	return result
}

func ParseTemplateWithParams(template_content string,maps map[string]interface{}) (string,error) {
	var buff bytes.Buffer
	reg := regexp.MustCompile(`{{\s*(.*?)\s*}}`)
	submatchs := reg.FindAllStringSubmatch(template_content, -1)
	for _,submatch := range submatchs{
		tmp,_ := regexp.Compile(submatch[1])
		template_content = tmp.ReplaceAllString(template_content,"."+submatch[1])
	}
	tmpl, err := template.New("parseTemplateWithParams").Parse(template_content)
	if err != nil {
		return "",err
	}

	err = tmpl.Execute(&buff, maps)
	if err != nil {

		return "",err
	}
	result := buff.String()
	for _,submatch := range submatchs{
		tmp,_ := regexp.Compile("."+submatch[1])
		result = tmp.ReplaceAllString(result,submatch[1])
	}

	return result,nil
}