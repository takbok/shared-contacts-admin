#!/bin/bash

if [ "$GOPATH" == "" ]; then
	echo "GOPATH Environment variable not defined."
	exit
else
	echo "GOPATH = $GOPATH"
fi

echo "Performing package update\installation. Please wait...."

echo "Installing\Updating 'google.golang.org/appengine'"
go get google.golang.org/appengine

echo "Installing\Updating 'google.golang.org/appengine/urlfetch'"
go get google.golang.org/appengine/urlfetch

echo "Installing\Updating 'golang.org/x/oauth2'"
go get golang.org/x/oauth2

echo "Installing\Updating 'golang.org/x/oauth2/google'"
go get golang.org/x/oauth2/google

echo "Package update\installation done."
