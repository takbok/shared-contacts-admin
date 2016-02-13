// Copyright 2016 Takbok, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE_2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package demo

import (
	"fmt"
)

const (
	NUM_COLUMNS int = 129
	ACTION_COL_IDX int = 0
	ID_COL_IDX int = 1
	NAME_COL_IDX int = 2
	COMPANY_COL_IDX int = 3
	JOBTITLE_COL_IDX int = 4
	NOTES_COL_IDX int = 5
	CATEGORIES_COL_IDX int = 6
	CONTACTTYPE_COL_IDX int = 7
	DOMAIN_COL_IDX int = 8
	APIID_COL_IDX int = 9
	WORKRELTYPE_COL_IDX int = 10
	BIRTHDAY_COL_IDX int = 11
	HOMERELTYPE_COL_IDX int = 12
	OTHERRELTYPE_COL_IDX int = 13
	MOBILERELTYPE_COL_IDX int = 14
	NAMELOWER_COL_IDX int = 15
	NAMELOWERUNIQUE_COL_IDX int = 16
	NICKNAME_COL_IDX int = 17
	EXTERNALID_COL_IDX int = 18
	OCCUPATION_COL_IDX int = 19
	DEPARTMENT_COL_IDX int = 20
	JOBDESCRIPTION_COL_IDX int = 21
	IM_COL_IDX int = 22
	IMREL_COL_IDX int = 23
	WEBSITE_COL_IDX int = 24
	WEBSITEREL_COL_IDX int = 25
	RELATION_COL_IDX int = 26
	RELATIONREL_COL_IDX int = 27
	CUSTOMKEY1_COL_IDX int = 28
	CUSTOMVALUE1_COL_IDX int = 29
	CUSTOMKEY2_COL_IDX int = 30
	CUSTOMVALUE2_COL_IDX int = 31
	CUSTOMKEY3_COL_IDX int = 32
	CUSTOMVALUE3_COL_IDX int = 33
	CUSTOMKEY4_COL_IDX int = 34
	CUSTOMVALUE4_COL_IDX int = 35
	CUSTOMKEY5_COL_IDX int = 36
	CUSTOMVALUE5_COL_IDX int = 37
	CUSTOMKEY6_COL_IDX int = 38
	CUSTOMVALUE6_COL_IDX int = 39
	BUSINESSFAX_COL_IDX int = 40
	BUSINESSPHONE_COL_IDX int = 41
	BUSINESSPHONE2_COL_IDX int = 42
	HOMEFAX_COL_IDX int = 43
	HOMEPHONE_COL_IDX int = 44
	HOMEPHONE2_COL_IDX int = 45
	OTHERPHONE_COL_IDX int = 46
	MOBILEPHONE_COL_IDX int = 47
	PAGER_COL_IDX int = 48
	HOMEADDRESS_COL_IDX int = 49
	BUSINESSADDRESS_COL_IDX int = 50
	OTHERADDRESS_COL_IDX int = 51
	E_MAILADDRESS_COL_IDX int = 52
	E_MAIL2ADDRESS_COL_IDX int = 53
	E_MAIL3ADDRESS_COL_IDX int = 54
	E_MAIL4ADDRESS_COL_IDX int = 55
	E_MAIL5ADDRESS_COL_IDX int = 56
	E_MAIL6ADDRESS_COL_IDX int = 57
	E_MAIL7ADDRESS_COL_IDX int = 58
	E_MAIL8ADDRESS_COL_IDX int = 59
	E_MAIL9ADDRESS_COL_IDX int = 60
	E_MAIL10ADDRESS_COL_IDX int = 61
	E_MAIL11ADDRESS_COL_IDX int = 62
	E_MAIL12ADDRESS_COL_IDX int = 63
	E_MAIL13ADDRESS_COL_IDX int = 64
	E_MAIL14ADDRESS_COL_IDX int = 65
	E_MAIL15ADDRESS_COL_IDX int = 66
	E_MAIL16ADDRESS_COL_IDX int = 67
	E_MAIL17ADDRESS_COL_IDX int = 68
	E_MAIL18ADDRESS_COL_IDX int = 69
	E_MAIL19ADDRESS_COL_IDX int = 70
	E_MAIL20ADDRESS_COL_IDX int = 71
	E_MAIL21ADDRESS_COL_IDX int = 72
	E_MAIL22ADDRESS_COL_IDX int = 73
	E_MAIL23ADDRESS_COL_IDX int = 74
	E_MAIL24ADDRESS_COL_IDX int = 75
	E_MAIL25ADDRESS_COL_IDX int = 76
	E_MAIL26ADDRESS_COL_IDX int = 77
	E_MAIL27ADDRESS_COL_IDX int = 78
	E_MAIL28ADDRESS_COL_IDX int = 79
	E_MAIL29ADDRESS_COL_IDX int = 80
	E_MAIL30ADDRESS_COL_IDX int = 81
	E_MAIL31ADDRESS_COL_IDX int = 82
	E_MAIL32ADDRESS_COL_IDX int = 83
	E_MAIL33ADDRESS_COL_IDX int = 84
	E_MAIL34ADDRESS_COL_IDX int = 85
	E_MAIL35ADDRESS_COL_IDX int = 86
	E_MAIL36ADDRESS_COL_IDX int = 87
	E_MAIL37ADDRESS_COL_IDX int = 88
	E_MAIL38ADDRESS_COL_IDX int = 89
	E_MAIL39ADDRESS_COL_IDX int = 90
	E_MAIL40ADDRESS_COL_IDX int = 91
	E_MAIL41ADDRESS_COL_IDX int = 92
	E_MAIL42ADDRESS_COL_IDX int = 93
	E_MAIL43ADDRESS_COL_IDX int = 94
	E_MAIL44ADDRESS_COL_IDX int = 95
	E_MAIL45ADDRESS_COL_IDX int = 96
	E_MAIL46ADDRESS_COL_IDX int = 97
	E_MAIL47ADDRESS_COL_IDX int = 98
	E_MAIL48ADDRESS_COL_IDX int = 99
	E_MAIL49ADDRESS_COL_IDX int = 100
	E_MAIL50ADDRESS_COL_IDX int = 101
	E_MAIL51ADDRESS_COL_IDX int = 102
	E_MAIL52ADDRESS_COL_IDX int = 103
	E_MAIL53ADDRESS_COL_IDX int = 104
	E_MAIL54ADDRESS_COL_IDX int = 105
	E_MAIL55ADDRESS_COL_IDX int = 106
	E_MAIL56ADDRESS_COL_IDX int = 107
	E_MAIL57ADDRESS_COL_IDX int = 108
	E_MAIL58ADDRESS_COL_IDX int = 109
	E_MAIL59ADDRESS_COL_IDX int = 110
	E_MAIL60ADDRESS_COL_IDX int = 111
	E_MAIL61ADDRESS_COL_IDX int = 112
	E_MAIL62ADDRESS_COL_IDX int = 113
	E_MAIL63ADDRESS_COL_IDX int = 114
	E_MAIL64ADDRESS_COL_IDX int = 115
	E_MAIL65ADDRESS_COL_IDX int = 116
	E_MAIL66ADDRESS_COL_IDX int = 117
	E_MAIL67ADDRESS_COL_IDX int = 118
	E_MAIL68ADDRESS_COL_IDX int = 119
	E_MAIL69ADDRESS_COL_IDX int = 120
	WEBSITEHOME_PAGE_COL_IDX int = 121
	WEBPAGE_COL_IDX int = 122
	WEBSITEBLOG_COL_IDX int = 123
	WEBSITEPROFILE_COL_IDX int = 124
	WEBSITEHOME_COL_IDX int = 125
	WEBSITEWORK_COL_IDX int = 126
	WEBSITEOTHER_COL_IDX int = 127
	WEBSITEFTP_COL_IDX int = 128
)

