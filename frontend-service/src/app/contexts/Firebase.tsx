import React, { Context, ReactElement } from 'react';

const FirebaseContext: Context<any> = React.createContext(null);

export const withFirebaseContext = (MyComponent: ((props:any)=>JSX.Element), props:any) => (
    <FirebaseContext.Consumer>
      {(firebase) => <MyComponent {...props} firebase={firebase}  />}
    </FirebaseContext.Consumer>
  );
  

export default FirebaseContext;