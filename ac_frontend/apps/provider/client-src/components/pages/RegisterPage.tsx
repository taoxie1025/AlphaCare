import React from 'react';
import {connect} from 'react-redux';
import styled from 'styled-components';
import {AppState} from '../../../@types/global';
import RegisterForm from '../RegisterForm';

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  width: 100%;
`;
const Register = (props: AppState): JSX.Element => {
  return (
    <Wrapper>
      <RegisterForm />
    </Wrapper>
  );
};

export default connect((state) => state)(Register);
