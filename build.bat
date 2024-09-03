REM build the project

del bin\lfsr*

set GOOS=windows
set GOARCH=amd64
go build -o bin\lfsr.exe 

set GOOS=linux
set GOARCH=amd64
go build -o bin\lfsr


..\..\bin\gomarkdoc.exe -u > .\Docs\doc.md