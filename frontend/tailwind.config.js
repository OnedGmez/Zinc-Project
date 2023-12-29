/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')

export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      black: colors.black,
      white: colors.white,
      gray: {
        '50':  '#f8f9f7',
        '100': '#efeeef',
        '200': '#ded9df',
        '300': '#bdb6bd',
        '400': '#988f98',
        '500': '#7c6e77',
        '600': '#65535c',
        '700': '#4e3f47',
        '800': '#372c32',
        '900': '#221c20',
      },
      red: {
        50: '#ffd6d6',
        100: '#FF4B4B',
      }
    },
    screens: {
      sm: '480px',
      md: '768px',
      lg: '976px',
      xl: '1440px',
    },
    fontFamily: {
      sans: ['Graphik', 'sans-serif'],
      serif: ['Merriweather', 'serif'],
    },
    extend: {
      spacing: {
        '8xl': '96rem',
        '9xl': '128rem',
      },
      borderRadius: {
        '4xl': '2rem',
      },
      height:{
        'normal': '94%',
        'restante': '6%'
      }
    },
  },
  plugins: [],
}

