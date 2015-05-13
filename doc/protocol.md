Protocol Specification
======================

REST API
--------

    Request:       POST /api/v1/games
    Response:      HTTP 200
    Response Body: { "id" : "0" }

Creates a new game, JSON response contains the ID of the newly created game.


    Request:       PATCH /api/v1/games/<id>/reset
    Response:      HTTP 200

Resets the game at <id>. Clients are notified through a WebSocket protocol
message of the reset event and all sockets/players are dropped from the game.

    Request:       POST /api/v1/games/<id>/players
    Request Body:  { "name" : "mark" }
    Response:      HTTP 200
    Response Body: { "websocket" : "XXXX" }

Response bodies are in JSON format. In addition to the indicated fields the
following fields are present for indicating API success or failure:

On success:

    { "status" : "OK" }

On failure:

    { "status" : "FAILED",
      "error" : "Maximum number of players reached for game." }

WebSocket Protocol
------------------
