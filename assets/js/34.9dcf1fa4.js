(window.webpackJsonp=window.webpackJsonp||[]).push([[34],{392:function(t,e,a){"use strict";a.r(e);var s=a(42),n=Object(s.a)({},(function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"und-mainchain-config-app-toml-reference"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#und-mainchain-config-app-toml-reference"}},[t._v("#")]),t._v(" "),a("code",[t._v(".und_mainchain/config/app.toml")]),t._v(" Reference")]),t._v(" "),a("p",[t._v("The "),a("code",[t._v("$HOME/.und_mainchain/config/app.toml")]),t._v(" file contains all the configuration options for the "),a("code",[t._v("und")]),t._v(" server binary. Below is a reference for the file.")]),t._v(" "),a("h4",{attrs:{id:"contents"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#contents"}},[t._v("#")]),t._v(" Contents")]),t._v(" "),a("p"),a("div",{staticClass:"table-of-contents"},[a("ul",[a("li",[a("a",{attrs:{href:"#main-base-config-options"}},[t._v("Main base config options")]),a("ul",[a("li",[a("a",{attrs:{href:"#minimum-gas-prices"}},[t._v("minimum-gas-prices")])]),a("li",[a("a",{attrs:{href:"#halt-height"}},[t._v("halt-height")])]),a("li",[a("a",{attrs:{href:"#halt-time"}},[t._v("halt-time")])]),a("li",[a("a",{attrs:{href:"#inter-block-cache"}},[t._v("inter-block-cache")])]),a("li",[a("a",{attrs:{href:"#pruning"}},[t._v("pruning")])])])])])]),a("p"),t._v(" "),a("h2",{attrs:{id:"main-base-config-options"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#main-base-config-options"}},[t._v("#")]),t._v(" Main base config options")]),t._v(" "),a("h3",{attrs:{id:"minimum-gas-prices"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#minimum-gas-prices"}},[t._v("#")]),t._v(" minimum-gas-prices")]),t._v(" "),a("p",[t._v("The minimum gas prices a validator is willing to accept for processing a transaction. A transaction's fees must meet the minimum of any denomination specified in this config (e.g. "),a("code",[t._v("0.25nund")]),t._v(").")]),t._v(" "),a("p",[t._v("Example")]),t._v(" "),a("div",{staticClass:"language-toml extra-class"},[a("pre",{pre:!0,attrs:{class:"language-toml"}},[a("code",[a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("minimum-gas-prices")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"0.25nund"')]),t._v("\n")])])]),a("h3",{attrs:{id:"halt-height"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#halt-height"}},[t._v("#")]),t._v(" halt-height")]),t._v(" "),a("p",[t._v("HaltHeight contains a non-zero block height at which a node will gracefully halt and shutdown that can be used to assist upgrades and testing.")]),t._v(" "),a("div",{staticClass:"custom-block tip"},[a("p",{staticClass:"custom-block-title"},[t._v("Note")]),t._v(" "),a("p",[t._v("Commitment of state will be attempted on the corresponding block.")])]),t._v(" "),a("p",[t._v("Example")]),t._v(" "),a("div",{staticClass:"language-toml extra-class"},[a("pre",{pre:!0,attrs:{class:"language-toml"}},[a("code",[a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("halt-height")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("12345")]),t._v("\n")])])]),a("h3",{attrs:{id:"halt-time"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#halt-time"}},[t._v("#")]),t._v(" halt-time")]),t._v(" "),a("p",[t._v("HaltTime contains a non-zero minimum block time (in Unix seconds) at which a node will gracefully halt and shutdown that can be used to assist upgrades and testing.")]),t._v(" "),a("div",{staticClass:"custom-block tip"},[a("p",{staticClass:"custom-block-title"},[t._v("Note")]),t._v(" "),a("p",[t._v("Commitment of state will be attempted on the corresponding block.")])]),t._v(" "),a("p",[t._v("Example")]),t._v(" "),a("div",{staticClass:"language-toml extra-class"},[a("pre",{pre:!0,attrs:{class:"language-toml"}},[a("code",[a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("halt-time")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1587565852")]),t._v("\n")])])]),a("h3",{attrs:{id:"inter-block-cache"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#inter-block-cache"}},[t._v("#")]),t._v(" inter-block-cache")]),t._v(" "),a("p",[t._v("InterBlockCache enables inter-block caching.")]),t._v(" "),a("p",[t._v("Example")]),t._v(" "),a("div",{staticClass:"language-toml extra-class"},[a("pre",{pre:!0,attrs:{class:"language-toml"}},[a("code",[a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("inter-block-cache")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("true")]),t._v("\n")])])]),a("h3",{attrs:{id:"pruning"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#pruning"}},[t._v("#")]),t._v(" pruning")]),t._v(" "),a("p",[t._v("Pruning sets the pruning strategy: "),a("code",[t._v("syncable")]),t._v(", "),a("code",[t._v("nothing")]),t._v(", "),a("code",[t._v("everything")])]),t._v(" "),a("div",{staticClass:"custom-block danger"},[a("p",{staticClass:"custom-block-title"},[t._v("IMPORTANT")]),t._v(" "),a("p",[t._v("There is a known issue with the "),a("code",[t._v("syncable")]),t._v(" pruning option in the Cosmos SDK. Since "),a("code",[t._v('pruning = "syncable"')]),t._v(" is the default value when "),a("code",[t._v("und init")]),t._v(" is run, it is recommended to set the value to either "),a("code",[t._v('pruning = "everything"')]),t._v(" or "),a("code",[t._v('pruning = "nothing"')]),t._v(". Note that setting to "),a("code",[t._v('pruning = "nothing"')]),t._v(" will increase storage usage considerably.")])]),t._v(" "),a("ul",[a("li",[a("code",[t._v("syncable")]),t._v(": only those states not needed for state syncing will be deleted (keeps last 100 + every 10000th)")]),t._v(" "),a("li",[a("code",[t._v("nothing")]),t._v(": all historic states will be saved, nothing will be deleted (i.e. archiving node)")]),t._v(" "),a("li",[a("code",[t._v("everything")]),t._v(": all saved states will be deleted, storing only the current state")])]),t._v(" "),a("p",[t._v("Example")]),t._v(" "),a("div",{staticClass:"language-toml extra-class"},[a("pre",{pre:!0,attrs:{class:"language-toml"}},[a("code",[a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("pruning")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"syncable"')]),t._v("\n")])])])])}),[],!1,null,null,null);e.default=n.exports}}]);