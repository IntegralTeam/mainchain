(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{367:function(e,t,a){"use strict";a.r(t);var s=a(42),n=Object(s.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"beacon-example"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#beacon-example"}},[e._v("#")]),e._v(" BEACON Example")]),e._v(" "),a("h4",{attrs:{id:"contents"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#contents"}},[e._v("#")]),e._v(" Contents")]),e._v(" "),a("p"),a("div",{staticClass:"table-of-contents"},[a("ul")]),a("p"),e._v(" "),a("p",[e._v('Register:\nundcli tx beacon register --moniker=beacon1 --name="Beacon 1" --from wrktest')]),e._v(" "),a("p",[e._v("then run:")]),e._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("undcli query tx [TX HASH]\n")])])]),a("p",[e._v("this will return the generated Beacon ID integer")]),e._v(" "),a("p",[e._v("Query metadata")]),e._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("undcli query beacon beacon 1\n")])])]),a("p",[e._v("Record Timestamp hash")]),e._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("undcli tx beacon record 1 --hash=d04b98f48e8 --subtime=$(date +%s) --from wrktest --gas=auto --gas-adjustment=1.15\n")])])]),a("p",[e._v("Query a Timestamp")]),e._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("undcli query beacon timestamp 1 1\n")])])])])}),[],!1,null,null,null);t.default=n.exports}}]);