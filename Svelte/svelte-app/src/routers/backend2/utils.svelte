<script>
    export function isEmpty(obj = {}) {
            return Object.keys(obj).length === 0
        }

    export function isString(value) {
        return typeof value === 'string' || value instanceof String
    }

    export function isNumber(value) {
        return typeof value == 'number' && !isNaN(value)
    }

    export function isBoolean(value) {
        return value === true || value === false
    }

    export function isNil(value) {
        return typeof value === 'undefined' || value === null
    }

    export function isDateString(value) {
        if (!isString(value)) return false

        return value.match(/^\d{2}-\d{2}-\d{4}$/)
    }

    export function convertDateString(value) {
        return value.substr(0, 4) + value.substr(5, 2) + value.substr(8, 2)
    }

    export function toLower(value) {
        if (isString(value)) {
            return value.toLowerCase()
        }
        return value
    }

    export function convertType(value) {
        if (isNumber(value)) {
            return value.toString()
        }

        if (isDateString(value)) {
            return convertDateString(value)
        }

        if (isBoolean(value)) {
            return value ? '1' : '-1'
        }

        return value
    }
</script>