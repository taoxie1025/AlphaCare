import React, {useEffect, useState} from 'react';
import {connect} from 'react-redux';
import styled from 'styled-components';
import {AppState} from '../../@types/global';
import {register} from '../store/actions/user';

const Wrapper = styled.div`
  max-width: 325px;
  width: 100%;
  border: 1px solid #ededed;
  border-radius: 3px;
  box-shadow: 0px 1px 3px rgb(0 0 0 / 25%);
  background-color: #fff;
`;
const Header = styled.div`
  font-size: 28px;
`;

interface Props {
  register?: typeof register;
}

const LoginForm = (props: AppState & Props): JSX.Element => {
  const [newUserName, setNewUsername] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');

  const onLogin = () => {
    props.register(newUserName, newPassword, firstName, lastName);
  };

  return (
    <Wrapper>
      <Header>
        <h2>Register</h2>
      </Header>
      <form onSubmit={onLogin}>
        <div>
          <input
            type="text"
            id="firstname"
            value={firstName}
            placeholder={'Jane'}
            onChange={(e) => setFirstName(e.target.value)}
            autoCorrect="off"
            autoCapitalize="off"
            spellCheck="false"
          />
          <label htmlFor="firstName">First Name</label>
        </div>
        <div>
          <input
            type="text"
            id="lastname"
            value={lastName}
            placeholder="Doe"
            onChange={(e) => setLastName(e.target.value)}
            autoCorrect="off"
            autoCapitalize="off"
            spellCheck="false"
          />
          <label htmlFor="lastName">Last Name</label>
        </div>
        <div>
          <input
            type="text"
            id="username"
            value={newUserName}
            placeholder="Nyx username"
            onChange={(e) => setNewUsername(e.target.value)}
            autoCorrect="off"
            autoCapitalize="off"
            spellCheck="false"
          />
          <label htmlFor="username">Username</label>
        </div>
        <div>
          <input
            id="password"
            type="password"
            value={newPassword}
            placeholder="••••••••••"
            onChange={(e) => setNewPassword(e.target.value)}
          />
          <label htmlFor="password">Password</label>
        </div>

        <button className="form__submit-btn" type="submit">
          Register
        </button>
      </form>
    </Wrapper>
  );
};

export default connect((state) => state, {register})(LoginForm);
