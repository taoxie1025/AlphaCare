import * as React from 'react';
import {render} from 'react-dom';
import {Provider} from 'react-redux';

import {App} from './components/app';
import {BrowserRouter as Router, Switch, Route, useHistory} from 'react-router-dom';
import HomePage from './components/pages/HomePage';
import initStore from './store/initStore';
import Nav from './components/Nav';
import LoginPage from './components/pages/LoginPage';
import RegisterPage from './components/pages/RegisterPage';
import Dashboard from './components/pages/Dashboard';
// import NotFound from './components/pages/NotFound';
import {Security, SecureRoute, LoginCallback} from '@okta/okta-react';
import {oktaAuthConfig, oktaSignInConfig} from './okta/config';
import {OktaAuth, toRelativeUrl} from '@okta/okta-auth-js';

const MOUNT_NODE = document.getElementById('root') as HTMLElement;

const oktaAuth = new OktaAuth(oktaAuthConfig);

function init() {
  const store = initStore();
  // const history = useHistory();

  const customAuthHandler = () => {
    // history.push('/login');
  };

  const restoreOriginalUri = async (_oktaAuth: any, originalUri: any) => {
    // history.replace(toRelativeUrl(originalUri, window.location.origin));
  };

  render(
    <React.StrictMode>
      <Provider store={store}>
        <App>
          <Router basename="provider">
            <Nav />
            <Security
              oktaAuth={oktaAuth}
              onAuthRequired={customAuthHandler}
              restoreOriginalUri={restoreOriginalUri}
            >
              <Switch>
                <Route exact path="/" component={HomePage} />
                <Route path="/login" component={LoginPage} />
                <Route path="/register" component={RegisterPage} />
                <SecureRoute path="/dashboard" component={Dashboard} />
                {/* <Route onEnter={checkAuth}>
            <Route path="/login" component={LoginPage} />
          </Route> */}
                {/* <Route exact path="*" component={NotFound} /> */}
              </Switch>
            </Security>
          </Router>
        </App>
      </Provider>
    </React.StrictMode>,
    MOUNT_NODE
  );
}
if (document.readyState !== 'loading') {
  init();
} else {
  document.addEventListener('DOMContentLoaded', init);
}
// TODO: report web vitals
