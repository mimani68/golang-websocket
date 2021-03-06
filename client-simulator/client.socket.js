/**
 *
 * main method of
 *
 */
function start() {
    var roomId = null
    var otherPlayerId = null
    var otherPlayerSession = null
    var gameId = $('#namespace').val()
    var jwt = $('#jwt').val()
    var BASE_URL = $('#serverAddress').val()
    var socket = io(BASE_URL, {
        path: '/balout/api/v1/match/', 
        transports: ['websocket'],
        // query: {
        //     'app': gameId,
        //     'token': jwt
        // },
        // extraHeaders: {
        //     'x-api-token': gameId,
        //     'X-test-Version': 'balout.simulator/v1.1.1',
        //     'X-authentication-token': jwt
        // },
        // transportOptions: {
        //     polling: {
        //         extraHeaders: {
        //             'x-api-token': gameId,
        //             'X-test-Version': 'balout.simulator/v1.1.1',
        //             'X-authentication-token': jwt
        //         }
        //     }
        // }
    })

    /************************************
     *
     *           E V E N T S
     *
     ************************************/
    socket.on('event', function (data) {})
    socket.on('disconnect', function () {
        console.log('=========| end connection |========')
    })
    socket.on('connect', function () {
        $('#sessionPlaceHolder').html(`<h5>${ socket.id }</h5>`)
        console.log('=========| start connection |========')
    })


    socket.on('balout:player:invalid-token', function (msg) {
        console.log(msg)
    })
    socket.on('balout:player:valid-token', function (msg) {
        console.log(msg)
    })
    socket.on('balout:player:identity', function (msg) {
        console.log(msg)
    })


    socket.on('balout:match:start', function (msg) {
        roomId = msg.body.room
        otherPlayerId = msg.body.opponent.id
        console.log(msg)
    })
    socket.on('balout:match:waiting', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:progress', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:finish', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:error', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:alert:is-same', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:you-are-disconnected', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:cheat', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:disconnect-other-player', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:join-again-other-player', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:player:act', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:player:act:retry', function (msg) {
        console.log(msg)
    })
    socket.on('balout:match:player:leave', function (msg) {
        console.log(msg)
    })


    socket.on('balout:system:ping', function (msg) {
        console.log(msg)
    })
    socket.on('balout:system:error', function (msg) {
        console.log(msg)
    })


    socket.on('balout:chat:send:ack', function (msg) {
        console.log(msg)
    })
    socket.on('balout:chat:inbox:latest', function (msg) {
        otherPlayerSession = msg.body.senderSession
        console.log(msg)
    })


    socket.on('balout:dev', function (msg) {
        console.log(msg)
    })

    /************************************
     *
     *           E V E N T S
     *
     ************************************/

    $('#new-match')
    .on('click', function () {
        socket.emit('balout:match:player:new-match', {},
        msg => {
            console.log(msg)
            roomId = msg.body.room
        })
    })
    $('#ready')
        .on('click', function () {
            socket.emit('balout:match:player:ready', { 
                body: {
                    room: roomId
                }
            },
            msg => {
                console.log(msg)
            })
        })

    $('#play')
        .on('click', function () {
            socket.emit('balout:match:player:act', {
                body: {
                    room: roomId,
                    newSerial: new Date().toISOString()
                }
            },
            msg => {
                console.log(msg)
            })
        })

    $('#retrySendAnswer')
        .on('click', function () {
            socket.emit('balout:match:player:act:retry', {
                body: {
                    room: roomId,
                    oldSerial: 'f4589',
                    newSerial: new Date().toISOString(),
                    value: '??????'
                }
            },
            msg => {
                console.log(msg)
            })
        })

    $('#buyCheat')
        .on('click', function () {
            socket.emit('balout:match:cheat', {
                body: {
                    room: roomId
                }
            },
            msg => {
                console.log(msg)
            })
        })

    $('#leave')
        .on('click', function () {
            socket.emit('balout:match:player:leave', {
                body: {
                    room: roomId
                }
            })
        },
        msg => {
            console.log(msg)
        })



    $('#ping')
        .on('click', function () {
            socket.emit('balout:system:ping', {
                body: {
                    text: 'salam from client'
                }
            });
        })

    $('#whoami')
        .on('click', function () {
            socket.emit('balout:player:identity', {})
        })
    $('#auth')
        .on('click', function () {
            socket.emit('balout:player:authenticate', {
                body: {
                    token: 'jwt BgSmtlGmRvbbbmbwIlTDaiVY...'
                }
            }, msg=>{
                console.log(msg)
            })
        })
        
    $('#dev')
        .on('click', function () {
            socket.emit('balout:dev', JSON.stringify({
                room: roomId
            }))
        })

    $('#cleanDb')
        .on('click', function () {
            socket.emit('balout:dev:clean-db', {})
        })



    $('#sendMessage')
        .on('click', function () {
            var id = $('#audiencePlayerId').val()
            socket.emit('balout:chat:send', {
                player: id,
                message: ' salam'
            })
        })

}