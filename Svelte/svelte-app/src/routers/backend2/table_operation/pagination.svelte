<script context='module'>

    export const Pagination = ({ activePage, count, rowsPerPage, totalPages, setActivePage }) => {
        const beginning = activePage === 1 ? 1 : rowsPerPage * (activePage - 1) + 1
        const end = activePage === totalPages ? count : beginning + rowsPerPage - 1
    }

    export function paginateRows(rows, curPage, rowsPerPage) {
        return rows.slice((curPage - 1) * rowsPerPage, curPage * rowsPerPage)
    }

</script>
