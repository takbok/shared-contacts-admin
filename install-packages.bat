@ECHO OFF

IF "%GOPATH%"=="" (
	ECHO GOPATH Environment variable not define.
	EXIT
)

REM SHOULD CHECK FOR GO INSTALLATION ALSO

ECHO Performing package update\installation. Please wait....

ECHO Installing\Updating 'google.golang.org/appengine'
go get google.golang.org/appengine
ECHO Installing\Updating 'google.golang.org/appengine/urlfetch'
go get google.golang.org/appengine/urlfetch
ECHO Installing\Updating 'golang.org/x/oauth2'
go get golang.org/x/oauth2
ECHO Installing\Updating 'golang.org/x/oauth2/google'
go get golang.org/x/oauth2/google

ECHO Package update\installation done.