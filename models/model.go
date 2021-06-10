package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const localDateTimeFormat string = "2006-01-02 15:04:05"

type LocalTime struct {
	time.Time
}

func (l LocalTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(l.Time).Format(localDateTimeFormat))
	// null 值返回空
	if stamp == `"0001-01-01 00:00:00"` {
		return []byte(`""`), nil
	}
	return []byte(stamp), nil
}

func (l LocalTime) UnmarshalJSON(data []byte) error {
	var err error
	l.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}
	return nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (l LocalTime) String() string {
	return l.Time.Format(localDateTimeFormat)
}

type LocalDate struct {
	time.Time
}

func (l LocalDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(l.Time).Format("2006-01-02"))
	// null 值返回空
	if stamp == `"0001-01-01 00:00:00"` {
		return []byte(`""`), nil
	}
	return []byte(stamp), nil
}

func (l LocalDate) UnmarshalJSON(data []byte) error {
	var err error
	l.Time, err = time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	return nil
}

func (t LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *LocalDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	Id        int64     `json:"id" gorm:"column:id"`                  //  id
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at"`  // 创建时间
	UpdatedAt LocalTime `json:"updated_at" gorm:"column:updated_at" ` // 修改时间
}

func ExecSql(sql string, values ...interface{}) error {
	d := DB.Master.Exec(sql, values...)
	return d.Error
}

//func (task *Task) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedAt", time.Now())
//	scope.SetColumn("ID", uuid.NewV4().String())
//	return nil
//}
//
//func (task *Task) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("UpdatedAt", time.Now())
//	return nil
//}
