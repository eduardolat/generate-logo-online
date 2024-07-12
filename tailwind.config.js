/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./internal/web/**/*.go'],
  plugins: [require('daisyui')],
  daisyui: {
    logs: false,
  },
  theme: {
    screens: {
      'desk': '768px', // only one breakpoint to keep it simple
    },
    extend: {},
  }
}
