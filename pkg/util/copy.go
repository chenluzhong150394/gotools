package util

import (
	"reflect"
)

// UpCopyStruct 指定字段 结构体之间数据复制 (注意字段映射关系对应的数据类型需要是一样的)
func UpCopyStruct(f map[string]string, src, dst interface{}) error {
	//s, _ := json.Marshal(src)
	//fmt.Println(string(s))
	//fmt.Println("this is ", src.(*cdp.AccountContent).SaleDeptCode)
	for k, v := range f {
		l := reflect.ValueOf(src).Elem().FieldByName(k)
		r := reflect.ValueOf(dst).Elem().FieldByName(v)
		if l.IsValid() == false {
			//log.Printf("%s is not valid", k)
			continue
		}
		r.Set(l)
	}
	return nil
}

func GetStructAllField(src interface{}) map[string]string {
	obj := reflect.TypeOf(src)

	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}

	fields := make(map[string]string, 0)
	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		fields[field.Name] = field.Name
		//fmt.Println("结构体里的字段名",field.Name)
		//fmt.Println("结构体里的字段属性:",field.Type)
		//fmt.Println("结构体里面的字段的tag标签",field.Tag)
	}

	return fields
}

// UpCopyStructAll 指定字段 结构体之间数据复制 (注意字段映射关系对应的数据类型需要是一样的)
func UpCopyStructAll(f map[string]string, src, dst interface{}) (err error) {
	//s, _ := json.Marshal(src)
	//fmt.Println(string(s))
	//fmt.Println("this is ", src.(*cdp.AccountContent).SaleDeptCode)
	if len(f) == 0 {
		f = GetStructAllField(src)
	}

	for k, v := range f {
		l := reflect.ValueOf(src).Elem().FieldByName(k)
		r := reflect.ValueOf(dst).Elem().FieldByName(v)
		if l.IsValid() == false {
			//log.Printf("%s is not valid", k)
			continue
		}

		if r.IsValid() == false {
			//log.Println("右边的结构体中没有这个值,找不到进行跳过", v)
			continue
		}

		r.Set(l)
	}
	return nil
}
