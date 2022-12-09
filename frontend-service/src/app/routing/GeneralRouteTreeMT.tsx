import React from "react";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import ScrollToTop from "./ScrollToTop";

interface Props {
  routeMap: Object;
}

function buildRouteTree(routes: Object) {
  return Object.entries(routes).map(
    ([route, { component, subRoutes }]: any) => {
      if (
        Object.keys(subRoutes).length === 0 &&
        Object.getPrototypeOf(subRoutes) === Object.prototype
      ) {
        return <Route key={route} path={route} element={component} />;
      } else {
        return (
          <Route
            key={route}
            path={route}
            element={
              route === "dashboard" ? (
                component
              ) : (
                <React.Fragment>
                  <ScrollToTop />
                  <Outlet />
                </React.Fragment>
              )
            }
          >
            {
              //Recursive call
              buildRouteTree(subRoutes)
            }
          </Route>
        );
      }
    }
  );
}

function GeneralRouteTreeMT(props: Props) {
  return (
    <BrowserRouter>
      <Routes>{buildRouteTree(props.routeMap)}</Routes>
    </BrowserRouter>
  );
}

export default GeneralRouteTreeMT;
