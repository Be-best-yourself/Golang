# -*- coding: utf-8 -*-  
import sys,re  
import urllib,urllib.request  
url = "https://degoo.com/g/19zdWfE?_branch_match_id=433423193969333623"  
params = urllib.parse.urlencode([('Email', 'fasdfsdfsdsfsdfs124453@qq.com'), ('Password', '123456789')]).encode(encoding='GB2312')  
req = urllib.request.Request(url)  
f = urllib.request.urlopen(req,params)  