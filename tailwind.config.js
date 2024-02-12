/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{go, templ}"],
  theme: {
    extend: {
      colors: {
        primary: "#061012",
        green: "#5CA05A",
        darkGreen: "#102916",
        darkYellow: "#2B2511",
        dark: "#11292B",
        yellow: "#B6A155",
      },
    },
  },
  plugins: [],
};
