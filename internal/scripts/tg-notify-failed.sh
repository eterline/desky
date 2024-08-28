#!/bin/sh

token="6835601638:AAFDe3FDMdrHhvUv9XRLiChbG_6sS_BzxmU"
chat_id="710085729"

curl -s -X POST https://api.telegram.org/bot$token/sendMessage \
        -d chat_id=$chat_id \
        -d text="Desky:%0A-----------------%0A$(date)%0AAuthorization attempt failed"
exit