module.exports = {
    purge: ["./content/**/*.md", "./content/**/*.html", "./layouts/**/*.html"],
};

const colors = require('tailwindcss/colors')

module.exports = {
  theme: {
    colors: {
      primary: colors.sky,
      black: colors.black,
      white: colors.white,
      gray: colors.gray,
      transparent: 'transparent',
    }
  }
}