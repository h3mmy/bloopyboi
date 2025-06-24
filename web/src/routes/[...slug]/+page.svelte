<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<h1>Bloopyboi</h1>
{#if data}
	<p>{data.appInfo}</p>
{/if}
