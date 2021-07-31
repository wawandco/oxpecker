module.exports = {
    plugins: {
        tailwindcss: {},
        '@fullhuman/postcss-purgecss': {
            content: ['./**/*.html', './**/*.js'],
        },
        autoprefixer: {
            browsers: [
                "last 2 versions",
                "Explorer >= 11",
            ]
        },
    }
}