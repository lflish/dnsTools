debug:
	go build -o dnsTools dnsTools.go speed.go dig.go

release: dnsTools-linux dnsTools.exe

dnsTools-linux:
	CGO_ENABLED=0 go build -o $@ -a -ldflags  "-extldflags '-static'" dnsTools.go speed.go dig.go

dnsTools.exe:
	CGO_ENABLED=0 GOOS=windows go build -o $@ -a -ldflags  "-extldflags '-static'" dnsTools.go speed.go dig.go

clean:
	rm -rf dnsTools
	rm -rf dnsTools-linux
	rm -rf dnsTools.exe
