import * as React from "react";
import DashboardScreen from "../../dashboard/DashboardScreen";
import Login from "../../login/Login";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import ChargeTokenScreen from "../../chargetokens/ChargeTokenScreen";
import { JsxElement } from "typescript";
import StationsScreen from "../../stations/StationsScreen";
import dashboardMenu from "../../dashboard/constants/dashboardMenu";
import appTheme from "../theme/AppTheme";
import { withSelectedMenuItemContext } from "../contexts/SelectedMenuItem";
import TransactionsScreen from "../../transactions/TransactionsScreen";
import { withRouteContext } from "../contexts/Route";
import { withFirebaseDataContext } from "../contexts/FirebaseData";

const appRoutes = {
  "/": {
    label: "Login page",
    component: withFirebaseDataContext(Login, {}),
    subRoutes: {},
  },
  login: {
    label: "Login page",
    component: withFirebaseDataContext(Login, {}),
    subRoutes: {},
  },
  dashboard: {
    label: "Dashboard",
    component: withSelectedMenuItemContext(DashboardScreen, {
      menu: dashboardMenu,
      theme: appTheme,
    }),
    subRoutes: {
      stations: {
        label: "Charging station list",
        component: React.Fragment,
        subRoutes: {
          list: {
            label: "Charging station list",
            component: withFirebaseDataContext(StationsScreen, {}),
            subRoutes: {},
          },
        },
      },
      chargetokens: {
        label: "RFIDs",
        component: React.Fragment,
        subRoutes: {
          list: {
            label: "Token list",
            component: withFirebaseDataContext(ChargeTokenScreen, {}),
            subRoutes: {},
          },
        },
      },
      transactions: {
        label: "Transactions",
        component: withFirebaseDataContext(TransactionsScreen, {}),
        subRoutes: {},
      },
    },
  },
};

export default appRoutes;
