package models

type UserDataRule struct {
	Id   int64  `json:"-"`
	Uid  int64  `json:"uid"`
	Orgs string `json:"orgs"`
	Org  string `json:"org"`
}

func (i *UserDataRule) TableName() string {
	return "user_data_rule"
}
