package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"time"

	"github.com/ckam225/golang/fiber/docs"
)

func PrintJson(v any) {
	if s, e := json.MarshalIndent(v, "", "  "); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(string(s))
	}
}

func SetSwaggerInfos() {
	docs.SwaggerInfo.Title = os.Getenv("APP_NAME")
	docs.SwaggerInfo.Description = os.Getenv("APP_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("APP_VERSION")
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func HashBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func GenerateCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%v", RandInt(0, 9))
	}
	return code
}

func GenerateHashCode() string {
	return base64.StdEncoding.EncodeToString([]byte(GenerateCode()))
}
