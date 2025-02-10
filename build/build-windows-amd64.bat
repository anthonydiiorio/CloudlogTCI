@echo off
REM Get today's date in YYYY-MM-DD format
for /f "tokens=2 delims==" %%I in ('"wmic os get localdatetime /value"') do set datetime=%%I
set datestamp=%datetime:~0,4%-%datetime:~4,2%-%datetime:~6,2%

REM Build the Go project
go build ../

REM Copy the config file
copy ..\config.yaml .

REM Create a ZIP archive with today's date
powershell -Command "Compress-Archive -Path CloudlogTCI.exe, config.yaml -DestinationPath CloudlogTCI-%datestamp%-Windows-amd64.zip"

REM Clean up
del CloudlogTCI.exe
del config.yaml

pause