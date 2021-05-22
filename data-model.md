# Model


## Room

* id
* type              string/enum {trun,...}
* players           Player[]
    * id
    * nickname
    * username
    * role (guest, owner, ...)
    * accessToken
    * refreshToken
* max_player {int}
* request_for_join  Player[]
* group             Group[]
    * id
    * title
    * max_member {int}
    * owner
    * members       Player[]
* policy            Policy[]
    * id
    * title
    * value
* action            Action[]
    * id
    * title (send card to deck)
    * status
* max_action {int}
* envirnoment       Envirnoment[]
    * id
    * title (deck one)
    * status
* turn_history      Act[]
    * id
    * action (external id)
    * group  (external id)
    * policy (external id)
    * envirnoment (external id)
* turn_current      Act[]
    * id
    * action (external id)
    * group  (external id)
    * policy (external id)
    * envirnoment (external id)
* blocked       Player[]
* winner        Player[]
* loser         Player[]
* date          Date{start, end}
* status
