const plugin = require('tailwindcss/plugin')

module.exports = plugin(function ({addComponents}) {
    addComponents({
        '.input svg': {
            '@apply w-4 h-4': {},
        },
    })
})
