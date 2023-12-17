/** @type {import('tailwindcss').Config} */
import plugin from "tailwindcss/plugin"

module.exports = {
  content: ["**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [
    // Requires that we have npm package installed, later.
    //
    // plugin(function({ addVariant }) {
    //   addVariant('htmx-settling', ['&.htmx-settling', '.htmx-settling &']),
    //   addVariant('htmx-request',  ['&.htmx-request',  '.htmx-request &']),
    //   addVariant('htmx-swapping', ['&.htmx-swapping', '.htmx-swapping &']),
    //   addVariant('htmx-added',    ['&.htmx-added',    '.htmx-added &']),
    //   addVariant('htmx-indicator', ['&.htmx-indicator',    '.htmx-indicator &']),
    // })
  ],
}

