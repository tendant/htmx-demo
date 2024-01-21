/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,tmpl}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
}
