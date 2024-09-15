/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./controllers/web/components/**/*.templ"],
  theme: {
    extend: {
      colors: {
        dune: "#D9CAB3",
        "rich-black": "#0D1821",
      },
      fontFamily: {
        mono: ["Courier Prime"],
      },
    },
  },
  plugins: [],
  corePlugins: {
    preflight: true,
  },
};
