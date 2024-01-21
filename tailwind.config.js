/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./static/templates/**/*.{html,tmpl}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
}
