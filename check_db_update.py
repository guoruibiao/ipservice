#!/usr/bin python
# coding: utf8
# `deprecated` useless script.
import urllib2
url = "https://www.ipip.net/product/client.html"
headers = {
    'User-Agent':'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36',
    "Host": "www.ipip.net",
    "Cookie": "__jsluid=f2935cca25b67df2c97c0496a718ece3; _ga=GA1.2.262138422.1552457307; LOVEAPP_SESSID=32da3602a71215abe5aa52ad04cf90ad222e026e; _gid=GA1.2.1189292749.1552975476; Hm_lvt_123ba42b8d6d2f680c91cb43c1e2be64=1552457307,1552975476; lang=CN; login_r=https%253A%252F%252Fwww.ipip.net%252Fproduct%252Fclient.html; user=Q6YFFzEZcGLl%7COgsggq%7CLpkdZ1z7Az08bxkXaK2r5HMjOPm0%29I8Y1SiNqtHoI%29GWtm3jdA; Hm_lpvt_123ba42b8d6d2f680c91cb43c1e2be64=1553047476"
}
req = urllib2.Request(url, headers=headers)
content = urllib2.urlopen(url)
print(content.read())
