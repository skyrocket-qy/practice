<script context='module'>

    let selectedRow = {};
    let address = "http://localhost:8080/story/";
	
    let postData = {
        name: '',
        age: '',
        is_manager: '',
        start_date: '',
    }

    import axios from 'axios';
    let auth = {
        username: 'abc',
        password: '123'
    };
/* 
    let dataCache = [];
    import { onMount } from 'svelte';
    onMount(dataCache = readAllOpe()) */

    async function readAllOpe(){
        try{
            const res = await axios.get(address, auth)
            const data = await res.data["context"]
            return data
        }catch(err){ 
            console.error(err)
        }
    }

    async function readOneOpe(id){
        try{
            const res = await axios.get(address + filter, auth)
            data = await res.data["context"]
            return data
        }catch(err){
            console.error(err)
        }
    }

	export async function deleteOpe(selectedRow) {
		for (const [key, value] of Object.entries(selectedRow)) {
            if (value != 1) continue
            try{
                const res = await axios.delete(address + key, auth)
                console.log(res)
            }catch(err){
                console.error(err)
            }
        }
	}

    async function createOpe(data){
        try{
            const res = await axios.post(address, data)
            console.log(res)
        }catch(err){
            console.error(err)
            return false
        }
    }

    async function updateOpe(id, data){
        try{
            const res = await axios.put(address + id, data)
            console.log(res)
        }catch(err){
            console.error(err)
            return false
        }
    }

    /* async function readOneOpe(id){
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
            console.error(err)
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
    } */
</script>

<button on:click={() => deleteOpe(selectedRow)}>delete</button>
<button on:click={() => createOpe()}>create</button>