/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./internal/web/**/*.go'],
  plugins: [
    // https://github.com/tailwindlabs/tailwindcss-forms
    require('@tailwindcss/forms'),
  ],
  theme: {
    screens: {
      'desk': '768px', // only one breakpoint to keep it simple
    },
    extend: {},
  }
}
