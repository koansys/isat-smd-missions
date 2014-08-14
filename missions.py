#!/usr/bin/env python
# Needs BeautifulSoup or lxml or similar

import requests
from lxml import etree

import lxml.html
from lxml.cssselect import CSSSelector

MISSIONS_URL = 'http://science.nasa.gov/missions/?group=all'

res = requests.get(MISSIONS_URL)
if res.status_code != 200:
    raise RuntimeError('Failed to get url=%s' % MISSIONS_URL)
html = res.text
tree = lxml.html.fromstring(res.text)
sel = CSSSelector('table tbody tr')
rows = sel(tree)
print "Row results: ", len(rows)
num_operating = 0
for row in rows:
    # This is unrealiable; I don't know how to get just the text 'Operating':
    #  phase = CSSSelector('td:nth-of-type(4)')(row)[0]
    #  lxml.html.tostring(phase)
    #  '<td><span class="hide">3</span>Operating</td>'
    phase = CSSSelector('td:nth-of-type(4)')(row)[0].text_content() # '3Operating'
    if 'Operating' in phase:
        num_operating += 1
    # Show phase because we may have stale ones in iSat
    division = CSSSelector('td:nth-of-type(1)')(row)[0]
    try:
        division = division.text.strip()
    except AttributeError, e:
        division = 'NOTFOUND'
    mission = CSSSelector('td:nth-of-type(2) > a')(row)[0]
    mission_name = mission.text.strip()
    mission_url = mission.get('href') # /missions/xmm-newton/
    mission_slug = mission_url.split('/')[2]
    num_operating += 1
    try:
        print '%-30s\t%-40s\t%-20s\t%-20s' % (mission_slug, mission_name.encode('ascii', 'ignore'), division, phase)
    except UnicodeEncodeError, e:
        print "Fucking unicode problem: ", e
        import pdb; pdb.set_trace()
print 'Operating:', num_operating

