_Author_        : Alexander Dolgunin  
_Date Created_  : 11/04/2016

**What is tested?**  
_Updating records by importing same CSV IDs again_

Import External Contacts from CSV
-----
1. Follow procedure described in [Test Case 8](https://github.com/takbok/shared-contacts-admin/blob/master/testcases/testcase-8.md) to import CSV file.

Re-Import changed External Contacts from CSV
-----
1. Open [sample test data](https://github.com/takbok/shared-contacts-admin/blob/master/testcases/test-data/sample1.csv) and edit it, leaving one row after the CSV headers:
```
update,3121790dfee845,Test Name 47,Test Company 47,Media Research Analyst,Test notes 47,,External,cloudtest2.com,http://www.google.com/m8/feeds/contacts/craft-bilt.com/base/3121790dfee845,1975-03-05,work,home,other,mobile,test name 47,test name 47|4823dc023f852af96a3ea583ac60a5e52,,,,,,,,http://www.example.com,work,,,Newsletters,1,Territory,HA,Type,Customer,Language,English,,,,,(111) 222-3288,(111) 111-2177,,,,,,(333) 555-6621,,,"Address #1 RR1, East Aylesford, NS B0P 46",,testmail46@test.com,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,http://test.com/,,
```
2. Edit some fields like NickName, Occupation, Department.
3. Save as sample3.csv and import again, following procedure above.
4. Even though Action was "update" in the CSV data, Directory now has a second contact with the majority of fields identical.

Export contacts after re-import
-----
1. Follow "Export changed records to CSV" procedure from [Test Case 8](https://github.com/takbok/shared-contacts-admin/blob/master/testcases/testcase-8.md)
2. The resulting CSV contains n + 1 records from sample1.csv, with the edited and updated record duplicated.
