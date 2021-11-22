import json
import time

import requests
from requests.exceptions import ConnectTimeout, RequestException

with open('/home/z1gzag/projects/radio.json', 'r') as jf:
    data = json.load(jf)


i = 0

for s in data['stations']:
    time.sleep(2)
    
    n, l = s.values()

    try:
        resp = requests.get(
            l,
            allow_redirects=False,
            timeout=0.0001
        )

        if (code := resp.status_code) in (200, 302):
            print(code)

        else:
            print(resp.status_code)

        resp.close()
    except ConnectTimeout as ct:
        pass

    


