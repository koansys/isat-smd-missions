======================================
 Parse Missions from Science.nasa.gov
======================================

We need to periodically update the iSat missions for the SMD group, so
we need to pull down the missions we currently have. Do this by
parsing the HTML on the page itself, since we don't have database
access.

For each of the missions, make sure it's in the iSat SMD pull-down.

If not, I manually have to find the NORAD/COSPAR identifier for the
satellite. Some of these missions are not satellites (e.g., their
airplane-mounted instruments) or they're not Earth-orbiting satellites
(e.g., deep-space or Mars-orbiting).

Then we'll need to integrate that with where iSat is pulling the SMD satellite list. TBD.

Install
=======

Do the virtualenv and pip thing::

  virtualenv .
  pip install -r requirements.txt

Run
===

It shows the Mission URL slug, mission name, SMD division, and operating status::

  (mission-extract-data)bash-3.2$ ./missions.py
  Row results:  208
  ace                           	ACE                                     	Heliophysics        	3Operating          
  acrimsat                      	ACRIMSAT                                	Earth               	3Operating          
  ...
  xmm-newton                    	XMM-Newton                              	Astrophysics        	3Operating          
  yohkoh                        	Yohkoh                                  	Heliophysics        	4Past               
  Operating: 277

Misfeatures
===========

The operating status gets the numeric database ID and the name as a
combined item, due to how my parsing grabs the 'hide' data::

  ...<td><span class="hide">3</span>Operating</td></tr>

