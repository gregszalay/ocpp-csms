import React, { Context, ReactElement } from "react";

const FirebaseDataContext: Context<any> = React.createContext(null);

export const withFirebaseDataContext = (
  MyComponent: (props: any) => JSX.Element,
  props: any
) => (
  <FirebaseDataContext.Consumer>
    {({
      firebase,
      stations,
      chargetokens,
      transactions,
      userInfo,
      handleNewToken,
      handleModifiedToken,
      handleDeletedToken,
      isSmUp,
      setRefresh,
    }) => (
      <MyComponent
        {...props}
        firebase={firebase}
        stations={stations}
        chargetokens={chargetokens}
        transactions={transactions}
        userInfo={userInfo}
        handleNewToken={handleNewToken}
        handleModifiedToken={handleModifiedToken}
        handleDeletedToken={handleDeletedToken}
        isSmUp={isSmUp}
        refresh={setRefresh}
      />
    )}
  </FirebaseDataContext.Consumer>
);

export default FirebaseDataContext;
