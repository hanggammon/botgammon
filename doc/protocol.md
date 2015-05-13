Protocol Specification
======================

REST API
--------
The following API has a game ID parameter which is intended to support multiple concurrent games, but initially we'll be using id=0 for all calls.

    /game/<id>/new
Resets the existing game in progress (all clients are notified of the reset event) and starts a new game with clean game state. Existing players and sockets are retained.

    /game/<id>/addplayer?name=<playername>
Adds the player 'playername' to the game. On success HTTP 200 is returned. If there are no free spots left in the game then HTTP 418 (I'm a teapot) is returned and the error message is sent as plaintext in the HTTP response.

WebSocket Protocol
------------------
