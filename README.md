# Go-loginclient

## V2 UPDATE:
A server written in go is now added. The server currently serves HTTP only as its running on localhost. To serve HTTPS simply replace http.ListenAndServe with http.ListenAndServeTLS with the link to your SSL certs.
In version 2 [JWT](https://jwt.io/introduction/) is added as a functionality for authorizing users.
The hashing method is changed to Sha256.
## Introduction V1
A simple login/register GUI made with Go and an API written in Nodejs. The goals of this project was to gain experience in writing apps with Go, designing and implementing an self-hosted API, and to gain knowledge in writing cross-platform GUI apps with Go-astilectron.

Some basic security features/measures for V1: 
 * The passwords are all hashed and salted with Scrypt Kdf. 
 * The requests to the API server are only HTTPS.
 * All database querys are parameterized. 
 
 
 ## Snapshots:
 ![img](https://imgur.com/2i0tKWP.png)
 ![img](https://imgur.com/YVRTVGi.png)
 ![img](https://imgur.com/IZHuZGQ.png)
 ![img](https://imgur.com/PT3bTs6.png)
 
 
 ## V2 SNAPSHOTS:
 ![img](https://imgur.com/ZKhMuto.png)
 ![img](https://imgur.com/NGpK7fi.png)
 ![img](https://imgur.com/pZHAiL6.png)
   
 
