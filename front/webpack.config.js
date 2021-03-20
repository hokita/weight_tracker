const path = require('path')
const webpack = require('webpack')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const src = path.resolve(__dirname, 'src')
const dist = path.resolve(__dirname, 'dist')

module.exports = {
  mode: 'development',
  entry: src + '/index.tsx',

  output: {
    path: dist,
    filename: 'bundle.js',
  },

  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        loader: 'ts-loader',
        include: [path.resolve(__dirname, 'src')],
        exclude: [/node_modules/],
      },
    ],
  },

  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
  },

  devServer: {
    host: '0.0.0.0',
    port: 8080,
  },

  plugins: [
    new webpack.ProgressPlugin(),
    new HtmlWebpackPlugin({
      template: 'index.html',
    }),
    new webpack.DefinePlugin({
      'process.env.API_DOMAIN': JSON.stringify(
        process.env.API_DOMAIN || 'localhost'
      ),
    }),
  ],
}
