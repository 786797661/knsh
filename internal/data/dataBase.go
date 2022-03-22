package data

//type RealworldDac interface {
//	LowadeWithPage(interface{},map[string]interface{})(interface{},error)
//	Find(interface{},map[string]interface{})(interface{},error)
//	Save(interface{})(interface{},error)
//}
import (
	"gorm.io/gorm"
)

type RealworldDac struct {
	gorm.Model
	result interface{}
	err    error
}

func (r *RealworldDac) LowadeWithPage(T interface{}, db *gorm.DB, dics ...map[string]interface{}) interface{} {
	var where = ""
	for index, dic := range dics {
		if index > 0 {
			where += "or (1=1"
		}
		for k, v := range dic {
			if where == "" {
				where = "where"
			} else {
				where += "and"
			}
			where += k + v.(string)
		}
		if index > 0 {
			where += ")"
		}
	}
	db.Where(where).Find(&T)
	return T
}
