module.exports = {
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            }
        ]
    },
    entry: "./static/js",
    output: {
        path: __dirname + '/static',
        publicPath: '',
        filename: 'js/bundle.js'
    },
    devServer: {
        contentBase: './distgetuk'
    }
};