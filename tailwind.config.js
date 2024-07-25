// tailwind config from tut

/** @type {import{'tailwindcss').Config} */
module.exports = {
  content: ["./view/*.templ", "./**/*.templ"],
  safelist: [],
  theme: { },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["synthwave"]
  }
}
