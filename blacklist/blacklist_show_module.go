package blacklist

import (
	"fmt"
	"github.com/khsadira/cleaner/auth"
	"net/http"
)

func BlacklistShowModule(w http.ResponseWriter, r *http.Request) {
	if auth.TryConnection(w, r) {

		var p, c, cg string

		for _, promo := range bList.Promotions {
			p += promo.Id + "<br />"
		}
		for _, camp := range bList.Campaigns {
			c += camp.Id + "<br />"
		}
		for _, cstmgrp := range bList.CustGrps {
			cg += cstmgrp.Id + "<br />"
		}
		ret := fmt.Sprintf(`<html>
		<body>
			<h1>Promotions</h1>
			<p>%s</p>
			<h1>Campaigns</h1>
			<p>%s</p>
			<h1>Customer Groups</h1>
			<p>%s</p>
			<p><br /><a href='/blacklist/'>done</a></p>
		</body>

	</html>`, p, c, cg)

		w.Write([]byte(ret))
	}
}
