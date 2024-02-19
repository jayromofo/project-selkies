/* When using a custom file name, you need to specify it in cmd when compiling with tailwind cli tool */
/* npx tailwindcss init tailwindcss-config.js */
/* npx tailwindcss -c ./tailwindcss-config.js -i input.css -o output.css */

import type { Config } from 'tailwindcss'

export default {
  important: '#app',
  content: [
    "./index.html",
    "./src/pages/**/*.{vue,js,ts,jsx,tsx}",
    "./src/components/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      'blue': '#1fb6ff',
      'purple': '#7e5bef',
      'pink': '#ff49db',
      'orange': '#ff7849',
      'green': '#13ce66',
      'yellow': '#ffc82c',
      'gray-dark': '#273444',
      'gray': '#8492a6',
      'gray-light': '#d3dce6',
      'heading': '#010203',
      'copy': '#234345',
      'accent': '#268baa',
      'border': '#e2e4e8',
      'dim': '#59797b',
    },
    fontFamily: {
      sans: ['Graphik', 'sans-serif'],
      serif: ['Merriweather', 'serif'],
    },
    fontSize: {
      sm: ['0.875rem', { lineHeight: '1rem' }],
    },
    screens: {
      lp: { max: '1440px' },
      tl: { max: '1199px' },
      tp: { max: '1023px' },
      ph: { max: '767px' },
    },
    extend: {
      spacing: {
        '8xl': '96rem',
        '9xl': '128rem',
        '15': '3.75rem',
        '23': '5.75rem',
        'content': '48.75rem',
      },
      borderRadius: {
        '4xl': '2rem',
      }
    }
  },
  plugins: [],
} satisfies Config