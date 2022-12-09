import React, { Context, ReactElement } from 'react';
import { JsxElement } from 'typescript';

const RouteContext: Context<any> = React.createContext(null);

export const withRouteContext = (MyComponent: ((props:any)=>JSX.Element), props:any) => (
    <RouteContext.Consumer>
      {routeMap => <MyComponent {...props} routeMap={routeMap} />}
    </RouteContext.Consumer>
  );
  

export default RouteContext;