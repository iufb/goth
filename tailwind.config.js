/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.{templ,go}",
    "./public/**/*.css",
    "./internal/handlers/**/*.{go,templ}",
    "./cmd/app/**/*.go",
  ],

  safelist: [],
};
