# Go-loginclient

A simple login/register GUI made with Go and an API written in Nodejs. The goals of this project was to gain experience in writing apps with Go, designing and implementing an self-hosted API, and to gain knowledge in writing cross-platform GUI apps with Go-astilectron.

Some basic security features/measures: 
 * The passwords are all hashed and salted with Scrypt Kdf. 
 * The requests to the API server are only HTTPS.
 * All database querys are parameterized. 
 
 Snapshots:
 ![Login](https://imgur.com/2i0tKWP)
 ![Register](https://imgur.com/YVRTVGi)
 ![Login fault](https://imgur.com/IZHuZGQ)
 ![Input](https://imgur.com/PT3bTs6)
 
 
