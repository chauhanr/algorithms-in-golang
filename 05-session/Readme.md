# Sessions with Golang web 

Session is about associating the user with a server state so that the server can remember things
about a user like user session state. 

The most popular way to handle user session is to use Cookies. The Cookie information will have a
session id (uuid) that uniquely identifies the user. The user information will be stored at the
server side. The cookies handler will look for uuid session and determine the user information. 



