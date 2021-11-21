import json
import subprocess as sp
import time

with open('/home/z1gzag/projects/radio.json', 'r') as jf:
    data = json.load(jf)


for station in data['stations']:
    name, link = station['name'], station['link']

    data = sp.Popen(f'curl -i {link}', stdout=sp.PIPE,
                    stdin=sp.PIPE, shell=True)

    break
