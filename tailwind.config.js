/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/**/*.gohtml"],
  theme: {
    extend: {
        colors: {
            "green": "#26BB81",
        },
    },
  },
  future: {
    hoverOnlyWhenSupported: true,
  },
  plugins: [],
}

