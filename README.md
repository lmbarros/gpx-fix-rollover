# gpx-fix-rollover

Fixes the GPS week rollover issue in GPX files by adding 1024 weeks to the timestamp of every point.

## Long story

The GPS system represents dates and times in a weird way. In particular, there is a 10-bit field representing "the week". About every 20 years this field overflows, causing some (arguably buggy) GPS receivers to go nuts. The last time this happened was in April 7th, 2019 and one of GPSs to go nuts was my beloved Holux GPSport 245 (which I have been happily using for geotagging my photographs for more than 7 years).

This program simply reads a GPX file and adds 1024 weeks to the date of each point. To be honest, I am not even sure this is a proper fix to the GPS rollover issue. Maybe there is some subtler point I'm not taking into account, but it worked well enough for my geotagging purposes.

If your are curious, there is some more information about the GPS rollover issue around the web: [here, for example](https://blog.fosketts.net/2019/04/06/gps-time-rollover-failures-keep-happening-but-theyre-almost-done/).

## Credits

By Leandro Motta Barros, but all the hard work is really done by Tomo Krajina's [gpxgo](https://github.com/tkrajina/gpxgo) library.
