env GOOS=darwin GOARCH=amd64 go build ../
xattr -c CloudlogTCI
cp ../config.yaml .
zip CloudlogTCI-macOS.zip CloudlogTCI config.yaml
rm CloudlogTCI