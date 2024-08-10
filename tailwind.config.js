/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./internal/web/**/*.go'],
  plugins: [require('daisyui')],
  daisyui: {
    logs: false,
    themes: ['light', 'dracula'],
    darkTheme: 'dracula'
  },
  theme: {
    screens: {
      desk: '768px' // only one breakpoint to keep it simple
    },
    extend: {}
  }
}
