
<script>

	import Home from './routers/Home.svelte';
	import Backend from './routers/backend/Backend.svelte';
	import Backend2 from './routers/backend2/Backend2.svelte'
	import router from 'page';

	let page
	
	router('/', () => page = Home)
	router('/backend', () => page = Backend)
	router('/backend2', () => page = Backend2)
	router.start()
</script>

<main>
	<svelte:component this={page} />
</main>

<style>

</style>