const path = require('path');
const CopyPlugin = require('copy-webpack-plugin');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const nodeExternals = require('webpack-node-externals');

const clientConfig = {
  target: 'web',
  mode: 'development', // todo: change based on args

  entry: {
    'client.bundle': path.join(__dirname, 'client-src/index.tsx'),
  },

  // Path and filename of your result bundle.
  // Webpack will bundle all JavaScript into this file
  output: {
    path: path.join(__dirname, 'static'),
    filename: '[name].js',
    publicPath: path.join('/provider/', 'static'),
  },

  resolve: {
    alias: {
      react: path.resolve('./node_modules/react'),
    },
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
  },

  module: {
    rules: [
      {
        test: /\.js|\.jsx$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
      },
      {
        test: /\.ts|\.tsx$/,
        loader: 'ts-loader',
        include: __dirname,
        exclude: /node_modules/,
        options: {
          configFile: 'tsconfig.json',
        },
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader'],
        include: __dirname,
      },
    ],
  },

  plugins: [
    new CopyPlugin({
      patterns: [
        {
          from: path.resolve(__dirname, './client-src/normalize.css'),
          to: path.resolve(__dirname, './static/'),
          force: true,
        },
      ],
      options: {
        concurrency: 100,
      },
    }),
    new CleanWebpackPlugin({cleanOnceBeforeBuildPatterns: ['**/*']}),
  ],
};

const serverConfig = {
  target: 'node',
  mode: 'development', // todo: change based on args

  entry: {
    app: path.join(__dirname, 'server/server.ts'),
  },

  // Path and filename of your result bundle.
  // Webpack will bundle all JavaScript into this file
  output: {
    path: path.join(__dirname, 'dist'),
    filename: 'server.bundle.js',
    publicPath: path.join('/provider/', 'static'),
  },

  externals: [
    nodeExternals({
      allowlist: [/\.mjs?$/],
    }),
  ],

  node: {__dirname: true},

  // devServer: {
  //     contentBase: path.join(__dirname, 'dist'),
  //     compress: true,
  //     port: 9000,
  // },

  resolve: {
    alias: {
      react: path.resolve('./node_modules/react'),
    },
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
  },

  module: {
    rules: [
      {
        test: /\.js|\.jsx$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
      },
      {
        test: /\.ts|\.tsx$/,
        loader: 'ts-loader',
        include: __dirname,
        exclude: /node_modules/,
        options: {
          configFile: 'tsconfig.json',
        },
      },
    ],
  },
};

module.exports = [clientConfig, serverConfig];
