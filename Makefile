.PHONY: win lin mac

name := encrypt-js
path :=


all: win lin

win:
	GOOS=windows GOARCH=amd64 go build -o out/$(name).exe -trimpath -ldflags="-s -w" $(path)
	$(call bak)

lin:
	GOOS=linux GOARCH=amd64 go build -o out/$(name) -trimpath -ldflags="-s -w" $(path)
	$(call bak)

mac:
	GOOS=darwin GOARCH=amd64 go build -o out/$(name) -trimpath -ldflags="-s -w" $(path)
	$(call bak)




define bak
	cp -R ./bak/* ./out/
endef

#	upx64 out/main -f -o out/$(name)
#	rm -rf out/main