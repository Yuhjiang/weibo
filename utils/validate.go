package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

const (
	BadRequest = "请求参数错误"
)

var trans ut.Translator
var Validate *validator.Validate

func init() {
	zhCh := zh.New()
	uni := ut.New(zhCh, zhCh)
	trans, _ = uni.GetTranslator("zh")

	Validate = binding.Validator.Engine().(*validator.Validate)
	err := zhTranslations.RegisterDefaultTranslations(Validate, trans)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleValidateError(err validator.ValidationErrors) []string {
	errMsg := make([]string, len(err))
	for i, e := range err {
		errMsg[i] = e.Translate(trans)
	}
	return errMsg
}

// ValidateErrorResp 返回参数验证的错误响应
func ValidateErrorResp(c *gin.Context, err error) {
	errors := err.(validator.ValidationErrors)
	errMsg := HandleValidateError(errors)
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  BadRequest,
		"detail": errMsg,
	})
}
