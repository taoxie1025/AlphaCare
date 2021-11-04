import * as React from 'react';
import styled from 'styled-components';

interface PropTypes {
  children: JSX.Element;
}

const Container = styled.div`
  margin: 0 auto;
  max-width: 1200px;
`;

export function App(props: PropTypes) {
  return <Container>{props.children}</Container>;
}
