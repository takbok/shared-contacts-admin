** Author 		: Prakhar Kumar
** Date Created : 02/10/2016
** What is tested?
** Importing external contacts & Export contacts

Import External Contacts via CSV
-----

1. Download [sample test data](https://github.com/takbok/shared-contacts-admin/blob/master/testcases/test-data/sample2.csv).
2. Open [GAE Test](http://www.gae-test1.com/) in your browser.
3. Click on 'Choose File' or 'Browse' button depending on your browser.
4. Select the downloaded file from step # 1.
5. Enter the domain name, that is hosted on Google Apps, to which the contacts need to be imported to.
6. Click on 'Import CSV' button.
7. On the OAuth consent screen, enter the credentials that have the necessary permissions on the domain from #5 above.
8. Results should be displayed on screen as below.

![Import Results](https://raw.githubusercontent.com/takbok/shared-contacts-admin/master/testcases/images/test-case-04-01.jpeg)

Export Contacts to CSV
-----

1. Open [GAE Test](http://www.gae-test1.com) in your browser.
2. Enter the domain name, that is hosted on Google Apps, and for which Domain Shared Contacts need to be exported.
3. Click on 'Set Domain & Export CSV'
4. On the OAuth consent screen, enter the credentials that have the necessary permissions on the domain from #2 above.
5. A CSV file containing the contacts will be downloaded
