Protocol Specification
======================

REST API
--------

### Create Game
    Request:       POST /api/v1/games
    Response:      HTTP 200
    Response Body: { "id" : "0" }

Creates a new game, JSON response contains the ID of the newly created game.


### Reset Game
    Request:       PATCH /api/v1/games/<id>/reset
    Response:      HTTP 200

Resets the game at id. Clients are notified through a Socket.IO protocol
message of the reset event and all sockets/players are dropped from the game.

### Add Player
    Request        : POST /api/v1/games/<id>/players
    Request Body   : { "name" : "mark" }
    Response       : HTTP 200
    Response Body  : { "websocket" : "<socket URI>",
                       "seat" : 1 }

    (optional)
    Request        : POST /api/v1/games/<id>/players
    Request Body   : { "name" : "baran",
                       "seat" : 1 }
    Response       : HTTP 200
    Response Body  : { "websocket" : "<socket URI>",
                       "seat" : 1 }

Adds a player to the game at id. The player name is passed in as a request
field through JSON. The response JSON contains a socket URI for the Socket.IO
client/server connection.

### Status Reporting
Response bodies are in JSON format. In addition to the indicated fields the
following fields are present for indicating API success or failure:

On success:

    { "status" : "OK" }

On failure:

    { "status" : "FAILED",
      "error" : "Maximum number of players reached for game." }

    { "status" : "FAILED",
      "error" : "Seat X is occupied." }


Socket.IO Protocol
------------------

### Board State JSON representation

The regular board slots are represented as an array with 24 elements. Each slot
element has a sub-array that contains an array. Pieces are represented as
integers denoting the team (0 for white, 1 for black). The collected and hit
pieces are represented by an array each, with the same semantics as a regular
slot.

    { "slots":
             [
               [0, 0],
               [],
               [],
               [],
               [],
               [1, 1, 1, 1, 1],
               [],
               [1, 1, 1],
               [],
               [],
               [],
               [0, 0, 0],
               [1, 1, 1],
               [],
               [],
               [],
               [0, 0, 0],
               [],
               [0, 0, 0, 0, 0],
               [],
               [],
               [],
               [],
               [1, 1]
             ],
     "hit" : [0, 1],
     "collected" : [1, 0]
   }

### Move Request

This is the JSON request sent by the server to a client to request a move
response. The indices are flipped around by the server so the client receiving
the request is always player 0, their pieces are represented by team 0, they are
on boards[0] and the board is flipped so their team is moving in the positive
direction (0->23). It is implicit that players[0] is playing against players[1]
and players[2] is playing against players[3].

    { "players": [ "baran", "mrj", "mrpdaem0n", "bpoess" ],
      "diceroll": [6, 2],
      "boards": [ {...}, {...}]
    }
