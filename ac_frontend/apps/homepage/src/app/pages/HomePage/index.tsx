import * as React from 'react';
import { Helmet } from 'react-helmet-async';

export function HomePage() {
  return (
    <>
      <Helmet>
        <title>AlphaCare Medical</title>
        <meta name="description" content="A Boilerplate application homepage" />
      </Helmet>
      <span>Hello Nyx!</span>
    </>
  );
}
