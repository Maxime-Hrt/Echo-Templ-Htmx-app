/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      './view/**/*.templ',
      './view/**/**/*.templ',
    ],
    theme: {},
    plugins: [
      require('@tailwindcss/forms'),
    ],
  }
  