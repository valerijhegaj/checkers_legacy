import React from 'react';

const MyContext = React.createContext();

export const Provider = (props) => {
  return <MyContext.Provider value={props.value}>
    {props.children}
  </MyContext.Provider>
}

export default MyContext