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
### reindeer and SnowWhite publishes, only reindeer watches SnowWhite

https://media.raku.cloud:7889/admin/
https://media.raku.cloud:7889/admin/2139542948799017
https://media.raku.cloud:7889/admin/2139542948799017/93507538637076
https://media.raku.cloud:7889/admin/2139542948799017/6053503735469801
https://media.raku.cloud:7889/admin/2139542948799017/1685459559131250
https://media.raku.cloud:7889/admin/2139542948799017/1078904380005661
https://media.raku.cloud:7889/admin/4711353647645573
https://media.raku.cloud:7889/admin/4711353647645573/3254409020427425

User: Display reindeer ID 3823503906827241 PvtID 2492745038  Session 2139542948799017
publishes:
Using handle 93507538637076 in Room 1234
subscribes to:
Using handle 1078904380005661 in  Room 1234 to SnowWhite with ID 464261124587590 PvtID 2492745038
Listeners:


User: Display SnowWhite ID 464261124587590 PvtID 3062447485  Session 4711353647645573
publishes:
Using handle 3254409020427425 in Room 1234
subscribes to:
Listeners:
reindeer listens on SnowWhite


Process finished with exit code 0
