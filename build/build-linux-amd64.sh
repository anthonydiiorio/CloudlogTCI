env GOOS=linux GOARCH=amd64 go build ../
cp ../config.yaml .
zip CloudlogTCI-linux-amd64.zip CloudlogTCI config.yaml
rm CloudlogTCI