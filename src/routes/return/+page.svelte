<script lang="ts">
	import { onMount } from 'svelte';
	let id;
	let borrowedBooks = [];
	let usn = '';
	let bookName = '';
	let bookID = 0;
	let loading = false;

	// Fetch borrowed books
	const fetchBorrowedBooks = async () => {
		try {
			loading = true;
			const response = await fetch(`http://localhost:9090/borrow/${id}`);
			if (!response.ok) throw new Error('Failed to fetch borrowed books');
			borrowedBooks = await response.json();
		} catch (error) {
			console.error('Error fetching borrowed books:', error);
		} finally {
			loading = false;
		}
	};

	// Return book by USN, Book Name, and Book ID
	const returnBook = async () => {
		if (!usn || !bookName || !bookID) {
			alert('Please fill in all the fields.');
			return;
		}

		try {
			loading = true;

			const response = await fetch('http://localhost:9090/books/return', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					usn: usn,
					book_name: bookName,
					book_id: bookID
				})
			});

			if (!response.ok) {
				const errorMessage = (await response.json()).message || 'Failed to return book';
				alert(`Error: ${errorMessage}`);
				return;
			}

			const data = await response.json();
			alert(data.message);

			// Refresh the list after returning the book
			await fetchBorrowedBooks();

			// Clear form fields
			usn = '';
			bookName = '';
			bookID = 0;
		} catch (error) {
			console.error('Failed to return book:', error);
		} finally {
			loading = false;
		}
	};

	onMount(() => {
		fetchBorrowedBooks();
	});
</script>

<!-- UI -->
<div class="min-h-screen bg-gradient-to-r from-blue-50 to-blue-100 p-10">
	<div class="mx-auto max-w-6xl">
		<h1 class="mb-8 text-5xl font-extrabold text-gray-800">📚 Return Borrowed Book</h1>

		<!-- Form for returning a book -->
		<div class="rounded-lg bg-white p-8 shadow-lg transition-transform hover:scale-105">
			<h2 class="mb-6 text-2xl font-semibold text-gray-700">Return a Book</h2>

			<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
				<div>
					<label class="block text-sm font-medium text-gray-700">USN</label>
					<input
						type="text"
						placeholder="Enter USN"
						bind:value={usn}
						class="mt-1 w-full rounded-md border-gray-300 p-3 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</div>

				<div>
					<label class="block text-sm font-medium text-gray-700">Book Name</label>
					<input
						type="text"
						placeholder="Enter Book Name"
						bind:value={bookName}
						class="mt-1 w-full rounded-md border-gray-300 p-3 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</div>

				<div>
					<label class="block text-sm font-medium text-gray-700">Book ID</label>
					<input
						type="number"
						placeholder="Enter Book ID"
						bind:value={bookID}
						class="mt-1 w-full rounded-md border-gray-300 p-3 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</div>
			</div>

			<button
				class="mt-6 w-full transform rounded-md bg-green-500 px-5 py-3 text-lg font-semibold text-white transition-transform hover:scale-105 hover:bg-green-600"
				on:click={returnBook}
				disabled={loading}
			>
				{loading ? 'Returning...' : 'Return Book'}
			</button>
		</div>

		<!-- Loading Spinner -->
		{#if loading}
			<div class="mt-10 flex justify-center">
				<div
					class="h-12 w-12 animate-spin rounded-full border-4 border-blue-400 border-t-transparent"
				></div>
			</div>
		{/if}

		<!-- Display Borrowed Books -->
		<h2 class="mt-12 mb-6 text-3xl font-bold text-gray-800">📚 Borrowed Books</h2>

		{#if borrowedBooks.length > 0}
			<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
				{#each borrowedBooks as book}
					<div
						class="transform rounded-lg bg-white p-8 shadow-lg transition-transform hover:scale-105"
					>
						<h3 class="text-2xl font-bold text-blue-600">{book.title}</h3>
						<p class="mt-2 text-gray-600">
							<span class="font-semibold">Author:</span>
							{book.author}
						</p>
						<p class="mt-2 text-gray-600">
							<span class="font-semibold">Book ID:</span>
							{book.book_id}
						</p>
						<p class="mt-2 text-gray-600">
							<span class="font-semibold">Borrowed Count:</span>
							{book.count}
						</p>
					</div>
				{/each}
			</div>
		{:else}
			<p class="mt-10 text-center text-gray-600">You have not borrowed any books.</p>
		{/if}
	</div>
</div>
