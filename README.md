# system-dashboard-pinger
Golang api to store Sets of IP addresses and then ping them to monitor network health. 

The school I work for could use a way to monitor which access points, servers and routers are currenting responding to network traffic. 

I started with a pinger adapted from https://github.com/tatsushid/go-fastping and modified it to ping multiple ip addresses via command line arguments. 

I then built a small rest API with go to store ip addresses and related data into mongoDB. 

I am currently adapting the pinger module to work from data taken from MongoDB instead of command line arguments. 

The front end is also in progress and will be a few simple react components that show the building that the router/access point is in and will ping the WIFI access points in that building, on click. 

There will be some flashy javascript animations on top of that to make it look cool.

The end. 

