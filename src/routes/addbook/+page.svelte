<script lang="ts">
	import { goto } from '$app/navigation';

	let title = '';
	let author = '';
	let year: number;
	let count: number; // Add count field

	const addBook = async () => {
		try {
			const response = await fetch('http://localhost:9090/books/add', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					title,
					author,
					year: Number(year), // Ensure year is an integer
					count // Include count field
				})
			});

			if (!response.ok) {
				const errorText = await response.text();
				console.error('Failed to add book:', errorText);
				alert(`Error: ${errorText}`);
				return;
			}

			const data = await response.json();
			console.log('Book added:', data);
			alert(data.message);
			goto('/books');
		} catch (error) {
			console.error('Error adding book:', error);
			alert('Failed to add book');
		}
	};
</script>

<!-- UI -->
<div class="min-h-screen bg-gray-100 p-10">
	<div class="mx-auto max-w-2xl rounded-lg bg-white p-8 shadow-lg">
		<h1 class="mb-6 text-3xl font-bold">➕ Add New Book</h1>

		<form on:submit|preventDefault={addBook} class="space-y-4">
			<input
				type="text"
				placeholder="Book Title"
				bind:value={title}
				required
				class="w-full rounded-lg border p-3"
			/>

			<input
				type="text"
				placeholder="Author"
				bind:value={author}
				required
				class="w-full rounded-lg border p-3"
			/>

			<input
				type="number"
				placeholder="Year"
				bind:value={year}
				required
				class="w-full rounded-lg border p-3"
			/>

			<input
				type="number"
				placeholder="Count"
				bind:value={count}
				required
				class="w-full rounded-lg border p-3"
			/>

			<button
				type="submit"
				class="w-full rounded-lg bg-blue-500 px-6 py-2 text-white hover:bg-blue-600"
			>
				➕ Add Book
			</button>
		</form>

		<button class="mt-4 text-blue-500 hover:underline" on:click={() => goto('/books')}>
			🔙 Back to Book List
		</button>
	</div>
</div>
