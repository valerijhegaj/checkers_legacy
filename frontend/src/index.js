import React from 'react';
import ReactDOM from "react-dom/client";
import App from "./App";
import store from "./usingRedux/store";
import {Provider} from "./redux/react-redux/Provider";

const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
  <Provider value={store}>
    <App/>
  </Provider>
)