var column_name_map map[string]int
var column_names []string

func buildColumnMap() {
	column_names = []string{}
	column_name_map = make(map[string]int)

	column_names = append(column_names, "Action")
	column_names = append(column_names, "ID", "Name", "Company", "Job Title", "Notes", "Categories", "contactType",
							"domain", "apiId", "WorkRelType", "birthday", "HomeRelType", "OtherRelType",
							"MobileRelType", "NameLower", "NameLowerUnique", "NickName", "ExternalId",
							"Occupation", "Department", "Job Description", "IM", "IM Rel", "Website",
							"Website Rel", "Relation", "Relation Rel")
	for n := 0; n < 6; n++ {
		column_names = append(column_names, fmt.Sprintf("Custom Key%v", n+1))
		column_names = append(column_names, fmt.Sprintf("Custom Value%v", n+1))
	}
	column_names = append(column_names, "Business Fax", "Business Phone", "Business Phone 2", "Home Fax", "Home Phone",
						   "Home Phone 2", "Other Phone", "Mobile Phone", "Pager", "Home Address", "Business Address",
							"Other Address", "E-mail Address")
	for n := 2; n <= 69; n++ {
		column_names = append(column_names, fmt.Sprintf("E-mail %v Address", n))
	}
	column_names = append(column_names, "Website Home-Page", "Web Page", "Website Blog", "Website Profile", "Website Home",
							"Website Work", "Website Other", "Website FTP")
							
	column_name_map["Action"] = ACTION_COL_IDX
	column_name_map["ID"] = ID_COL_IDX
	column_name_map["Name"] = NAME_COL_IDX
	column_name_map["Company"] = COMPANY_COL_IDX
	column_name_map["Job Title"] = JOBTITLE_COL_IDX
	column_name_map["Notes"] = NOTES_COL_IDX
	column_name_map["Categories"] = CATEGORIES_COL_IDX
	column_name_map["contactType"] = CONTACTTYPE_COL_IDX
	column_name_map["domain"] = DOMAIN_COL_IDX
	column_name_map["apiId"] = APIID_COL_IDX
	column_name_map["WorkRelType"] = WORKRELTYPE_COL_IDX
	column_name_map["birthday"] = BIRTHDAY_COL_IDX
	column_name_map["HomeRelType"] = HOMERELTYPE_COL_IDX
	column_name_map["OtherRelType"] = OTHERRELTYPE_COL_IDX
	column_name_map["MobileRelType"] = MOBILERELTYPE_COL_IDX
	column_name_map["NameLower"] = NAMELOWER_COL_IDX
	column_name_map["NameLowerUnique"] = NAMELOWERUNIQUE_COL_IDX
	column_name_map["NickName"] = NICKNAME_COL_IDX
	column_name_map["ExternalId"] = EXTERNALID_COL_IDX
	column_name_map["Occupation"] = OCCUPATION_COL_IDX
	column_name_map["Department"] = DEPARTMENT_COL_IDX
	column_name_map["Job Description"] = JOBDESCRIPTION_COL_IDX
	column_name_map["IM"] = IM_COL_IDX
	column_name_map["IM Rel"] = IMREL_COL_IDX
	column_name_map["Website"] = WEBSITE_COL_IDX
	column_name_map["Website Rel"] = WEBSITEREL_COL_IDX
	column_name_map["Relation"] = RELATION_COL_IDX
	column_name_map["Relation Rel"] = RELATIONREL_COL_IDX
	column_name_map["Custom Key1"] = CUSTOMKEY1_COL_IDX
	column_name_map["Custom Value1"] = CUSTOMVALUE1_COL_IDX
	column_name_map["Custom Key2"] = CUSTOMKEY2_COL_IDX
	column_name_map["Custom Value2"] = CUSTOMVALUE2_COL_IDX
	column_name_map["Custom Key3"] = CUSTOMKEY3_COL_IDX
	column_name_map["Custom Value3"] = CUSTOMVALUE3_COL_IDX
	column_name_map["Custom Key4"] = CUSTOMKEY4_COL_IDX
	column_name_map["Custom Value4"] = CUSTOMVALUE4_COL_IDX
	column_name_map["Custom Key5"] = CUSTOMKEY5_COL_IDX
	column_name_map["Custom Value5"] = CUSTOMVALUE5_COL_IDX
	column_name_map["Custom Key6"] = CUSTOMKEY6_COL_IDX
	column_name_map["Custom Value6"] = CUSTOMVALUE6_COL_IDX
	column_name_map["Business Fax"] = BUSINESSFAX_COL_IDX
	column_name_map["Business Phone"] = BUSINESSPHONE_COL_IDX
	column_name_map["Business Phone 2"] = BUSINESSPHONE2_COL_IDX
	column_name_map["Home Fax"] = HOMEFAX_COL_IDX
	column_name_map["Home Phone"] = HOMEPHONE_COL_IDX
	column_name_map["Home Phone 2"] = HOMEPHONE2_COL_IDX
	column_name_map["Other Phone"] = OTHERPHONE_COL_IDX
	column_name_map["Mobile Phone"] = MOBILEPHONE_COL_IDX
	column_name_map["Pager"] = PAGER_COL_IDX
	column_name_map["Home Address"] = HOMEADDRESS_COL_IDX
	column_name_map["Business Address"] = BUSINESSADDRESS_COL_IDX
	column_name_map["Other Address"] = OTHERADDRESS_COL_IDX
	column_name_map["E-mail Address"] = E_MAILADDRESS_COL_IDX
	column_name_map["E-mail 2 Address"] = E_MAIL2ADDRESS_COL_IDX
	column_name_map["E-mail 3 Address"] = E_MAIL3ADDRESS_COL_IDX
	column_name_map["E-mail 4 Address"] = E_MAIL4ADDRESS_COL_IDX
	column_name_map["E-mail 5 Address"] = E_MAIL5ADDRESS_COL_IDX
	column_name_map["E-mail 6 Address"] = E_MAIL6ADDRESS_COL_IDX
	column_name_map["E-mail 7 Address"] = E_MAIL7ADDRESS_COL_IDX
	column_name_map["E-mail 8 Address"] = E_MAIL8ADDRESS_COL_IDX
	column_name_map["E-mail 9 Address"] = E_MAIL9ADDRESS_COL_IDX
	column_name_map["E-mail 10 Address"] = E_MAIL10ADDRESS_COL_IDX
	column_name_map["E-mail 11 Address"] = E_MAIL11ADDRESS_COL_IDX
	column_name_map["E-mail 12 Address"] = E_MAIL12ADDRESS_COL_IDX
	column_name_map["E-mail 13 Address"] = E_MAIL13ADDRESS_COL_IDX
	column_name_map["E-mail 14 Address"] = E_MAIL14ADDRESS_COL_IDX
	column_name_map["E-mail 15 Address"] = E_MAIL15ADDRESS_COL_IDX
	column_name_map["E-mail 16 Address"] = E_MAIL16ADDRESS_COL_IDX
	column_name_map["E-mail 17 Address"] = E_MAIL17ADDRESS_COL_IDX
	column_name_map["E-mail 18 Address"] = E_MAIL18ADDRESS_COL_IDX
	column_name_map["E-mail 19 Address"] = E_MAIL19ADDRESS_COL_IDX
	column_name_map["E-mail 20 Address"] = E_MAIL20ADDRESS_COL_IDX
	column_name_map["E-mail 21 Address"] = E_MAIL21ADDRESS_COL_IDX
	column_name_map["E-mail 22 Address"] = E_MAIL22ADDRESS_COL_IDX
	column_name_map["E-mail 23 Address"] = E_MAIL23ADDRESS_COL_IDX
	column_name_map["E-mail 24 Address"] = E_MAIL24ADDRESS_COL_IDX
	column_name_map["E-mail 25 Address"] = E_MAIL25ADDRESS_COL_IDX
	column_name_map["E-mail 26 Address"] = E_MAIL26ADDRESS_COL_IDX
	column_name_map["E-mail 27 Address"] = E_MAIL27ADDRESS_COL_IDX
	column_name_map["E-mail 28 Address"] = E_MAIL28ADDRESS_COL_IDX
	column_name_map["E-mail 29 Address"] = E_MAIL29ADDRESS_COL_IDX
	column_name_map["E-mail 30 Address"] = E_MAIL30ADDRESS_COL_IDX
	column_name_map["E-mail 31 Address"] = E_MAIL31ADDRESS_COL_IDX
	column_name_map["E-mail 32 Address"] = E_MAIL32ADDRESS_COL_IDX
	column_name_map["E-mail 33 Address"] = E_MAIL33ADDRESS_COL_IDX
	column_name_map["E-mail 34 Address"] = E_MAIL34ADDRESS_COL_IDX
	column_name_map["E-mail 35 Address"] = E_MAIL35ADDRESS_COL_IDX
	column_name_map["E-mail 36 Address"] = E_MAIL36ADDRESS_COL_IDX
	column_name_map["E-mail 37 Address"] = E_MAIL37ADDRESS_COL_IDX
	column_name_map["E-mail 38 Address"] = E_MAIL38ADDRESS_COL_IDX
	column_name_map["E-mail 39 Address"] = E_MAIL39ADDRESS_COL_IDX
	column_name_map["E-mail 40 Address"] = E_MAIL40ADDRESS_COL_IDX
	column_name_map["E-mail 41 Address"] = E_MAIL41ADDRESS_COL_IDX
	column_name_map["E-mail 42 Address"] = E_MAIL42ADDRESS_COL_IDX
	column_name_map["E-mail 43 Address"] = E_MAIL43ADDRESS_COL_IDX
	column_name_map["E-mail 44 Address"] = E_MAIL44ADDRESS_COL_IDX
	column_name_map["E-mail 45 Address"] = E_MAIL45ADDRESS_COL_IDX
	column_name_map["E-mail 46 Address"] = E_MAIL46ADDRESS_COL_IDX
	column_name_map["E-mail 47 Address"] = E_MAIL47ADDRESS_COL_IDX
	column_name_map["E-mail 48 Address"] = E_MAIL48ADDRESS_COL_IDX
	column_name_map["E-mail 49 Address"] = E_MAIL49ADDRESS_COL_IDX
	column_name_map["E-mail 50 Address"] = E_MAIL50ADDRESS_COL_IDX
	column_name_map["E-mail 51 Address"] = E_MAIL51ADDRESS_COL_IDX
	column_name_map["E-mail 52 Address"] = E_MAIL52ADDRESS_COL_IDX
	column_name_map["E-mail 53 Address"] = E_MAIL53ADDRESS_COL_IDX
	column_name_map["E-mail 54 Address"] = E_MAIL54ADDRESS_COL_IDX
	column_name_map["E-mail 55 Address"] = E_MAIL55ADDRESS_COL_IDX
	column_name_map["E-mail 56 Address"] = E_MAIL56ADDRESS_COL_IDX
	column_name_map["E-mail 57 Address"] = E_MAIL57ADDRESS_COL_IDX
	column_name_map["E-mail 58 Address"] = E_MAIL58ADDRESS_COL_IDX
	column_name_map["E-mail 59 Address"] = E_MAIL59ADDRESS_COL_IDX
	column_name_map["E-mail 60 Address"] = E_MAIL60ADDRESS_COL_IDX
	column_name_map["E-mail 61 Address"] = E_MAIL61ADDRESS_COL_IDX
	column_name_map["E-mail 62 Address"] = E_MAIL62ADDRESS_COL_IDX
	column_name_map["E-mail 63 Address"] = E_MAIL63ADDRESS_COL_IDX
	column_name_map["E-mail 64 Address"] = E_MAIL64ADDRESS_COL_IDX
	column_name_map["E-mail 65 Address"] = E_MAIL65ADDRESS_COL_IDX
	column_name_map["E-mail 66 Address"] = E_MAIL66ADDRESS_COL_IDX
	column_name_map["E-mail 67 Address"] = E_MAIL67ADDRESS_COL_IDX
	column_name_map["E-mail 68 Address"] = E_MAIL68ADDRESS_COL_IDX
	column_name_map["E-mail 69 Address"] = E_MAIL69ADDRESS_COL_IDX
	column_name_map["Website Home-Page"] = WEBSITEHOME_PAGE_COL_IDX
	column_name_map["Web Page"] = WEBPAGE_COL_IDX
	column_name_map["Website Blog"] = WEBSITEBLOG_COL_IDX
	column_name_map["Website Profile"] = WEBSITEPROFILE_COL_IDX
	column_name_map["Website Home"] = WEBSITEHOME_COL_IDX
	column_name_map["Website Work"] = WEBSITEWORK_COL_IDX
	column_name_map["Website Other"] = WEBSITEOTHER_COL_IDX
	column_name_map["Website FTP"] = WEBSITEFTP_COL_IDX
}