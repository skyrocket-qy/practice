<script>
    import { initSortConfig, sortConfig, sort } from './table_operation/sort_rows.svelte'
    import { filter, filterConfig, initFilterConfig } from './table_operation/filter_rows.svelte'
    import { paginateRows } from './table_operation/pagination.svelte'
    import { deleteOpe} from './crud.svelte'
    import { columns, rows } from './data.svelte'

    let columnNames = []
    for (const column of columns) columnNames.push(column.name)
    let sortBy = ''
    let filterName = ''
    let filterValue = ''
    let rowsPerPage = 5
    let curPage = 1
    let selectedRow = {}
    

    $: icons = sortConfig.icons
    $: start = (curPage-1) * rowsPerPage + 1
    $: end = curPage * rowsPerPage 
    $: lenOfRows = finalRows.length
    $: totalPages = Math.ceil(finalRows.length / rowsPerPage)

    $: paginatedRows = paginateRows(finalRows, curPage, rowsPerPage)
    $: filteredRows = filter(filterConfig, rows, filterName, filterValue)
    $: finalRows = sort(sortConfig, sortBy, filteredRows)

    function init(){
        sortBy = ''
        initSortConfig(sortConfig, columnNames)
        sortConfig.icons = sortConfig.icons
        initFilterConfig(filterConfig, columnNames)
        curPage = 1
        for (const name of columnNames) selectedRow[name] = false
        filteredRows = filter(filterConfig, rows, filterName, filterValue)
    }
    init()

    function activateSort(columnName){
        sortBy = ''
        sortBy = columnName
        sortConfig.icons = sortConfig.icons
    }

    function activateFilter(columnName, value){
        filterName = columnName
        filterValue = value
    }

    function setCurPage(page){
        curPage = page
    }

    function select(id){
        selectedRow[id] ^= 1
    }

    function add(e){
        const formData = new FormData(e.target)
        
        const data = {}
        for (let field of formData){
            const[key, value] = field
            data[key] = value
        }
        console.log(data)
    }

</script>
<button type="button" on:click={deleteOpe}>delete</button>
<table>
    <thead>
        <tr>
            <th>select</th>
            {#each columns as column}
            <th>
                {column.name}
                <button on:click={() => activateSort(column.name)}>{icons[column.name]}</button>
            </th>
            {/each}
        </tr>
        <tr>
            <th></th>
            {#each columns as column}
            <th>
                <input type="text" placeholder={`search ${column.name}`} value={filterConfig['filters'][column.name]} on:input={(event) => activateFilter(column.name, event.target.value)}/>
            </th>
            {/each}
        </tr>
    </thead>
    <tbody>
        {#each paginatedRows as row}
        <tr>
            <td>
                <input type="checkbox" bind:checked={selectedRow[row['id']]} on:click={() => select(row['id'])}>
            </td>
            {#each Object.entries(row) as [k, cell]}
                <td>
                    {cell}
                </td>
            {/each}
        </tr>
        {/each}
    </tbody>
</table>
<div class="pagination">
    <button disabled={curPage === 1} on:click={() => setCurPage(1)}>
        ⏮️ First
    </button>
    <button disabled={curPage === 1} on:click={() => setCurPage(curPage-1)}>
        ⬅️ Previous
    </button>
    <button disabled={curPage === totalPages} on:click={() => setCurPage(curPage+1)}>
        Next ➡️
    </button>
    <button disabled={curPage === totalPages} on:click={() => setCurPage(totalPages)}>
        Last ⏭️
    </button>
</div>
<p>
Page {curPage} of {totalPages}
</p>
<p>
Rows: {end >= lenOfRows ? `${start} - ${lenOfRows}` : `${start} - ${end}`} of {lenOfRows}
</p>
<div>
    <p>
      <button on:click={init}>Clear all Condition</button>
    </p>
</div>
<form on:submit|preventDefault={add}>
    {#each columns as column}
        <label for={column.name}>{column.name}</label>
        <input type="text" name={column.name}>
    {/each}
    <button type="submit">Add</button>
    <button type="reset">Clear input</button>
</form>


<style>
    * {
        box-sizing: border-box;
    }
    
    html {
        font-family: sans-serif;
    }
    
    table {
        width: 100%;
        border: 1px solid black;
    }
    
    th,
    td {
        border: 1px solid black;
        padding: 0.5rem;
    }
    
    th span {
        margin-right: 1rem;
    }
    
    input {
        border: 1px solid gray;
        padding: 0.5rem;
        width: 100%;
        max-width: 100%;
        background: #f0f0f0;
    }
    
    .pagination {
        margin-top: 0.25rem;
        display: flex;
        justify-content: space-between;
        border: 1px solid black;
        padding: 0.5rem;
    }
</style>