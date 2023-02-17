<script>
	let columns = ["Select", "Author", "Title", "Context"];

    let selectedRow = {};
    let address = "http://localhost:8080/story/";
	

    import axios from 'axios';
    let auth = {
        username: 'abc',
        password: '123'
    };
    
    export let stories = [];
    import { onMount } from 'svelte';
    onMount(async () => {
        const res = await axios.get(address);
        stories = await res.data["context"];
	});

    export let fStories = [];
    function filteredStories() {
        fStories = [];
        let f = document.getElementById("filter")
        for (const story of stories){
            for (const [_, value] of Object.entries(story)){
                if (value.include(f)){
                    fStories.push(story);
                    break;
                }
            }
        }
    }
    

    function filterToURL(filter) {
        res  = "?"
        for (const [key, value] of Object.entries(filter)){
            let add  = "";
            add += key;
            add += "="
            add += value;
            res += add;
            res += "&";
        }
        res -= "&";
        return res;
    }

    async function readStory(){
        let filter = document.getElementById("filter");
        filter = filterToURL(filter)
        try{
            if (filter == {}){
                const res = await axios.get(address, auth)
            }else{
                const res = await axios.get(address + filter, auth)
            }
            data = await res.data["context"];
            console.log(res)
        }catch(err){
            console.log(err)
        }
    }

	async function deleteStories(selectedRow) {
		for (const [key, value] of Object.entries(selectedRow)) {
            if (value != 1) continue;
            try{
                const res = await axios.delete(address + key, auth);
                console.log(res)
            }catch(err){
                console.log(err)
            }
        }
	}

    function createStory(){
        
    }

    function updateStory(){
        
    }

    function click(id){
        selectedRow[id] ^= 1; 
    }

</script>

<div class="sticky"><a href="/">head</a></div>
<input type="text" id="filter" name="filter" placeholder="filter..." on:keyup={filteredStories}>
<button on:click={() => deleteStories(selectedRow)}>delete</button>
<button on:click={() => createStory()}>create</button>

<table>
    <thead>
        <tr>
            {#each columns as column}
                <th>{column}</th>
            {/each}
        </tr>
    </thead>
    <tbody>
        {#each stories as row}
            <tr>
                <td>
                    <input type="checkbox" bind:checked={selectedRow[row[0]]} on:click={() => click(row[0])}>
                </td>
                {#each columns.slice(1) as cell}
                    <td contenteditable="true" bind:innerHTML={row[cell]} />
                {/each}
            </tr>
        {/each}
    </tbody>
</table>


<style>
    :global(body){
        background-color: #fdf6e3;
    }

    div.sticky{
        position: -webkit-sticky;
        position: sticky;
        top: 0;
        padding: 5px;
        background-color: #cae8ca;
        border: 2px solid #eee8d5;
    }

    table, th, td{
        border: 2px solid rgb(138, 10, 218);
    }

    table{
        border-collapse: collapse;
        width: 80%;
        margin: 25px;
        position: absolute;
        left: 0%;
        top: 20%;
    }

    td{
        padding:5px;
    }

    th{
        height: 28px;
        background-color:rgb(103, 10, 113);
        color:white;
    }

</style>