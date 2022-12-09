import { ThemeProvider } from "@mui/material/styles";
import { Box, CssBaseline, useMediaQuery } from "@mui/material";

import dashboardMenu from "../dashboard/constants/dashboardMenu";
import appTheme from "./theme/AppTheme";

import React, { useRef, useState } from "react";
import appRoutes from "./routing/appRoutes";
import SelectedMenuItemContext from "./contexts/SelectedMenuItem";
import GeneralRouteTreeMT from "./routing/GeneralRouteTreeMT";
import FirebaseDataContext from "./contexts/FirebaseData";
import {
  collection,
  doc,
  Firestore,
  getDoc,
  getDocs,
  getDocsFromServer,
  getFirestore,
  limit,
  onSnapshot,
  orderBy,
  query,
  where,
} from "firebase/firestore";
import { setUpSnapshotListener } from "./apis/firebase/dbUtils";
import { User, UserCredential } from "firebase/auth";
import Firebase from "./apis/firebase/Firebase";
import { ExistingFirestoreCollection } from "./apis/firebase/MTFirestore";
import * as firebaseAuthApi from "firebase/auth";
import {
  createRecordInDB,
  createRecordInDBWithoutId,
  deleteRecordInDB,
  updateRecordInDB,
} from "./apis/firebase/dbFunctions";
import { startTime } from "..";
import AppPermissionException from "./apis/firebase/AppPermissionException";

import { IdTokenType } from "ocpp-messages-ts/types/AuthorizeRequest";
import { Transaction } from "../transactions/typedefs/Transaction";
import { ChargingStation } from "../stations/typedefs/ChargingStation";

/***************************************************************************/

interface Props {
  firebase: Firebase;
}

/***************************************************************************/

function App(props: Props) {
  console.log("function App called");

  const [selectedMenuItem, setSelectedMenuItem] = React.useState(() => {
    let mainIndex = 0;
    let childIndex = 0;
    dashboardMenu.forEach((mainItem, index1) => {
      mainItem.children.forEach((childItem, index2) => {
        if (window.location.href.includes(childItem.route)) {
          mainIndex = index1;
          childIndex = index2;
        }
      });
    });
    return dashboardMenu[mainIndex].children[childIndex];
  });
  const [authUser, setAuthUser] = useState<UserCredential | null>(null);

  const [userInfo, setUserInfo] = useState<User | null>(null);
  const [db, setDb] = useState<Firestore | null>(null);
  const [stations, setStations] = useState<ChargingStation[]>([]);
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [chargetokens, setChargeTokens] = useState<IdTokenType[]>([]);

  const isSmUp = useMediaQuery(appTheme.breakpoints.up("md"));

  const handleMenuClick = (newMenuItem: MenuItem) => {
    setSelectedMenuItem(newMenuItem);
  };
  //Wait for db and userInfo to arrive (it does not load on component load right away)

  console.log(`App.tsx first render - elapsed: ${Date.now() - startTime} ms`);

  if (props.firebase.userInfo) {
    console.log(
      `App.tsx userInfo arrived - elapsed: ${Date.now() - startTime} ms`
    );
  }
  if (props.firebase.app) {
    console.log(
      `App.tsx firt app obj arrived - elapsed: ${Date.now() - startTime} ms`
    );
  }

  React.useEffect(() => {
    const timeOutID = setTimeout(() => {
      setUserInfo(props.firebase.userInfo);
      setDb(getFirestore(props.firebase.app));
    }, 1000);
    return () => clearTimeout(timeOutID);
  });

  React.useEffect(() => {
    if (!db || !userInfo || !userInfo.email) {
      console.log(
        "=> useffect STATIONS -------!db || !userInfo || !userInfo.email "
      );
      return;
    }
    let q = null;
    q = query(collection(db, "chargingstations"));
    return setUpSnapshotListener(
      (resultItems: ChargingStation[]) => {
        console.log("=> stationResults");
        console.log("=> ", [...resultItems]);
        setStations(resultItems);
      },
      "chargingstations",
      q,
      db,
      userInfo
    );
  }, [db, userInfo]);

  React.useEffect(() => {
    if (!db || !userInfo || !(stations.length > 0)) {
      console.log(
        "=> useffect TRANSACTIONS ------- !db  !currentUserRecord ||!userInfo || !currentUserPermissions.transactions || !(stations.length > 0)"
      );
      return;
    }
    let q = null;
    q = query(collection(db, "transactions"), orderBy("energyTransferStarted", "desc"));
    return setUpSnapshotListener(
      (resultItems: Transaction[]) => {
        console.log("=> transactionResults");
        console.log("=> ", [...resultItems]);
        setTransactions(resultItems);
      },
      "transactions",
      q,
      db,
      userInfo
    );
  }, [db, userInfo, stations]);

  
  React.useEffect(() => {
    if (!db || !userInfo) {
      console.log(
        "=> useffect TOKENS ------- !db  !currentUserRecord ||!userInfo || !currentUserPermissions.transactions || !(stations.length > 0)"
      );
      return;
    }
    let q = null;
    q = query(collection(db, "idTokens") /*orderBy("status"))*/);
    return setUpSnapshotListener(
      (resultItems: IdTokenType[]) => {
        console.log("=> tokenResults");
        console.log("=> ", [...resultItems]);
        setChargeTokens(resultItems);
      },
      "idTokens",
      q,
      db,
      userInfo
    );
  }, [db, userInfo, transactions]);

  const handleNewToken = async (newToken: IdTokenType, id: string) => {
    try {
      await createRecordInDB(
        {
          firebase: props.firebase,
          //db: db,
          record: newToken,
          collectionName: /*ExistingFirestoreCollection.idTokens*/ "idTokens",
        },
        id
      );
    } catch (err) {
      console.error(err);
    }
  };

  const handleModifiedToken = async (
    modifiedToken: IdTokenType,
    id: string
  ) => {
    try {
      await updateRecordInDB(
        {
          firebase: props.firebase,
          record: modifiedToken,
          collectionName: ExistingFirestoreCollection.idTokens,
        },
        id
      );
    } catch (err) {
      console.error(err);
    }
  };

  const handleDeletedToken = async (deletedToken: IdTokenType, id: string) => {
    try {
      await deleteRecordInDB(
        {
          firebase: props.firebase,
          record: deletedToken,
          collectionName: ExistingFirestoreCollection.idTokens,
        },
        id
      );
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <FirebaseDataContext.Provider
      value={{
        firebase: props.firebase,
        stations: stations,
        chargetokens: chargetokens,
        transactions: transactions,
        userInfo: userInfo,
        handleNewToken: handleNewToken,
        handleModifiedToken: handleModifiedToken,
        handleDeletedToken: handleDeletedToken,
        isSmUp: isSmUp,
        setRefresh: (newAuthUser: UserCredential) => setAuthUser(newAuthUser),
      }}
    >
      <SelectedMenuItemContext.Provider
        value={{
          selectedMenuItem: selectedMenuItem,
          handleMenuSelection: handleMenuClick,
        }}
      >
        <ThemeProvider theme={appTheme}>
          <CssBaseline />
          <GeneralRouteTreeMT routeMap={appRoutes} />
        </ThemeProvider>
      </SelectedMenuItemContext.Provider>
    </FirebaseDataContext.Provider>
  );
}

export default App;
