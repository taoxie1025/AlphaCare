import {readFileSync} from 'fs';
import Koa from 'koa';
import {resolve} from 'path';
import nunjucks, {Environment, FileSystemLoader} from 'nunjucks';

export async function render(ctx: Koa.Context) {
  const templatedirectory = resolve(__dirname, '..', 'templates/');
  const fileLoader = new FileSystemLoader(templatedirectory);
  const environemnt = new Environment(fileLoader);
  const template = readFileSync(templatedirectory + '/index.njk', 'utf-8');
  const precompiledTemplate = nunjucks.compile(template, environemnt);

  try {
    ctx.response.type = 'html';

    const body = await precompiledTemplate.render({
      staticDir: 'static',
    });
    ctx.body = body;
  } catch (err) {
    ctx.throw(err);
  }
}
