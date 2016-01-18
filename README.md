# shared-contacts-admin
a web based application for administering Google Apps Domain Shared Contacts

This is a GAE application written in golang that provides a web based user interface for the following API:

https://developers.google.com/admin-sdk/domain-shared-contacts/

Build Instructions
===========

https://cloud.google.com/appengine/docs/go/

Prepare
-------
  
  go get golang.org/x/oauth2
  
  go get golang.org/x/oauth2/google
  
  go get google.golang.org/appengine
  
  go get google.golang.org/appengine/urlfetch

Dev Serve
---------

  goapp serve app.yaml

Deploy
------

  goapp deploy app.yaml


