(window.webpackJsonp=window.webpackJsonp||[]).push([[20],{378:function(e,t,s){"use strict";s.r(t);var a=s(42),i=Object(a.a)({},(function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[s("h1",{attrs:{id:"fees-and-gas"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#fees-and-gas"}},[e._v("#")]),e._v(" Fees and Gas")]),e._v(" "),s("p",[e._v("Transactions consume gas, and the sender must pay a fee in order for the transaction be processed by the validator nodes. The fee is calculated from the amount of gas a Tx will consume multiplied by the gas price.")]),e._v(" "),s("div",{staticClass:"custom-block tip"},[s("p",{staticClass:"custom-block-title"},[e._v("NOTE")]),e._v(" "),s("p",[e._v("The gas price for a transaction is set by the sender of the Tx, but each validator will have set their own "),s("code",[e._v("minimum-gas-prices")]),e._v(" value, and will not process transactions that do not meet this minimum requirement.")])]),e._v(" "),s("p",[e._v("Fees are paid in "),s("code",[e._v("nund")]),e._v(", and may be either set or calculated depending on which flags are passed to the "),s("code",[e._v("undcli")]),e._v(" command.")]),e._v(" "),s("div",{staticClass:"custom-block tip"},[s("p",{staticClass:"custom-block-title"},[e._v("NOTE")]),e._v(" "),s("p",[e._v("only "),s("code",[e._v("--fees")]),e._v(" or "),s("code",[e._v("--gas-prices")]),e._v(" may be used - not both at the same time.")])]),e._v(" "),s("div",{staticClass:"custom-block tip"},[s("p",{staticClass:"custom-block-title"},[e._v("TIP")]),e._v(" "),s("p",[s("code",[e._v("--gas-prices")]),e._v(" can be used along with the "),s("code",[e._v("--gas=auto")]),e._v(" and "),s("code",[e._v("--gas-adjustment")]),e._v(" flags to estimate the gas requirement and automatically calculate the Tx fees.")])]),e._v(" "),s("h2",{attrs:{id:"example-1-setting-fees"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#example-1-setting-fees"}},[e._v("#")]),e._v(" Example 1: setting --fees")]),e._v(" "),s("p",[e._v("In this example, we're manually setting the fee for a standard "),s("code",[e._v("send")]),e._v(" transaction. The validator has a "),s("code",[e._v("minimum-gas-prices")]),e._v(" of "),s("code",[e._v("0.25nund")]),e._v(". We'll set the "),s("code",[e._v("--fee")]),e._v(" to 20000nund. For the purposes of simpler calculations, we'll assume the amount of "),s("strong",[e._v("gas")]),e._v(" consumed for this "),s("code",[e._v("send")]),e._v(" transaction, including a small "),s("code",[e._v("--memo")]),e._v(" will be around 65000.")]),e._v(" "),s("div",{staticClass:"custom-block tip"},[s("p",{staticClass:"custom-block-title"},[e._v("NOTE")]),e._v(" "),s("p",[e._v("gas is defined on the chain as a flat cost per byte for a Tx, e.g. 10 gas per byte. The total size of our Tx will be around 6500 bytes, and therefore the gas consumed by the Tx will be 6500 * 10 = 65000.")])]),e._v(" "),s("div",{staticClass:"language- extra-class"},[s("pre",{pre:!0,attrs:{class:"language-text"}},[s("code",[e._v('undcli tx send [from] [to] 123456nund --memo="some und from me to you" --fees=20000nund\n')])])]),s("p",[e._v("In this instance, the "),s("code",[e._v("gas-price")]),e._v(" is implied as approximately 0.31nund (fee / gas: 20000 / 65000), so the Validator will accept the Tx and include it in the block, since 0.31nund > 0.25nund.")]),e._v(" "),s("p",[e._v("If we had set the "),s("code",[e._v("--fees")]),e._v(" to 10000, it would not have been processed by the Validator (10000 / 65000 = 0.15nund).")]),e._v(" "),s("div",{staticClass:"custom-block tip"},[s("p",{staticClass:"custom-block-title"},[e._v("NOTE")]),e._v(" "),s("p",[e._v("the Tx with lower fees may remain the Tx pool until a validator with lower "),s("code",[e._v("minimum-gas-prices")]),e._v(" picks it up and proposes the block.")])]),e._v(" "),s("h2",{attrs:{id:"example-2-setting-gas-and-gas-prices"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#example-2-setting-gas-and-gas-prices"}},[e._v("#")]),e._v(" Example 2: setting --gas and --gas-prices")]),e._v(" "),s("p",[e._v("In this example, we'll set our own "),s("code",[e._v("--gas-prices")]),e._v(", and ask "),s("code",[e._v("undcli")]),e._v(" to estimate the amount of gas the Tx will consume based on the Tx input by passing the "),s("code",[e._v("--gas=auto")]),e._v(" flag. We can also use the "),s("code",[e._v("--gas-adjustment")]),e._v(" flag to increase/decrease this gas estimate. We'll assume again that the calculated estimate will be around 65000 gas:")]),e._v(" "),s("div",{staticClass:"language- extra-class"},[s("pre",{pre:!0,attrs:{class:"language-text"}},[s("code",[e._v('undcli tx send from to 123456nund --memo="some und from me to you" --gas=auto --gas-prices=0.25nund\n')])])]),s("p",[e._v("In this example, the Tx "),s("strong",[e._v("fee")]),e._v(" will be calculated and included in the transaction for us. The fee will be around 16250nund ("),s("code",[e._v("gas * gas-prices")]),e._v(": 65000 * 0.25). Since we have set "),s("code",[e._v("gas-prices")]),e._v(" already to 0.25 (and assuming the gas estimate is also correct), this Tx will be processed by the validator.")]),e._v(" "),s("div",{staticClass:"custom-block tip"},[s("p",{staticClass:"custom-block-title"},[e._v("NOTE")]),e._v(" "),s("p",[e._v("Adding the "),s("code",[e._v("--gas-adjustment")]),e._v(" flag, for example "),s("code",[e._v("--gas-adjustment=1.5")]),e._v(", will increase the gas estimate and therefore the fee, but will increase the chances of the Tx being processed.")])]),e._v(" "),s("p",[e._v("Validators will prioritise Txs with higher "),s("code",[e._v("gas-prices")]),e._v(", so it is worth setting higher prices to ensure your Tx is processed.")])])}),[],!1,null,null,null);t.default=i.exports}}]);