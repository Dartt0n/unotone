/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./controllers/www/components/**/*.templ"],
  theme: {
    extend: {
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
