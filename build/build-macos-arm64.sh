env GOOS=darwin GOARCH=arm64 go build ../
xattr -c CloudlogTCI
cp ../config.yaml .
zip CloudlogTCI-macOS-arm64.zip CloudlogTCI config.yaml
rm CloudlogTCI