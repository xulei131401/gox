package validate

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator
var trans ut.Translator

func init() {
	uni = ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

// ValidStruct 校验结构体参数
func ValidStruct(s interface{}) error {
	// logx.Infof("go-zero解析后的参数:%+v", s)
	log.Printf("go-zero解析后的参数:%+v\n", s)

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	// 参数类型出错，比如nil,不是struct，是time.Time类型等
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	// 校验规则出错
	validationErrs := err.(validator.ValidationErrors)
	for _, err := range validationErrs {
		log.Println("Namespace:", err.Namespace())
		log.Println("Field:", err.Field())
		log.Println("StructNamespace:", err.StructNamespace())
		log.Println("StructField:", err.StructField())
		log.Println("Tag:", err.Tag())
		log.Println("ActualTag:", err.ActualTag())
		log.Println("Kind:", err.Kind())
		log.Println("Type:", err.Type())
		log.Println("Value:", err.Value())
		log.Println("Param:", err.Param())

		return errors.New(err.Translate(trans))
	}

	return nil
}
