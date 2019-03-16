Request:

POST /invite/send
Content-Type: application/json
{
    "event": {
        "name": "Karuppiah's Birthday Party!",
        "description": "It's going to be held on 1st Jan 2020! Let's party!",
        "email": "alex@gmail.com"
    }
}

Response:

200 OK
Content-Type: application/json
{
    "message": "mail will be sent"
}