import React, { Context, ReactElement } from 'react';
import { JsxElement } from 'typescript';

const SelectedMenuItemContext: Context<any> = React.createContext(null);

export const withSelectedMenuItemContext = (MyComponent: ((props:any)=>JSX.Element), props:any) => (
    <SelectedMenuItemContext.Consumer>
      {({selectedMenuItem, handleMenuSelection}) => <MyComponent {...props} selectedMenuItem_fromContext={selectedMenuItem} 
      handleMenuClick_fromContext={handleMenuSelection}/>}
    </SelectedMenuItemContext.Consumer>
  );
  

export default SelectedMenuItemContext;