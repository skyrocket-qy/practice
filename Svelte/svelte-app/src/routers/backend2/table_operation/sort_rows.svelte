<script context="module">
    export let sortConfig = {
        'by': '',
        'sequences': {},
        'icons': {},
    }

    export function initSortConfig(sortConfig, columnNames){
        sortConfig['by'] = ''
        for (const columnName of columnNames){
            sortConfig["sequences"][columnName] = ''
            sortConfig['icons'][columnName] = '️↕️'
        }
    }

    export function sort(sortConfig, newColName, rows){
        let sortBy = sortConfig['by']
        if (sortBy != '' && sortBy != newColName){
            sortConfig['sequences'][sortBy] = ''
            sortConfig['icons'][sortBy] = '️↕️'
        }

        sortConfig['sequences'][newColName] = sortConfig['sequences'][newColName] === 'asc' ? 'desc' : 'asc'
        sortConfig['icons'][newColName] = orderToIcon(sortConfig['sequences'][newColName])
        return sortRows(rows, sortConfig['sequences'][newColName], newColName)
    }

    export function orderToIcon(order){
        return order === 'asc' ? '⬆️' : order === 'desc' ? '⬇️' : '️↕️'
    }

    function sortRows(rows, order, colName){
        return rows.sort((a, b) => {
            const [aVal, bVal] = [a[colName], b[colName]]

            if (isNil(aVal)) return 1
            if (isNil(bVal)) return -1

            const [aLocale, bLocale] = [convertType(aVal), convertType(bVal)]

            return order === 'asc' ? aLocale.localeCompare(bLocale) : bLocale.localeCompare(aLocale)
        })
    }

   function isNil(value) {
        return typeof value === 'undefined' || value === null
    }

    function convertType(value) {
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

    function isString(value) {
        return typeof value === 'string' || value instanceof String
    }

    function convertDateString(value) {
        return value.substr(0, 4) + value.substr(5, 2) + value.substr(8, 2)
    }

    function isNumber(value) {
        return typeof value == 'number' && !isNaN(value)
    }

    function isBoolean(value) {
        return value === true || value === false
    }

    function isDateString(value) {
        if (!isString(value)) return false

        return value.match(/^\d{2}-\d{2}-\d{4}$/)
    }
</script>
