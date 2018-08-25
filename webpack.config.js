const isProduction = process.env.NODE_ENV === "production";
const filenameF = function() { return isProduction ? '[name]-[hash]' : '[name]'; };

const defaults = require('gostatic-webpack')(__dirname, filenameF);
const imageGenerator = require('gostatic-webpack-images')(__dirname, filenameF);

const relativePath = function(p) { return require('path').resolve(__dirname, p); };

module.exports = {
    mode: isProduction ? "production" : "development",
    devtool: "cheap-source-map",

    entry: Object.assign(defaults.entry(), {
        client: relativePath('assets/js/client.js'),
    }),
    output: Object.assign(defaults.output(), {
        // add your own
    }),

    module: {
        rules: defaults.allRules()
            .concat(imageGenerator.responsiveRules(relativePath('content/markdowns/images'), 'content/images/'))
            .concat(imageGenerator.responsiveRules(relativePath('content/markdowns/swiper'), 'content/swiper/'))
            .concat([
                {
                    test: /\.js$/,
                    exclude: /node_modules\/(?!(dom7|ssr-window|swiper)\/).*/,
                    loader: 'babel-loader'
                }
            ])
    },

    plugins: defaults.allPlugins().concat([
        // add your own
    ])
};