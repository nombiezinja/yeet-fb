# yeet-fb

YEET FB AAAAAA

## TODOs
CLI only
* investigate fb developer graphql endpoints/ new developer best practices aaaaa
* rough repo structure
* generate server
* db set up 
* redis set up
* alk;djflkasjdfl; 
* use short lived access token for development
* get 1 post works
* delete 1 post works 
* a;ldkfjal;sdkjf;alskd aaaaaa
* get user authn working
* get user authz working
* as;dlkfjasdlkfjlk;j aaaaa
add frontend aaaaa
deploy aaaaa
asl;dkjfas;dlkfj

## brain dump 
* tool choices 
  * generate go server 
  * chi
  * flag to start, add cobra later
  * redis 
  * psql 
  * testify/mock 
  * OPA
  * tbd whether to enable github.com/nombiezinja/k2so run individual tasks at scale?? mebe not in prod but would be fun to do for development
  * tbd otel tools (pending review of new stuffs)
  * tbd frontend 
  * tbd ci/cd stuff (github actions?)
  * tbd terraform or something else 
  * tbd queue (prob not needed for a while)
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