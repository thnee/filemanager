<script>
	import Sidebar from "$lib/Sidebar.svelte";
</script>

<div class="flex">
	<div class="flex-none">
		<Sidebar>
			<h2 class="text-2xl m-4">Settings</h2>
			<ul class="menu">
				<li><a href="/settings">Main</a></li>
				<li><a href="/settings/ui/theme">Theme</a></li>
			</ul>
		</Sidebar>
	</div>

	<div class="p-4">
		<slot />
	</div>
</div>
