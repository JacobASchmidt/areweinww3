/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.htmlm", "./src/**/*.{js,ts,tsx,jsx}"],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui : {
    themes: ["luxury" ]
  }
}

