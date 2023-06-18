<script lang='ts'>
	// The ordering of these imports is critical to your app working properly
	import '@skeletonlabs/skeleton/themes/theme-skeleton.css';
	// If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
	import '@skeletonlabs/skeleton/styles/skeleton.css';
	// Most of your app wide CSS should be put in this file
	import '../app.postcss';
	import '../app.css';

	import { AppShell } from '@skeletonlabs/skeleton';
	import { AppRail, AppRailAnchor } from '@skeletonlabs/skeleton';

	let currentTile: number = 0;
</script>

<AppShell>
	<svelte:fragment slot="sidebarLeft">
<AppRail>
	<svelte:fragment slot="lead">
		<AppRailAnchor href="/" >Home</AppRailAnchor>
	</svelte:fragment>
	<!-- --- -->
</AppRail>
	</svelte:fragment>
	<slot />
</AppShell>
