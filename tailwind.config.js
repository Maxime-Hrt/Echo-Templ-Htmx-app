/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      './view/**/*.{templ,html}',
      './view/**/**/*.templ',
    ],
    theme: {},
    plugins: [
      require('@tailwindcss/forms'),
    ],
  }
  