const Webpack = require("webpack");
const Glob = require("glob");
const path = require('path');
const CopyWebpackPlugin = require("copy-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const { WebpackManifestPlugin } = require('webpack-manifest-plugin');
const { CleanWebpackPlugin } = require("clean-webpack-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const LiveReloadPlugin = require('@kooneko/livereload-webpack-plugin');

const configurator = {
  entries: function(){
    var entries = {
      application: [
        './node_modules/jquery-ujs/src/rails.js',
        './app/assets/css/application.scss',
      ],
    }

    Glob.sync("./app/assets/*/*.*").forEach((entry) => {
      if (entry === './app/assets/css/application.scss') {
        return
      }

      let key = entry.replace(/(\.\/app\/assets\/(src|js|css|go)\/)|\.(ts|js|s[ac]ss|go)/g, '')
      if(key.startsWith("_") || (/(ts|js|s[ac]ss|go)$/i).test(entry) == false) {
        return
      }

      if( entries[key] == null) {
        entries[key] = [entry]
        return
      }

      entries[key].push(entry)
    })
    return entries
  },

  plugins() {
    var static_file_mapping = {};
    var plugins = [
      new Webpack.ProvidePlugin({$: "jquery",jQuery: "jquery"}),
      new MiniCssExtractPlugin({filename: "[name].[contenthash].css"}),
      new CopyWebpackPlugin({
        patterns: [{
            from: "./app/assets",
            to: "[path][name][ext]",
            filter: async (resourcePath) => {
              let ignore = resourcePath.match(/.*\/app\/assets\/(css|js)\/.*/g);
              return !ignore;
            },
            transformPath(targetPath, absosutePath) {
              relative_path = path.relative(path.resolve('./assets'),absosutePath)
              static_file_mapping[relative_path] = targetPath;

              return targetPath;
            }
        }]
      }),
      new Webpack.LoaderOptionsPlugin({minimize: true,debug: false}),
      new WebpackManifestPlugin({fileName: "manifest.json", seed: static_file_mapping}),
      new CleanWebpackPlugin()
    ];

    return plugins
  },

  moduleOptions: function() {
    return {
      rules: [
        { test: require.resolve("jquery"),loader: "expose-loader",options: {exposes: ["$", "jQuery"]} },
        {
          test: /\.s[ac]ss$/,
          use: [
            { loader: MiniCssExtractPlugin.loader, options: {publicPath: ''} },
            { loader: "css-loader", options: {sourceMap: true}},
            { loader: "postcss-loader", options: {sourceMap: true}},
            { loader: "sass-loader", options: {sourceMap: true}}
          ]
        },
        
        { test: /\.jsx?$/, loader: "babel-loader", exclude: /node_modules/ },
        { test: /\.(eot|woff|woff2|ttf|svg|png)(\?v=\d+\.\d+\.\d+)?$/, type: 'asset/resource' },
      ]
    }
  },

  buildConfig: function(){
    // NOTE: If you are having issues with this not being set "properly", make
    // sure your GO_ENV is set properly as `buffalo build` overrides NODE_ENV
    // with whatever GO_ENV is set to or "development".
    const env = process.env.NODE_ENV || "development";

    var config = {
      mode: env,
      entry: configurator.entries(),
      output: {
        filename: "[name].[contenthash].js",
        path: `${__dirname}/public/assets`,
        publicPath: ''
      },
      plugins: configurator.plugins(),
      module: configurator.moduleOptions(),
      resolve: {
        extensions: ['.ts', '.tsx', '.js', '.json']
      }
    }

    if( env === "development" ){
      config.plugins.push(new LiveReloadPlugin({appendScript: true}))
      
      return config
    }

    config.optimization = {
      minimizer: [configurator.terser()]
    }

    return config
  },
  
  // Terser returns the unglyfier used in production mode.
  terser: function() {
    return new TerserPlugin({
      terserOptions: {
        compress: {},
        mangle: {
          keep_fnames: true
        },
        output: {
          comments: false,
        },
      },
      extractComments: false,
    })
  }
}

module.exports = configurator.buildConfig()