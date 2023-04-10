### dnsTools
dns test tool

### usage

###### total dns domain delay
``` bash
~/github/dnsTools  ./dnsTools speed -d github.com -e 1

domain:github.com., server:20.27.177.113, delay:5ms
domain:github.com., server:20.27.177.113, delay:3ms
domain:github.com., server:20.27.177.113, delay:2ms
domain:github.com., server:20.27.177.113, delay:2ms
domain:github.com., server:20.27.177.113, delay:2ms
domain:github.com., server:20.27.177.113, delay:2ms
domain:github.com., server:20.27.177.113, delay:4ms
domain:github.com., server:20.27.177.113, delay:2ms
domain:github.com., server:20.27.177.113, delay:2ms
domain:github.com., server:20.27.177.113, delay:2ms
dnsTools Total:10, Succ:10, Error:0, averageTime:2ms
```

###### show dns domain detail information
```bash
.~/github/dnsTools /dnsTools dig -d github.com

;; opcode: QUERY, status: NOERROR, id: 58535
;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;github.com.    IN       A

;; ANSWER SECTION:
github.com.     1       IN      A       20.27.177.113

dnsTools delay:5ms
```