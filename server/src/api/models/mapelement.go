package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
	"../../lib/common"
	"encoding/json"
)

type MapElement struct {
	Id           	int64  			`orm:"pk;auto;column(id)" json:"id"`
	Name        	string			`orm:"" json:"name"`
	Description 	string			`orm:"" json:"description"`
	Status		string			`orm:"" json:"status"`
	PrototypeType	string			`orm:"" json:"prototype_type"`
	PrototypeId	int64			`orm:"" json:"prototype_id"`
	Prototype	interface{}		`orm:"-" json:"prototype"`
	Weight 		int64			`orm:"" json:"weight"`
	Layer		*MapLayer		`orm:"rel(fk)" json:"layer"`
	Map		*Map			`orm:"rel(fk)" json:"map"`
	GraphSettings	string			`orm:"column(graph_settings)" json:"graph_settings"`
	Created_at   	time.Time		`orm:"auto_now_add;type(datetime);column(created_at)" json:"created_at"`
	Update_at    	time.Time		`orm:"auto_now;type(datetime);column(update_at)" json:"update_at"`
}

func (m *MapElement) TableName() string {
	return beego.AppConfig.String("db_map_elements")
}

func init() {
	orm.RegisterModel(new(MapElement))
}

// AddMapElement insert a new MapElement into database and returns
// last inserted Id on success.
func AddMapElement(m *MapElement) (id int64, err error) {

	switch m.PrototypeType {
	case "text":
		prototype := &MapText{}
		b, _ := json.Marshal(m.Prototype)
		json.Unmarshal(b, &prototype)
		if prototype != nil {
			if prototype.Id == 0 {
				if _, err = AddMapText(prototype); err != nil {
					return
				}

				m.PrototypeId = prototype.Id
				m.PrototypeType = "text"
			}
		}
	case "image":
		prototype := &MapImage{}
		b, _ := json.Marshal(m.Prototype)
		json.Unmarshal(b, &prototype)
		if prototype != nil {
			if prototype.Id == 0 {
				if _, err = AddMapImage(prototype); err != nil {
					return
				}

				m.PrototypeId = prototype.Id
				m.PrototypeType = "image"
			}
		}
	case "device":
	case "script":
	default:

	}

	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMapElementById retrieves MapElement by Id. Returns error if
// Id doesn't exist
func GetMapElementById(id int64) (v *MapElement, err error) {
	o := orm.NewOrm()
	v = &MapElement{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMapElement retrieves all MapElement matches certain condition. Returns empty list if
// no records exist
func GetAllMapElement(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, meta *map[string]int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MapElement))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MapElement
	qs = qs.OrderBy(sortFields...)
	objects_count, err := qs.Count()
	if err != nil {
		return
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		meta = &map[string]int64{
			"objects_count": objects_count,
			"limit":         limit,
			"offset":        offset,
		}
		return ml, meta, nil
	}
	return nil, nil, err
}

// UpdateMapElement updates MapElement by Id and returns error if
// the record to be updated doesn't exist
func UpdateMapElementById(m *MapElement) (err error) {

	// load old data
	//
	var oldMapElement *MapElement
	if oldMapElement, err = GetMapElementById(m.Id); err != nil {
		return
	}

	// delete old map image
	//
	if oldMapElement.PrototypeId == 0 {
		oldMapElement.PrototypeType = ""
	}

	switch oldMapElement.PrototypeType {
	case "text":
		DeleteMapText(oldMapElement.PrototypeId)
	case "image":
		DeleteMapImage(oldMapElement.PrototypeId)
	case "device":
	case "script":
	}

	//
	switch m.PrototypeType {
	case "text":
		prototype := &MapText{}
		b, _ := json.Marshal(m.Prototype)
		json.Unmarshal(b, &prototype)
		if prototype != nil {
			if prototype.Id == 0 {
				if _, err = AddMapText(prototype); err != nil {
					return
				}
			} else {
				if err = UpdateMapTextById(prototype); err != nil {
					return
				}
			}

			m.PrototypeId = prototype.Id
			m.PrototypeType = "text"
		} else {
			m.PrototypeId = 0
			m.PrototypeType = ""
		}
	case "image":

		prototype := &MapImage{}
		b, _ := json.Marshal(m.Prototype)
		json.Unmarshal(b, &prototype)
		if prototype != nil {
			if prototype.Id == 0 {
				if _, err = AddMapImage(prototype); err != nil {
					return
				}
			} else {
				if err = UpdateMapImageById(prototype); err != nil {
					return
				}
			}

			m.PrototypeId = prototype.Id
			m.PrototypeType = "image"
		} else {
			m.PrototypeId = 0
			m.PrototypeType = ""
		}
	case "device":
	case "script":

	}
	//-----------

	// ascertain id exists in the database
	o := orm.NewOrm()
	var num int64
	if num, err = o.Update(m); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return
}

// DeleteMapElement deletes MapElement by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMapElement(id int64) (err error) {

	var oldMapElement *MapElement
	if oldMapElement, err = GetMapElementById(int64(id)); err != nil {
		return
	}

	switch oldMapElement.PrototypeType {
	case "text":
		DeleteMapText(oldMapElement.PrototypeId)
	case "image":
		DeleteMapImage(oldMapElement.PrototypeId)
	case "device":
	case "script":
	}

	o := orm.NewOrm()
	v := MapElement{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MapElement{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func (m *MapElement) CompareWith(element *MapElement) bool {
	return reflect.DeepEqual(m, element)
}

func (m *MapElement) GetPrototype() (*MapElement, error) {

	switch m.PrototypeType {
	case "text":
		text, err := GetMapTextById(m.PrototypeId)
		if err != nil {
			return nil, err
		}

		m.Prototype = text
	case "image":
		image, err := GetMapImageById(m.PrototypeId)
		if err != nil {
			return nil, err
		}

		o := orm.NewOrm()
		if _, err := o.LoadRelated(image, "Image"); err != nil {
			return nil, err
		}

		image.Image.Url = common.GetLinkPath(image.Image.Image)
		m.Prototype = image
	case "device":
		device, err := GetMapDeviceById(m.PrototypeId)
		if err != nil {
			return nil, err
		}

		o := orm.NewOrm()
		_, err = o.LoadRelated(device, "Device")
		_, err = o.LoadRelated(device.Device, "States")
		_, err = o.LoadRelated(device, "States", 2)
		if err != nil {
			return nil, err
		}

		// update image url
		for _, state := range device.States {
			if state.Image != nil {
				state.Image.Url = common.GetLinkPath(state.Image.Image)
			}
		}

		m.Prototype = device

	case "script":
	default:

	}

	return m, nil
}