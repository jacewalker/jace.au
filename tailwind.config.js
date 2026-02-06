/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*/*.html"],
  theme: {
    screens: {
      'xs': '450px',
      'sm': '640px',
      'md': '768px',
      'lg': '1024px',
      'xl': '1280px',
      '2xl': '1536px',
    },
    extend: {
      fontFamily: {
        'display': ['Bricolage Grotesque', 'sans-serif'],
        'sans': ['DM Sans', 'sans-serif'],
        'mono': ['Fira Code', 'monospace'],
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
