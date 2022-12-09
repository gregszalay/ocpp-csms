import React from "react";
import ReactDOM from "react-dom";
import App from "./app/App";

import appRoutes from "./app/routing/appRoutes";
import Firebase from "./app/apis/firebase/Firebase";
import FirebaseContext from "./app/contexts/Firebase";
import RouteContext from "./app/contexts/Route";
import "./index.css";

console.log("index file ran");

export const startTime = Date.now();

const firebase = new Firebase();

ReactDOM.render(
  <React.StrictMode>
    <FirebaseContext.Provider value={firebase}>
      <RouteContext.Provider value={appRoutes}>
        
        <App firebase={firebase} />
      </RouteContext.Provider>
    </FirebaseContext.Provider>
  </React.StrictMode>,
  document.getElementById("root")
);
