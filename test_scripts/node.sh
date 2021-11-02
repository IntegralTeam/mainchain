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


start_main_node() {
    $BIN start --home $HOME_PATH_MAIN \
    --rpc.laddr tcp://$RPC_ADDR_MAIN \
    --grpc.address $GRPC_ADDR_MAIN \
    --address tcp://$LISTEN_ADDR_MAIN \
    --p2p.laddr tcp://$P2P_ADDR_MAIN 
}

start_client_node() {
    $BIN start --home $HOME_PATH_CLIENT \
    --rpc.laddr tcp://$RPC_ADDR_CLIENT \
    --grpc.address $GRPC_ADDR_CLIENT \
    --address tcp://$LISTEN_ADDR_CLIENT \
    --p2p.laddr tcp://$P2P_ADDR_CLIENT
}

init_main_node() {
    $BIN init $MONIKER_MAIN $CHAIN_ID --home $HOME_PATH_MAIN
    $BIN keys add $KEY_NAME_VALIDATOR_MAIN --home $HOME_PATH_MAIN --output json > $HOME_PATH_MAIN/validator-phrase.json
    $BIN keys add $KEY_NAME_ORCHESTRATOR_MAIN --home $HOME_PATH_MAIN --output json > $HOME_PATH_MAIN/orchestrator-phrase.json
    VALIDATOR_ADDR=$( $BIN keys show $KEY_NAME_VALIDATOR_MAIN -a )
    ORCHESTRATOR_ADDR=$( $BIN keys show $KEY_NAME_ORCHESTRATOR_MAIN -a )
    $BIN add-genesis-account $VALIDATOR_ADDR $ALLOCATION_VALIDATOR_MAIN --home $HOME_PATH_MAIN
    $BIN add-genesis-account $ORCHESTRATOR_ADDR $ALLOCATION_ORCHESTRATOR_MAIN --home $HOME_PATH_MAIN
    # TODO
    tmp=$(mktemp)
    jq --arg id "$GRAVITY_ID" '.app_state.gravity.params.gravity_id = $id' $HOME_PATH_MAIN/config/genesis.json  > "$tmp" && mv "$tmp" $HOME_PATH_MAIN/config/genesis.json 
    # Enable api in main node
    $BIN gentx $KEY_NAME_VALIDATOR_MAIN $AMMOUNT_VALIDATOR_MAIN $ETH_ADDR_MAIN $ORCHESTRATOR_ADDR --home $HOME_PATH_MAIN $CHAIN_ID
    $BIN collect-gentxs --home $HOME_PATH_MAIN
    $BIN start --home $HOME_PATH_MAIN \
    --rpc.laddr tcp://$RPC_ADDR_MAIN \
    --grpc.address $GRPC_ADDR_MAIN \
    --address tcp://$LISTEN_ADDR_MAIN \
    --p2p.laddr tcp://$P2P_ADDR_MAIN 
}

init_client_node() {
    $BIN init $MONIKER_CLIENT --home $HOME_PATH_CLIENT $CHAIN_ID 

    ID_SEED=$(curl -X get http://$RPC_ADDR_MAIN/status? | jq --raw-output '.result.node_info.id')
    P2P_SEEDS="$ID_SEED@$P2P_ADDR_MAIN"

    curl -s http://$RPC_ADDR_MAIN/genesis | jq -r .result.genesis > $HOME_PATH_CLIENT/config/genesis.json
    sed -i 's/seeds =.*/seeds = "'$P2P_SEEDS'"/g' $HOME_PATH_CLIENT/config/config.toml
    $BIN keys add $KEY_NAME_VALIDATOR_CLIENT --home $HOME_PATH_CLIENT --output json > $HOME_PATH_CLIENT/validator-phrase.json
    $BIN keys add $KEY_NAME_ORCHESTRATOR_CLIENT --home $HOME_PATH_CLIENT --output json > $HOME_PATH_CLIENT/orchestrator-phrase.json
    
    SENDER=$($BIN keys show $KEY_NAME_VALIDATOR_MAIN -a )
    VALIDATOR_ADDR=$($BIN keys show $KEY_NAME_VALIDATOR_CLIENT -a )
    ORCHESTRATOR_ADDR=$($BIN keys show $KEY_NAME_ORCHESTRATOR_CLIENT -a )
    yes | $BIN tx bank send $SENDER $VALIDATOR_ADDR $AMMOUNT_SEND --from $KEY_NAME_VALIDATOR_MAIN $CHAIN_ID --node http://$RPC_ADDR_MAIN
    yes | $BIN tx bank send $SENDER $ORCHESTRATOR_ADDR $AMMOUNT_SEND --from $KEY_NAME_VALIDATOR_MAIN $CHAIN_ID --node http://$RPC_ADDR_MAIN
    yes | $BIN tx staking create-validator \
        --amount=$AMMOUNT_VALIDATOR_CLIENT \
        --pubkey=$($BIN tendermint show-validator --home $HOME_PATH_CLIENT) \
        --chain-id=und \
        --commission-rate="0.10" \
        --commission-max-rate="0.20" \
        --commission-max-change-rate="0.01" \
        --min-self-delegation="100000" \
        --gas="2000000" \
        --gas-prices="0.025stake" \
        --from=$KEY_NAME_VALIDATOR_CLIENT \
        --node http://$RPC_ADDR_MAIN

    VALOPER=$($BIN keys show $VALIDATOR_ADDR --bech val --output json | jq -r '.address')
    yes | $BIN tx staking delegate $VALOPER $AMMOUNT_DELEGATE --from $KEY_NAME_VALIDATOR_CLIENT $CHAIN_ID --node http://$RPC_ADDR_MAIN

    $BIN start --home $HOME_PATH_CLIENT \
    --rpc.laddr tcp://$RPC_ADDR_CLIENT \
    --grpc.address $GRPC_ADDR_CLIENT \
    --address tcp://$LISTEN_ADDR_CLIENT \
    --p2p.laddr tcp://$P2P_ADDR_CLIENT
}


if [ $NODE = 1 ]; then
    
    if [ -d "$HOME_PATH_MAIN" ]; then
        start_main_node
    else
        init_main_node
    fi
else 
    if [ -d "$HOME_PATH_CLIENT" ]; then
        start_client_node
    else
        init_client_node
    fi
fi






