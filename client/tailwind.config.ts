import type { Config } from 'tailwindcss'
import {nextui} from "@nextui-org/react";
import next from 'next';

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'pf-primary-1': '#87BFCA',
        'pf-primary-2': '#286572',
        'pf-primary-3': '#134651',
        'pf-secondary-1': '#D6EE8E',
        'pf-secondary-2': '#83A128',
        'pf-secondary-3': '#567009',
        'pf-accent-1': '#FCD894',
        'pf-accent-2': '#FCBF49',
        'pf-accent-3': '#C18717',
        'pf-gray-100': '#EAE9E9',
        'pf-gray-300': '#AEBDC0',
        'pf-gray-500': '#4F6C72',
        'pf-gray-700': '#0B3038',
        'pf-gray-900': '#00171C',

      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic':
          'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
    },
  },
  plugins: [],
}
export default config
