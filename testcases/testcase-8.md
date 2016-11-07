_Author_        : Alexander Dolgunin  
_Date Created_  : 11/04/2016

**What is tested?**  
_Importing and exporting specific fields from and to CSV_

Import External Contacts from CSV
-----
1. Download [sample test data](https://github.com/takbok/shared-contacts-admin/blob/master/testcases/test-data/sample1.csv).
2. Edit the first row to add birthday in format: YYYY-MM-DD and take note of the email in the row (e.g. testmail46@test.com)
3. Open [GAE Test](http://www.gae-test1.com/) in your browser.
4. Click on 'Choose File' or 'Browse' button depending on your browser.
5. Select the downloaded and edited file from step #2.
6. Opposite "Import CSV" button, enter the domain name, that is hosted on Google Apps, to which the contacts need to be imported to.
7. Click "Import CSV" button.
8. On the OAuth consent screen, enter the credentials that have the necessary permissions on the domain from #6 above and click Allow button.
![Import Results](https://raw.githubusercontent.com/takbok/shared-contacts-admin/master/testcases/images/Screenshot-Request%20for%20Permission%20Shared%20Contacts.png)

9. The output of web app should be:
```
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
Result: 201 Created
```
(status 201 for each processed row).

10. With the same Google Account, visit [Google Contacts](https://www.google.com/contacts/?hl=en#contacts) and click Directory.
11. Click at the contact you edited in CSV, e.g. testmail46@test.com
12. Birthday field is blank - set it to some other date in relation to #2.

Export changed records to CSV
-----
1. Open [GAE Test](http://www.gae-test1.com) in your browser.
2. Opposite "Export CSV" button, enter the domain name from the previous test.
3. Click "Export CSV".
4. The order of returned row is different. Search for testmail46@test.com
5. Birthday field in CSV keeps the date you set at #2 of previous test, not the one you edited in #12.

This is true for many other fields, like phone numbers, addresses etc.
