# JanusCapture
Example how to capture and parse Janus information by golang using its web-based admin interface. 

## About Janus
Janus is a WebRTC framework at https://github.com/meetecho/janus-gateway


# Prerequisits
Golang environment
Get jsong json query utility
$ go get github.com/jmoiron/jsonq

# Example usage

$ git hub clone MalinYamato/JanusCapture
$ go run capture.go

# -- sample ouput
### reindeer and SnowWhite publish, only reindeer watches SnowWhite

https://media.raku.cloud:7889/admin/
https://media.raku.cloud:7889/admin/2139542948799017
https://media.raku.cloud:7889/admin/2139542948799017/93507538637076
https://media.raku.cloud:7889/admin/2139542948799017/6053503735469801
https://media.raku.cloud:7889/admin/2139542948799017/1685459559131250
https://media.raku.cloud:7889/admin/2139542948799017/1078904380005661
https://media.raku.cloud:7889/admin/4711353647645573
https://media.raku.cloud:7889/admin/4711353647645573/3254409020427425
<<<<<<< HEAD
<br />
User: Display reindeer ID 3823503906827241 PvtID 2492745038  Session 2139542948799017 <br />
publishes: <br />
Using handle 93507538637076 in Room 1234 <br />
subscribes to: <br />
Using handle 078904380005661 in  Room 1234 to SnowWhite with ID 464261124587590 PvtID 2492745038 <br />
Listeners:<br />

User: Display SnowWhite ID 464261124587590 PvtID 3062447485  Session 4711353647645573 <br />
publishes: <br />
Using handle 3254409020427425 in Room 1234 <br />
subscribes to: <br />
Listeners: <br />
reindeer listens on SnowWhite <br />
=======

User: Display reindeer ID 3823503906827241 PvtID 2492745038  Session 2139542948799017 <br/>
publishes:<br/>
Using handle 93507538637076 in Room 1234<br/>
subscribes to: <br/>
Using handle 1078904380005661 in  Room 1234 to SnowWhite with ID 464261124587590 PvtID 2492745038 <br/>
Listeners: <br/>

User: Display SnowWhite ID 464261124587590 PvtID 3062447485  Session 4711353647645573<br/>
publishes:<br/>
Using handle 3254409020427425 in Room 1234<br/>
subscribes to:<br/>
Listeners:<br/>
reindeer listens on SnowWhite
>>>>>>> origin/master


Process finished with exit code 0
