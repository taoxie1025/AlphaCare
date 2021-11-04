import React from 'react';
import {connect} from 'react-redux';
import styled from 'styled-components';
import {AppState} from '../../../@types/global';
import LoginForm from '../LoginForm';

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  width: 100%;
`;
const Login = (props: AppState): JSX.Element => {
  return (
    <Wrapper>
      <LoginForm />
    </Wrapper>
  );
};

export default connect((state) => state)(Login);
