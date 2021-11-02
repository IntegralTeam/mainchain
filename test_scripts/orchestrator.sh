#!/bin/bash

set -a
source .env
set +a 


# what node is starting?
# 1 - main
# 2 - client
if [ -n "$1" ]; then
    NODE=$1
else 
    echo "set what node is starting: 1 - main, 2 - client"
    exit
fi

init_orchestrator() {
    $BIN_ORCHESTRATOR --home $1 init
    $BIN_ORCHESTRATOR --home $1 -a $ADDR_PREFIX keys register-orchestrator-address \
        --validator-phrase "$2" \
        --cosmos-phrase "$3" \
        --fees $FEES \
        --ethereum-key $4 \
        --cosmos-grpc http://$5
    $BIN_ORCHESTRATOR --home $1 -a $ADDR_PREFIX orchestrator -f $FEES $ETH_RPC --cosmos-grpc http://$5 -g $CONTRACT_ADDR
}

start_orchestrator() {
    $BIN_ORCHESTRATOR --home $1 -a $ADDR_PREFIX orchestrator -f $FEES $ETH_RPC --cosmos-grpc http://$2 -g $CONTRACT_ADDR
}

if [ $NODE = 1 ]; then

    HOME_PATH=$HOME_PATH_MAIN/orchestrator

    if [ -d $HOME_PATH ]; then

        start_orchestrator $HOME_PATH $GRPC_ADDR_MAIN

    else

        VALIDATOR_PHRASE=$(cat $HOME_PATH_MAIN/validator-phrase.json | jq --raw-output '.mnemonic')
        ORCHESTRATOR_PHRASE=$(cat $HOME_PATH_MAIN/orchestrator-phrase.json | jq --raw-output '.mnemonic')
        
        init_orchestrator $HOME_PATH "$VALIDATOR_PHRASE" "$ORCHESTRATOR_PHRASE" $ETH_PRIVATE_KEY_MAIN $GRPC_ADDR_MAIN

    fi
else

    HOME_PATH=$HOME_PATH_CLIENT/orchestrator

    if [ -d $HOME_PATH ]; then

        start_orchestrator $HOME_PATH $GRPC_ADDR_CLIENT

    else

        VALIDATOR_PHRASE=$(cat $HOME_PATH_CLIENT/validator-phrase.json | jq --raw-output '.mnemonic')
        ORCHESTRATOR_PHRASE=$(cat $HOME_PATH_CLIENT/orchestrator-phrase.json | jq --raw-output '.mnemonic')
        
        init_orchestrator $HOME_PATH "$VALIDATOR_PHRASE" "$ORCHESTRATOR_PHRASE" $ETH_PRIVATE_KEY_CLIENT $GRPC_ADDR_CLIENT

    fi
fi