/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/**/*.gohtml"],
  theme: {
    extend: {
        colors: {
            // "green": "#18a867", //"#26BB81",
            "off-white": "#f7f7f5" // "#FAF3E0"// "#e8e6e1"
        },
        fontFamily: {
            "fancy": ["Great Vibes", "cursive"],
        }
    },
  },
  future: {
    hoverOnlyWhenSupported: true,
  },
  plugins: [],
}

