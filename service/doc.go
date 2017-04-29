// Package service contains functions of Service, including service discovery and load balancing
package service

/*
Service is designed for connecting all kinds of concepts, including
Authenticator, Throttler and Router.

Events:

1. register - name, labels - can not be modified
2. remove - name, labels
3. bind - upstream to service
4. register - authenticator
5. add authenticator to service
6. register - throttler
7. add throttler to service
*/
