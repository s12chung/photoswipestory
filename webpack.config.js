const isProduction = process.env.NODE_ENV === "production";
const filename = isProduction ? '[name]-[hash]' : '[name]';

const defaults = require('gostatic-webpack')(__dirname, filename, isProduction);
const imageDefaults = require('gostatic-webpack-images')(__dirname, filename, isProduction);

const relativePath = function(p) { return require('path').resolve(__dirname, p); };

module.exports = {
    mode: isProduction ? "production" : "development",
    devtool: isProduction ? undefined : "eval",

    entry: Object.assign(defaults.entry(), {
        // entryChunkName: relativePath('assets/js/filename.js'),
    }),
    output: Object.assign(defaults.output(), {
        // customize
    }),

    module: {
        rules: defaults.allRules()
            .concat(imageDefaults.responsiveRules(relativePath('content/markdowns/images'), 'content/markdowns/images/'))
            .concat(imageDefaults.responsiveRules(relativePath('content/markdowns/swiper'), 'content/markdowns/swiper/'))
            .concat(imageDefaults.responsiveRules(relativePath('content/demo/images'), 'content/demo/images/'))
            .concat(imageDefaults.responsiveRules(relativePath('content/demo/swiper'), 'content/demo/swiper/'))
            .concat([
                {
                    test: /\.js$/,
                    exclude: /node_modules\/(?!(dom7|ssr-window|swiper)\/).*/,
                    loader: 'babel-loader'
                }
            ])
    },

    plugins: defaults.allPlugins()
        .concat([
            // customize
        ])
};
