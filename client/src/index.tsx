import React from 'react';
import ReactDOM from 'react-dom';
import { App } from './App';
import { mergeStyles, loadTheme } from '@fluentui/react';
import * as serviceWorker from './serviceWorker';
import { Provider } from 'react-redux'
import rootReducer from './rootReducer'
import { configureStore } from '@reduxjs/toolkit'
import { WebSocketProvider } from './WebSocket';
import { initializeIcons } from '@uifabric/icons';

initializeIcons();

// Inject some global styles
mergeStyles({
  selectors: {
    ':global(body), :global(html), :global(#root)': {
      margin: 0,
      padding: 0,
      height: '100vh'
    }
  }
});

loadTheme({
  palette: {
    themePrimary: '#914483',
    themeLighterAlt: '#fbf5fa',
    themeLighter: '#edd9ea',
    themeLight: '#debbd8',
    themeTertiary: '#bd81b2',
    themeSecondary: '#9f5591',
    themeDarkAlt: '#833d76',
    themeDark: '#6e3464',
    themeDarker: '#512649',
    neutralLighterAlt: '#faf9f8',
    neutralLighter: '#f3f2f1',
    neutralLight: '#edebe9',
    neutralQuaternaryAlt: '#e1dfdd',
    neutralQuaternary: '#d0d0d0',
    neutralTertiaryAlt: '#c8c6c4',
    neutralTertiary: '#a19f9d',
    neutralSecondary: '#605e5c',
    neutralPrimaryAlt: '#3b3a39',
    neutralPrimary: '#323130',
    neutralDark: '#201f1e',
    black: '#000000',
    white: '#ffffff',
  }});

const store = configureStore({
  reducer: rootReducer
});

ReactDOM.render(
    <Provider store={store}>
      <WebSocketProvider>
        <App />
      </WebSocketProvider>
    </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
