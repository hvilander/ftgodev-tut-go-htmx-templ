// tailwind config from tut

/** @type {import{'tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html",  "./**/*.templ", "./**/*.go"],
  safelist: [],
  theme: { },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["synthwave"]
  }
}
