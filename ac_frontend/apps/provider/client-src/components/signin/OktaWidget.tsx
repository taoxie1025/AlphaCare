declare module '@okta/okta-signin-widget';

import React, {useEffect, useRef} from 'react';
import OktaSignIn from '@okta/okta-signin-widget';
// import '@okta/okta-signin-widget/dist/css/okta-sign-in.min.css';

interface Props {
  config?: any;
  onSuccess: (tokens: any) => void;
  onError: (error: Error) => void;
}

const OktaSignInWidget = ({config, onSuccess, onError}: Props) => {
  const widgetRef = useRef();
  useEffect(() => {
    if (widgetRef.current) {
      const widget = new OktaSignIn(config);

      widget
        .showSignInToGetTokens({
          el: widgetRef.current,
        })
        .then(onSuccess)
        .catch(onError);

      return;
    }

    // toDO:
    // return () => widget.remove();
  }, [config, onSuccess, onError]);

  return <div ref={widgetRef} />;
};

export default OktaSignInWidget;
