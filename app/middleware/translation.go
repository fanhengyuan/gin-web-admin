package middleware

import (
	"gin-web-admin/common"
	"github.com/gin-gonic/gin"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

//设置Translation
func TranslationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := common.Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(common.Validate, trans)
			break
		case "en":
			en_translations.RegisterDefaultTranslations(common.Validate, trans)
			break
		case "zh_tw":
			zh_tw_translations.RegisterDefaultTranslations(common.Validate, trans)
			break
		default:
			zh_translations.RegisterDefaultTranslations(common.Validate, trans)
			break
		}

		//设置trans到context
		c.Set("trans", trans)
		c.Next()
	}
}
