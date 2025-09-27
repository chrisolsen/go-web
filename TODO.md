# TODOS

- [ ] Determine why I only have a Caddyfile.dev (for dev)
- Create a landing page and a route
- Create a login page and a forgot your password
- Get Sqlite setup
- Create auth service with required migrations
  - Create session table with list of tokens that link to the user id
    - tokens will have an expiry timestamp
    - session service will need a `clean` function to clean out the expires tokens
    - Attributes
      - token
      - userId
      - expiresOn
  - Create a function to check if the user is authorized or not

## From Book
 - [ ] TemplateCache pg165
 - [ ] Router (Alice) pg192
 - [ ] Validation pg217
 - [ ] FormDecoder pg224
 - [ ] Cron-like lib github.com/robfig/cron
     - remove reliance on Caddy
     - user own pem files pg 253
 - [ ] Panic Recovery
    - panics won't bring down the server, but they may result in a false 200 OK and incomplete render


