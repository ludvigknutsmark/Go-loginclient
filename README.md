# Go-loginclient

/* This program was written when I was a beginner in both Nodejs and Go, therefore some functions are spaghetti, I'm currently working on a version 2 with a REST-api written in Go and uses Json Web Tokens(JWT) for authenticating users on protected resources. */
A simple login/register GUI made with Go and an API written in Nodejs. The goals of this project was to gain experience in writing apps with Go, designing and implementing an self-hosted API, and to gain knowledge in writing cross-platform GUI apps with Go-astilectron.

Some basic security features/measures: 
 * The passwords are all hashed and salted with Scrypt Kdf. 
 * The requests to the API server are only HTTPS.
 * All database querys are parameterized. 
 
 Snapshots:
 
 
 ![img](https://i.imgur.com/2i0tKWP.png)[/img]
 ![img](https://imgur.com/YVRTVGi.png)[/img]
 ![img](https://imgur.com/IZHuZGQ.png)[/img]
 ![img](https://imgur.com/PT3bTs6.png)[/img]
 
 
