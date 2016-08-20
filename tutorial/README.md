# Shared Contacts Admin (SCA)

### Manual on how to install SCA on GAE with google apps

### Originally published 17 April 2016
Updated 20 August 2016

Credit:
This tutorial is based off the work and slides originally created by
Edwin Vinas, 
ULAPPH Corporation

### Download installer zip file

From:
https://github.com/takbok/shared-contacts-admin
click the green "Clone or download" button followed by "Download ZIP".  Then Unzip the downloaded ZIP file.

![img](2_download_installer_zip.png)

Unzip all files under "shared-contacts-admin-master" folder

### Take note of the file "oauth.go" which contains authentication details

* Note that the ClientID and ClientSecret can be configured on [https://console.cloud.google.com](https://console.cloud.google.com)

* Since there are multiple settings to be done in Developers Console and Google Apps, they will be covered separately

* For now, don't change anything yet on the oauth.go, we will try to compile & upload first to appengine


		var (
		    config = &oauth2.Config{
		        ClientID:     `?????????-??????????.apps.googleusercontent.com`,
		        ClientSecret: `?????????????`,
		        RedirectURL:  `ideally-should-be-set-later`,
		        Scopes:       []string{`http://www.google.com/m8/feeds/contacts/`},
		        Endpoint:     google.Endpoint,
		    }
		
		    yeah = "yeah"
		)### Create an appengine project on https://console.cloud.google.com

![img](4_create_an_appengine_project.png)

Take note of the email address and project ID

### Remove unneeded folders & download required golang packages

After un-zipping the ZIP delete the following three folders:

* tutorial
* video-tutorial
* vendor

See Issue # 9 for background on the reason for deleting the "vendor" folder.  https://github.com/takbok/shared-contacts-admin/issues/9

Because the "vendor" folder has been removed, some needed golang packages must now down be downloaded.  If using Microsoft Windows, run

* "install-packages.bat"

to install the required golang packages.

### Open a command prompt & upload app to appengine

![img](5_open_a_command_prompt_&_upload_app_to_appengine.png)

appcfg.py --application=PROJ_ID --email=EMAIL@EMAIL.COM --no_cookies update shared-contacts-admin-master

### After successful upload, try the website URL

![img](6_after_successful_upload_try_the_website_url.png)

This is the one-page web interface of the SCA

### The Imports section

![img](7_the_imports_section.png)

This is where you can Import Contacts

### The Exports Section

![img](8_the_exports_section.png)

This is where you can Export Contacts

### The Delete Section

![img](9_the_delete_section.png)

This is where you can Delete Contacts

### Now you have  working URL but it is not yet usable...

* As expected, there will be errors when you try the SCA as is...

* What needs to be done so we can make use of this app?

1) Setup Google Developer Console

2) Setup Google Apps for Work

## Setup Google Developer Console (Part 1)

### Enable Contacts API

![img](12_enable_contacts_api.png)

You must have the Contacts API enabled

### The API Credentials

![img](13_the_api_credentials.png)

Note that an entry was automatically added by Google Apps to enable sharing of contacts

### Looking at the "domain-wide delegation"

![img](14_looking_at_the_domain-wide_delegation.png)

This will appear only after configuring Google Apps

### Redirect URIs

* http://g-apps-sca001.appspot.com/contacts/export

* http://g-apps-sca001.appspot.com/import/do

* http://g-apps-sca001.appspot.com/contacts/exportxml

* http://g-apps-sca001.appspot.com/contacts/delete

### Looking at the oAuth Keys

![img](16._redirect_uris.png)

No need to download JSON, just update the oauth.go

### Update Consent screen

![img](17_update_consent_screen.png)

You can change it to SCA; Nagitgit is just a sample name

## Setup Google Apps

### Go to Google Apps http://apps.google.com as admin

![img](19_1460748308.png)

Go to Users to configure contacts sharing

### Click admin user

![img](20_click_admin_user.png)

Go to admin user

### Click to show apps

![img](21_click_to_show_apps.png)

Go to apps list

### Click to configure contacts

![img](22_configure_contacts.png)

Go to contacts

### Click to go to advanced settings

![img](23_advanced_settings.png)

Go to advanced settings

### Enable options for Contacts Sharing

![img](24_enable_options_for_contacts_sharing.png)

Enable contacts sharing

## Setup Google Developer Console (Part 2)

### Manage Service Accounts

![img](26_manage_service_accounts.png)

Click to manage service accounts

### Select App Engine service account

![img](27_select_app_engine_service_account.png)

Click to Edit App Engine service account

### Enable Google Apps Domain-Wide Delegation

![img](28_enable_google_apps_domain-wide_delegation.png)

Enable the domain-wide delegation

### The Newly Added Service account client

![img](29_newly_added_service_account_client.png)

Notice that a new client has been created automatically for the service account

## APIs and Google Apps settings done!

### Now recompile Go project and upload to appengine

![img](31_recompile_and_upload_go.png)

appcfg.py --application=PROJ_ID --email=EMAIL@EMAIL.COM --no_cookies update shared-contacts-admin-master

## Testing the SCA

## Importing contacts from CSV

### Go to the Imports section

![img](34_imports_section.png)

Go to the Imports section and click Import Contacts

### Point to CSV file and Enter Google Apps Domain

![img](35_point_csv_and_enter_google_app_domain.png)

Point to CSV contacts file and Enter Google Apps Domain

### Results will show successful imports

![img](36_results_will_show_successful_imports.png)

Number of records imported will be shown

## Exporting contacts to CSV

### Go to Exports Section

![img](38_exports_section.png)

Go to Exports section and click Export CSV

### Enter Google Apps Domain

![img](39_enter_google_apps_domain.png)

Enter Google Apps Domain

### CSV File will be downloaded

![img](40_csv_file_downloaded.png)

Double-click on the CSV download item

### Open the exported excel file

![img](41_open_the_exported_excel_file.png)

Contacts successfully exported to Excel

## Exporting contacts to XML

### Go to Exports Section

![img](43_exports_section.png)

Go to Exports section and click Export CSV

### Enter Google Apps domain name

![img](44_enter_google_apps_domain_name.png)

Enter Google Apps domain

### XML file will be downloaded

![img](45_xml_file_downloaded.png)

Double-click on the XML download item

### Open the exported XML file

![img](46_open_the_exported_xml_file.png)

Contacts successfully exported to XML

## Viewing the Shared Contacts in contacts.google.com

### Visit Contacts https://contacts.google.com (as admin)

![img](48_visit_contacts.google.com_admin.png)

Contacts are now visible on the domain

### Visit Contacts https://contacts.google.com (as ordinary user)

![img](49_visit_contacts.google.com_user.png)

Contacts are now visible on the domain

## Deleting Contacts

### Go to the Delete Contacts section

![img](51_delete_contacts_section.png)

Go to Delete Contacts section

### Enter Google Apps domain

![img](52_enter_google_apps_domain.png)

Enter the Google Apps domain

### Results will show batch delete request

![img](53_batch_delete_request_result.png)

Records were sent for batch delete request
