package main

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

type SignUpForm struct {
	Age        uint8  `form:"age" binding:"required,gte=1,lte=80"`
	Name       string `form:"name" binding:"required,min=3"`
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	RePassword string `form:"re_password" binding:"required,eqfield=Password"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	// 构建validator的中文翻译信息
	err := initTransValidator("zh")
	if err != nil {
		fmt.Println("initial trans error", err.Error())
	}

	route := gin.Default()


	route.GET("/bookable", getBookable)

	route.POST("/sign", sign)
	route.Run(":8085")
}

var translator ut.Translator

func initTransValidator(lan string) error {
	// 修改gin的validator引擎
	// 断言其为 gin 的 validator 类型
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
		zhT := zh.New()
		enT := en.New()
		// 第一个参数是备用语言环境 后面的是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		translator, ok = uni.GetTranslator(lan)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", lan)
		} else {
			fmt.Printf("lan is %s\n", lan)
		}

		switch lan {
		// 根据不同语言 进行翻译注册
		case "en":
			en_translations.RegisterDefaultTranslations(v, translator)
			fmt.Println("using en translator")
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, translator)
			fmt.Println("using zh translator")
		default:
			en_translations.RegisterDefaultTranslations(v, translator)
			fmt.Println("using default translator")
		}
		return nil
	}
	return nil
}



func sign(c *gin.Context) {
	var sign SignUpForm
	if err := c.ShouldBind(&sign); err != nil {
		// 这里启用自定义的报错 中文翻译
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"err": "internal error",
			})
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.Translate(translator),
		})
		return
	}
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
