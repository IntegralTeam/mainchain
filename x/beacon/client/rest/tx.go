package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"github.com/unification-com/mainchain-cosmos/x/beacon/internal/types"
)

type registerBeaconReq struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Moniker    string       `json:"moniker"`
	BeaconName string       `json:"name"`
	Owner      string       `json:"owner"`
}

type recordBeaconTimestampReq struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	BeaconID   uint64       `json:"id"`
	SubmitTime uint64       `json:"subtime"`
	Hash       string       `json:"hash"`
	Owner      string       `json:"owner"`
}

// registerTxRoutes - define REST Tx routes
func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/beacon/reg"), registerBeaconHandler(cliCtx)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/beacon/rec"), recordBeaconTimestampHandler(cliCtx)).Methods("POST")
}

func registerBeaconHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req registerBeaconReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgRegisterBeacon(req.Moniker, req.BeaconName, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func recordBeaconTimestampHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req recordBeaconTimestampReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgRecordBeaconTimestamp(req.BeaconID, req.Hash, req.SubmitTime, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
