# SaveService
*A service designed to host game data, states ETC*

> **WARNING**
>
> This is a work in progress, and as a result there is **NO** database or other persistent storage soloution

SaveService is a service designed to allow users to store and load a *state* from a server. This state is a struct which is stored as a generic parameter, allowing your state to be most simple data types.

> I made this for a larger project, it is attended to accompany a game which i will eventually make in TS using webgl bindings and a custom engine.

# Example code

Can be found at 

    ./example/main.go

# API

This service has several routes. 

## `/new`
*Allows the end user to create a new user*

This does not need any particular format.

This set a cookie called user on the client's browser. This cookie is required to perform `/save` and `/load`.


## `/save`
*Allows the end user to save data to their user object, overwriting whatever was originally saved*

This request **requires** the user cookie to be set via `/new`.

This request **requires** the request to have a body matching the json representation of the GameState_T struct.

## `/load`
*Allows the end user to read data saved on their user object.*

This request **requires** the user cookie to be set via `/new`.

This request returns the json representation of the GameState_T struct owned by the user.

## `/delete`
*Allows the end user to delete data saved under themselves.*

This request **requires** the user cookie to be set via `/new`.
This request deletes the user entry AND the user from the service. After this, the user cookie will have no use, and will return 403 Forbidden when used.

