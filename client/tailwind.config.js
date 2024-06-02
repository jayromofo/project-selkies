/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "./formkit.theme.ts" // formkit theme file
  ],
  darkmode: 'class',
  theme: {
    colors: {
      primary: "#ff0000",
    },
    grey: {
      100: "#f7f7f7",
      200: "#e5e5e5",
    },
    spacing: {},
    fontSize: {},
    fontFamily: {},
    extend: {},
  },
  plugins: [],
}