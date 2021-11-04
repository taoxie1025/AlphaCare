import Koa from 'koa';

import bodyParser from 'koa-body';
import cookieParser from 'koa-cookie';
import userAgent from 'koa-useragent';
import serve from 'koa-static-server';
import cors from '@koa/cors';
import subdomain from 'koa-subdomain';

import router from './router';

const port = '3333';
const app = new Koa();
const subDomain = new subdomain();

async function main() {
  // if localhost assign 1
  if (process.env.NODE_ENV === 'development') {
    app.subdomainOffset = 1
  }

  // setup subdomain module
  subDomain.use('*', router.routes());

  app
    .use(async (ctx: Koa.Context, next: () => Promise<any>) => {
      try {
        await next();
      } catch (error) {
        // TODO: error reporting
        ctx.status = error.statusCode || error.status;
        error.status = ctx.status;
        ctx.body = {error};
        ctx.app.emit('error', error, ctx);
      }
    })
    .use(bodyParser())
    .use(cookieParser())
    .use(userAgent)
    .use(
      serve({
        rootDir: 'static',
        rootPath: '/static',
      })
    )
    .use(subDomain.routes())
    // .use(router.routes())
    .use(router.allowedMethods())
    .use(cors());

  app.listen(port, () => {
    console.log(`app is running on port ${port}`);
  });
}

main().catch(console.log);
