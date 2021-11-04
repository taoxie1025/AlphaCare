/*
 * HomePage
 *
 * This is the first thing unauthenticated users see of the app
 * Route: /
 *
 */

import React from 'react';
import {connect} from 'react-redux';
import {AppState} from '../../../@types/global';

const HomePage = (props: AppState): JSX.Element => {
  return (
    <article>
      <div>Hello, homepage</div>
    </article>
  );
};

export default connect((state) => state)(HomePage);
