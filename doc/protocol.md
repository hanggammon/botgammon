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
    Request:       POST /api/v1/games/<id>/players
    Request Body:  { "name" : "mark" }
    Response:      HTTP 200
    Response Body: { "websocket" : "<socket URI>" }

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

Socket.IO Protocol
------------------

### Board State JSON representation

    { "board":
             { "slots":
                      [
                        { "index": "0", "pieces":
                                                [
                                                   { "id": "4", "team": "white"},
                                                   { "id": "5", "team": "white"}
                                                ]
                        },
                        { "index": "1", "pieces": [] },
                        { "index": "2", "pieces": [] },
                        { "index": "3", "pieces": [] },
                        { "index": "4", "pieces": [] },
                        { "index": "5", "pieces":
                                                [
                                                   { "id": "13", "team": "black"},
                                                   { "id": "14", "team": "black"}
                                                ]
                        },

                        ...

                        { "index": "23", "pieces":
                                                 [
                                                    { "id": "29", "team": "black"},
                                                    { "id": "30", "team": "black"}
                                                 ]
                        },
                        { "index": "Hit", "pieces":
                                                  [
                                                     {"id": "11", "team": "white"}
                                                  ]
                        }
                      ]
            }
    }
