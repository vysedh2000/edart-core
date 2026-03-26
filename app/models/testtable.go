package models

type TestTable struct{
	ID int  `gorm:"column:id"`
	Text string `gorm:"column:text"`
}

func (TestTable) TableName() string {
	return "testtable"
}