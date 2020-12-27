env GOOS=windows GOARCH=amd64 go build ../
cp ../config.yaml .
zip CloudlogTCI-windows-amd64.zip CloudlogTCI.exe config.yaml
rm CloudlogTCI.exe