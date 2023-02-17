<script context="module">
    export let filterConfig = {
        'filters': {},
    }

    export function initFilterConfig(filterConfig, columnNames){
        for (const columnName of columnNames) filterConfig["filters"][columnName] = ''
    }

    export function filter(filterConfig, rows, newColName, eventValue){
        filterConfig['filters'][newColName] = eventValue
        return filterRows(rows, filterConfig['filters'])
        
    }

    function filterRows(rows, filters) {
        if (isEmpty(filters)) return rows

        return rows.filter((row) => {
            return Object.keys(filters).every((filterBy) => {
                const value = row[filterBy]
                const searchValue = filters[filterBy]

                if (searchValue === "") return true
                else if (isString(value)) return toLower(value).includes(toLower(searchValue))
                else if (isBoolean(value)) return value.toString().includes(searchValue)//(searchValue === 'true' && value) || (searchValue === 'false' && !value)
                else if (isNumber(value)) return value.toString().includes(searchValue)

                return false
            })
        })
    }

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

    export function toLower(value) {
        if (isString(value)) return value.toLowerCase()
        return value
    }
</script>