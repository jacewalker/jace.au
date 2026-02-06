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
        'serif': ['Instrument Serif', 'serif'],
        'sans': ['Outfit', 'sans-serif'],
        'mono': ['JetBrains Mono', 'monospace'],
      },
      colors: {
        'accent': '#00e5a0',
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
