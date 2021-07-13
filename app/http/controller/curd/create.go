/**
* @Author: lik
* @Date: 2021/3/9 11:25
* @Version 1.0
 */
package curd

import (
	"gitee.com/open-product/dtcloud-api/app/service/curd"
	"gitee.com/open-product/dtcloud-api/routers/middleware/my_jwt"
	"regexp"
	"strings"
)

type CreateData struct {
}

func (c *CreateData) PublicCreate(dat map[string]interface{}, claims *my_jwt.CustomClaims) bool {

	v := make(map[string]interface{})
	mtm := make(map[string]interface{})

	for k := range dat {

		if regexp.MustCompile(`form-field-boolean-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-char-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-integer-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-many2one-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-float-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-date-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-datetime-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-html-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-text-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-radio-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-select-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		if regexp.MustCompile(`form-field-binary-`).FindAllStringSubmatch(k, -1) != nil {
			field := strings.Split(k, "-")[3]
			v[field] = dat[k]
		}
		// TODO
		if regexp.MustCompile(`form-field-m2m`).FindAllStringSubmatch(k, -1) != nil {
			if regexp.MustCompile(`form-field-m2m0\?`).FindAllStringSubmatch(k, -1) != nil {
				field := strings.Split(k, "?")[1]

				mtm["ids"] = dat[k]
				mtm["model"] = dat["model"]
				mtm["m2m_fields"] = field
				mtm["many2many"] = 0

				//for i := l.Front(); i != nil; i = i.Next() {
				//	fmt.Println(i.Value)
				//}

			} else if regexp.MustCompile(`form-field-m2m1\?`).FindAllStringSubmatch(k, -1) != nil {
				field := strings.Split(k, "?")[1]

				mtm["ids"] = dat[k]
				mtm["model"] = dat["model"]
				mtm["m2m_fields"] = field
				mtm["many2many"] = 1

			} else if regexp.MustCompile(`form-field-m2m2\?`).FindAllStringSubmatch(k, -1) != nil {
				field := strings.Split(k, "?")[1]

				mtm["ids"] = dat[k]
				mtm["model"] = dat["model"]
				mtm["m2m_fields"] = field
				mtm["many2many"] = 2

			} else if regexp.MustCompile(`form-field-m2m3\?`).FindAllStringSubmatch(k, -1) != nil {
				field := strings.Split(k, "?")[1]

				mtm["ids"] = dat[k]
				mtm["model"] = dat["model"]
				mtm["m2m_fields"] = field
				mtm["many2many"] = 3

			} else if regexp.MustCompile(`form-field-m2m4\?`).FindAllStringSubmatch(k, -1) != nil {
				field := strings.Split(k, "?")[1]

				mtm["ids"] = dat[k]
				mtm["model"] = dat["model"]
				mtm["m2m_fields"] = field
				mtm["many2many"] = 4

			} else if regexp.MustCompile(`form-field-m2m5\?`).FindAllStringSubmatch(k, -1) != nil {
				field := strings.Split(k, "?")[1]

				mtm["ids"] = dat[k]
				mtm["model"] = dat["model"]
				mtm["m2m_fields"] = field
				mtm["many2many"] = 5

			} else if regexp.MustCompile(`form-field-m2m6\?`).FindAllStringSubmatch(k, -1) != nil {
				field := strings.Split(k, "?")[1]
				mtm["ids"] = dat[k]
				mtm["model"] = dat["model"]
				mtm["m2m_fields"] = field
				mtm["many2many"] = 6

			}
		}
		v["form-field-m2m"] = mtm

	}

	return curd.CreateUserCurdFactory().PublicCreate(v, dat, claims)

}
