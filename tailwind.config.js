/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.{html,js,go}",
    "./controllers/**/*.go",
    "./main.go"
  ],
  safelist: [
    // Include common color classes
    {
      pattern: /bg-(red|green|blue|yellow|purple|pink|indigo|gray)-(50|100|200|300|400|500|600|700|800|900)/,
    },
    {
      pattern: /text-(red|green|blue|yellow|purple|pink|indigo|gray)-(50|100|200|300|400|500|600|700|800|900)/,
    },
    {
      pattern: /border-(red|green|blue|yellow|purple|pink|indigo|gray)-(50|100|200|300|400|500|600|700|800|900)/,
    },
    // Common utility classes
    'container',
    'mx-auto',
    'px-4',
    'py-2',
    'p-10',
    'flex',
    'grid',
    'hidden',
    'block'
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#eff6ff',
          500: '#3b82f6',
          600: '#2563eb',
          700: '#1d4ed8',
        }
      }
    },
  },
  plugins: [],
}

