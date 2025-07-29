# TODOS

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
