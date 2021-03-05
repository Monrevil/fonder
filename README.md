This is a simple api server written in GO.
With CRUD operations for posts and comments.
Security is managed by jwt tokens.

Used :  
Gorm    for Database
Echo    for http
Swaggo  for docs

For swagger doc:
http://localhost:1323/swagger/index.html

OAUTH2:
Login with Google works, but app is in testing mode.
Google should only allow users, that have been registered as test useres.
/home to get login with google link.
Facebook and Twitter not yet implemented
