# Install
```shel
go install github.com/gucchi0421/goya@latest
```

### Whois Search
```
$ goya whois example.com

Domain Name: EXAMPLE.COM
   Registry Domain ID: 2336799_DOMAIN_COM-VRSN
   Registrar WHOIS Server: whois.iana.org
   Registrar URL: http://res-dom.iana.org
   Updated Date: 2024-08-14T07:01:34Z
   Creation Date: 1995-08-14T04:00:00Z
   Registry Expiry Date: 2025-08-13T04:00:00Z
   Registrar: RESERVED-Internet Assigned Numbers Authority
   Registrar IANA ID: 376
   Registrar Abuse Contact Email:
   Registrar Abuse Contact Phone:
etc...
```

### DNS Record Search
```
$ goya dns example.com

DNS for example.com

IPv4 : 93.184.215.14
IPv6 : 2606:2800:21f:cb07:6820:80da:af6b:8b2c
A    : 93.184.215.14
A    : 2606:2800:21f:cb07:6820:80da:af6b:8b2c
NS1  : a.iana-servers.net.
NS2  : b.iana-servers.net.
MX   : .
CNAME: example.com.
TXT  : wgyf8z8cgvm2qmxpnbnldrcltvk4xqfn
TXT  : v=spf1 -all
```

### Password Generater
```
$ goya newpass
fi2AwqJjgh4y

$ goya newpass -l=18
6rzVMxST9w4Vlqoyll

$ goya newpass -c=2
BXZjXEvDiFOQ
e98DwTCnodEL
```