# Backend checkers

## The main goal
In this project I will try to write a restful server with the authorization database and the queries described below,
it will be server to store data

## Available http requests
    POST   /api/user        - create new user 
    POST   /api/session     - create new access token for user 
    
    POST   /api/game/create - create new game
    POST   /api/game/move   - make move
    
    GET    /api/game        - get game field
    POST   /api/game        - log in game
    PUT    /api/game        - change settings
    DELETE /api/game        - delete game
    
    
## Steps
1. Handle POST   /api/user (implemented)
2. Handle POST   /api/session (implemented)
3. Handle POST   /api/game/create (implemented)
4. Handle POST   /api/game/move (implemented)
5. Handle GET    /api/game (implemented)
6. Handle POST   /api/game (implemented)
7. Handle PUT    /api/game 
8. Handle DELETE /api/game 
9. add db

## Details of http requests
### /api/user
#### POST
    request:
      Cookie: not required
      body: {
        "username":your_username, 
        "password":your_password
      }
    response:
      201 - success creted user
      400 - bad request
      403 - permission denied, user already exist, even password is right
### /api/session
#### POST
    request:
      Cookie: not required
      body: {
        "username":your_username, 
        "password":your_password,
        "max_age":cookies_lifetime
      }
    response:
      201 - success log in
        Set-Cookies: token=your_access_token
      400 - bad request
      403 - permission denied, username or password incorrect
### /api/game/create
#### POST
    request:
      Cookie: token=your_access_token
      body: {
        "gamename":gamename, 
        "password":game_password,
        "settings": {
          "gamer0":int (0 - man, 1 - bot),
          "gamer1":int (0 - man, 1 - bot),
          "level0":int (bot level from 0 - 3),
          "level1":int
        }
      }
    response:
      201 - success create game
      400 - bad request
      403 - permission denied, game already exist
      500 - something went wrong :-(
### /api/game/move
#### POST
    request:
      Cookie: token=your_access_token
      body: {
        "gamename":gamename, 
        "from": {"x":int,"y":int},
        "to": [{"x":int,"y":int}, ...]
      }
    response:
      201 - successfully moved
      400 - bad request
      403 - permission denied
      404 - game not found
      405 - incorrect move
      500 - something went wrong :-(
### /api/game
#### GET
    request:
      Cookie: token=your_access_token
      body: {
        "gamename":gamename
      }
    response:
      200 - successfully moved
      body: {
        "figures": [
          {
            "x":int, 
            "y":int, 
            "figure":string ("checker", "king"), 
            "gamer_id":int
          }
        ],
        "turnGamerId": int,
        "winner": int (-1 if no winner, 0 if winner 0, 1 <-> 1)
      }
      400 - bad request
      401 - not authorized
      403 - permission denied
      404 - game not found
      500 - something went wrong :-(
#### POST
    request:
      Cookie: token=your_access_token
      body: {
        "gamename":gamename,
        "password":password
      }
    response:
      201 - successfully loged in
      400 - bad request
      401 - not authorized
      403 - permission denied
      404 - game not found
      500 - something went wrong :-(
    

