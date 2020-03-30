package models

type UserDataRuleExt struct {
	Id    int64  `json:"-"`
	Uid   int64  `json:"uid"`
	OrgId string `json:"org_id"`
}

func (i *UserDataRuleExt) TableName() string {
	return "user_data_rule_ext"
}
