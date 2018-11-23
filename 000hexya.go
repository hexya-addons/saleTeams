// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package saleTeams

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-addons/web/controllers"
	_ "github.com/hexya-addons/webKanban"

	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME = "saleTeams"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PostInit: func() {},
	})

	GroupSaleSalesman = security.Registry.NewGroup("sale_teams_group_sale_salesman", "User: Own Documents Only", base.GroupUser)
	GroupSaleSalesmanAllLeads = security.Registry.NewGroup("sale_teams_group_sale_salesman_all_leads", "User: All Documents", GroupSaleSalesman)
	GroupSaleManager = security.Registry.NewGroup("sale_teams_group_sale_manager", "Manager", GroupSaleSalesmanAllLeads)

	controllers.BackendLess = append(controllers.BackendLess, "/static/saleTeams/src/less/sales_team_dashboard.less")
	controllers.BackendCSS = append(controllers.BackendCSS, "/static/saleTeams/src/css/sales_team.css")
	controllers.BackendJS = append(controllers.BackendJS,
		"/static/saleTeams/src/js/sales_team.js",
		"/static/saleTeams/src/js/sales_team_dashboard.js",
	)
}
