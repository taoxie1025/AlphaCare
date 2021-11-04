import React from 'react';
import {Link} from 'react-router-dom';

const notFound = (): JSX.Element => {
  return (
    <article>
      <h1>Page not found.</h1>
      <Link to="/"> Home </Link>
    </article>
  );
};

export default notFound;
