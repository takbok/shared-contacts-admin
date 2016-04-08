# Shared Contact Admin (SCA)
A web based application for administering Google Apps Domain Shared Contacts.

* Live Demo: http://g-apps-sca001.appspot.com/
* Slides: http://goo.gl/SqRpHT

This is a Google App Engine (GAE) application written in Golang that provides a web based user interface for importing, exporting and deleting Google Apps shared contacts. It is based on the Google Domain Shared Contacts API:

https://developers.google.com/admin-sdk/domain-shared-contacts/

What is Domain Shared Contacts API?
===========

The [Shared Contacts API](https://developers.google.com/admin-sdk/domain-shared-contacts/) allows client applications to retrieve and update external contacts that are shared to all users in a Google Apps domain. Shared contacts are visible to all users of an Apps domain and all Google services have access to the contact list. To retrieve and update Google Apps domain users contact information, use the Directory API instead.

The Domain Shared Contacts API is only available to Google Apps for Business and Education accounts. The contacts sharing is disabled by default and this can be enabled by the Admin of the domain.

**API Features**

http://developers.google.com/admin-sdk/domain-shared-contacts/#Creating

> Creating shared contacts **
> Retrieving shared contacts  **
> Retrieving shared contacts using query parameters 
> Retrieving a single shared contact
> Retrieving a photo for a contact 
> Updating shared contacts **
> Shared contact photo management 
> Deleting shared contacts Batch operations **

** means SCA is now supporting these functionalities

Build Instructions
===========

For specific GAE setup, please read first the documentation: 

> https://cloud.google.com/appengine/docs/go/

Prepare
-------
  

>   go get golang.org/x/oauth2
>      go get golang.org/x/oauth2/google
>      go get google.golang.org/appengine
>      go get google.golang.org/appengine/urlfetch

Detailed Steps
---------

**Donwload & Install SCA**

1. Download installer zip file
2. Unzip all files under "shared-contacts-admin-master" folder
3. Take note of the file "oauth.go" which contains authentication details

> * Note that the ClientID and ClientSecret can be configured on https://console.cloud.google.com
> * Since there are multiple settings to be done in Developers Console and Google Apps, they will be covered separately
> * For now, don't change anything yet on the oauth.go, we will try to compile & upload first to appengine

<pre>
    var (
        config = &oauth2.Config{
            ClientID:     `?????????-??????????.apps.googleusercontent.com`,
            ClientSecret: `?????????????`,
            RedirectURL:  `ideally-should-be-set-later`,
            Scopes:       []string{`http://www.google.com/m8/feeds/contacts/`},
            Endpoint:     google.Endpoint,
        }
    
        yeah = "yeah"
    )
</pre>

4. Create an appengine project on https://console.cloud.google.com
* Take note of the email address and project ID
5. Open a command prompt & upload app to appengine
6. Execute upload and compile command to appengine

> * appcfg.py --application=PROJ_ID --email=EMAIL@EMAIL.COM --no_cookies update shared-contacts-admin-master

7. After successful upload, try the website URL

> * http://PROJECT-ID.appspot.com
> * The Imports section - This is where you can Import Contacts
> * The Exports Section - This is where you can Export Contacts
> * The Delete Section - This is where you can Delete Contacts

8. Now you have working URL but it is not yet usable...

> * As expected, there will be errors when you try the SCA as is...
> * What needs to be done so we can make use of this app? We need to setup Google Developer Console and also we need to setup Google Apps
> for Work

**Setup Google Developer Console (Part 1)**

> * Enable Contacts API
> * The API Credentials - contains Open Auth 2.0 client ID/secret
> * Add the following redirect URIs (replace with your project ID):
> Blockquote

<pre>
http://g-apps-sca001.appspot.com/contacts/export
http://g-apps-sca001.appspot.com/import/do
http://g-apps-sca001.appspot.com/contacts/exportxml
http://g-apps-sca001.appspot.com/contacts/delete
</pre>
* Put the client ID and client secret in the **oauth.go** program
* Update Consent screen

**Setup Google Apps for Work/Education**

> * Go to Google Apps http://apps.google.com as admin
> * Go to Users to configure contacts sharing
> * Click admin user
> * Click to show Google apps enabled
> * Click to configure contacts
> * Click to go to advanced settings
> * Enable options for Contacts Sharing

**Setup Google Developer Console (Part 2)**

> * Manage Service Accounts
> * Select App Engine service account
> * Enable Google Apps Domain-Wide Delegation
> * Notice that a new client has been created automatically for the service account

**APIs and Google Apps settings done!**

> * Now recompile Go project and upload to appengine

Using the SCA
---------
**1. Importing contacts from CSV**

> * Go to the Imports section
> * Point to CSV file and Enter Google Apps Domain
> * Results will show successful imports

**2. Exporting contacts to CSV**

> * Go to Exports Section
> * Enter Google Apps Domain
> * CSV File will be downloaded
> * Open the exported excel file

**3. Exporting contacts to XML**

> * Go to Exports Section
> * Enter Google Apps domain name
> * XML file will be downloaded
> * Open the exported XML file

**4. Deleting Contacts**

> * Go to the Delete Contacts section
> * Enter Google Apps domain
> * Results will show batch delete request

Viewing the Shared Contacts
---------
1. Visit Contacts https://contacts.google.com (as admin)

> * Click on the Directory
> * Contacts are now visible on the domain