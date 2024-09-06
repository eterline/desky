#!/bin/sh

token=""
chat_id=""

curl -s -X POST https://api.telegram.org/bot$token/sendMessage \
        -d chat_id=$chat_id \
        -d text="Desky:%0A-----------------%0A$(date)%0AAuthorization occurred"
exit