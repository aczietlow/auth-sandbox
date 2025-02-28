# Keycloak setup 

As I'm building out an oauth client, I'll need a OIDC Authorization Server. For that I'll be utilizing Keycloak, as it's a great open source IAM solutuion.

## Setup

### Create a dummy user

u: user
p: pass

### Create a new client within KC

1. Create
- Nav to Clients -> Create Client
- Client Type will be "OIDC"
- client authentication "on"
  - Standard flow: on (This is auth flow)
  - direct access grant: off (This is ROPC)
- settings/access settings
  - valid urls: http://localhost:1322/oauth2/callback
- grab client secret (under the credentials tab) 

2. Configure mapper for client


## Interesting but for later

KC supports well-known

`/realms/{realm-name}/.well-known/openid-configuration`
