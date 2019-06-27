# app.py

import requests

res = requests.get('http://observatory-service/')
#res = requests.get('http://gohttp-service/')
data = res.json()
print("SAHIL:resquestapi",data)
#print("SAHIL:resquestapi",res)
