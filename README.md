# yeet-fb

YEET FB AAAAAA

## brain dump 
* user auth flow 
  * user can register by invite code 
  * only allowed principals can register 
  * registered principals can login via fb oauth flow and obtain short-lived access code 
  * access tokens stored in Redis, AES ecnryption
* users can
  * list posts with pagination
  * get post by id 
  * delete post by id 
  * delete all posts 