/**
 *
 * Nav.tsx
 *
 * This component renders the navigation bar
 *
 */

import React from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router-dom';
import styled from 'styled-components';
import {AppState} from '../../@types/global';

const Button = styled(Link)`
  font-size: 0.8em;
  margin-left: 1em;
  padding: 0.5em 1.5em;
  color: #6cc0e5;
  border: 1px solid #6cc0e5;
  border-radius: 3px;
  text-decoration: none;
  user-select: none;
  display: inline-block;
`;

const NavWrapper = styled.div`
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 0 auto;
  padding-bottom: 40px;
`;

const Logo = styled.div`
  font-size: 1em;
  margin: 0;
  user-select: none;
  color: #222;
  text-decoration: none;
`;

const Nav = (props: AppState): JSX.Element => {
  // Render either the Log In and register buttons, or the logout button
  // based on the current authentication state.
  const navButtons = props.user.isLoggedIn ? (
    <div>
      <Button to="/dashboard">Dashboard</Button>
    </div>
  ) : (
    <div>
      <Button to="/register">Register</Button>
      <Button to="/login">Login</Button>
    </div>
  );

  return (
    <NavWrapper>
      <Logo>NYX</Logo>
      {navButtons}
    </NavWrapper>
  );
};

export default connect((state) => state)(Nav);
