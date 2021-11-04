import Koa from 'koa';
import Router from 'koa-router';
import jsonwebtoken from 'jsonwebtoken';

import {render} from './middlewares/render';
import {createGzip} from 'zlib';

const router = new Router({prefix: '/provider'});

const routes = ['/', '/login', '/register', '/dashboard'];

routes.forEach((route) => {
  router.get(route, (ctx: Koa.Context) => {
    render(ctx);
  });
});

export interface LoginDetails {
  userId: string;
  password: string;
}

router.post('/login', async (ctx: Koa.Context) => {
  let loginDetails: LoginDetails = ctx.request.body;
  try {
    if (loginDetails.userId && loginDetails.password) {
      // Make call to auth service
      // TODO:

      // return result
      const body = {
        token: jsonwebtoken.sign(
          {
            data: ctx.request.body,
            exp: Math.floor(Date.now() / 1000) - 60 * 60, // 60 seconds * 60 minutes = 1 hour
          },
          'secret'
        ),
      };
      ctx.body = body;
    }
  } catch (error) {
    console.log(error);
  }
});

export default router;
