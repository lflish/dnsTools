debug:
	go build -o dnsTools dnsTools.go speed.go dig.go

release: dnsTools dnsTools.exe
dnsTools:
	CGO_ENABLED=0 go build -o $@ -a -ldflags  "-extldflags '-static'" dnsTools.go
dnsTools.exe:
	CGO_ENABLED=0 GOOS=windows go build -o $@ -a -ldflags  "-extldflags '-static'" dnsTools.go

clean:
	rm -rf dnsTools
	rm -rf dnsTools.exe
