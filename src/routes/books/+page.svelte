<script lang="ts">
	import { onMount } from 'svelte';

	let books = [];
	let loading = false;
	let errorMsg = '';
	let title = '';
	let author = '';

	// Fetch all books initially
	const fetchBooks = async () => {
		loading = true;
		errorMsg = '';

		try {
			const response = await fetch('http://localhost:9090/books');
			if (!response.ok) throw new Error('Failed to fetch books');

			books = await response.json();
		} catch (error) {
			errorMsg = 'Error fetching books. Please try again!';
			console.error(error);
		} finally {
			loading = false;
		}
	};

	// Search books based on title and author
	const searchBooks = async () => {
		loading = true;
		errorMsg = '';

		try {
			const response = await fetch('http://localhost:9090/books/search', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ title, author })
			});

			if (!response.ok) {
				throw new Error('Failed to fetch books');
			}

			books = await response.json();
		} catch (error) {
			errorMsg = 'Error fetching books. Please try again!';
			console.error(error);
		} finally {
			loading = false;
		}
	};

	onMount(() => {
		fetchBooks(); // Load all books initially
	});
</script>

<!-- Search Form -->
<div class="container">
	<header>
		<h1>📚 Library Books</h1>
		<p>Search and explore your favorite books!</p>
	</header>

	<div class="search-bar">
		<input type="text" placeholder="Enter Title" bind:value={title} class="search-input" />
		<input type="text" placeholder="Enter Author" bind:value={author} class="search-input" />
		<button on:click={searchBooks} class="btn search-btn">🔍 Search</button>
		<button on:click={fetchBooks} class="btn reset-btn">🔄 Reset</button>
	</div>

	<!-- Book List -->
	{#if loading}
		<div class="spinner"></div>
	{:else if errorMsg}
		<p class="error-message">{errorMsg}</p>
	{:else if books.length > 0}
		<div class="book-list">
			{#each books as book (book.book_id)}
				<div class="book-card">
					<div class="book-cover">
						<img src={book.cover_url} alt={book.title} class="book-thumbnail" />
					</div>
					<div class="book-info">
						<h2 class="book-title">{book.title}</h2>
						<p class="book-author">by {book.author}</p>
						<p class="book-year">Published: {book.year}</p>
						<p class="book-count">Available: {book.count} copies</p>
					</div>
				</div>
			{/each}
		</div>
	{:else}
		<p class="no-books">No books available at the moment. 📚</p>
	{/if}
</div>

<!-- Styles -->
<style>
	/* Overall container */
	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 40px 20px;
		background: #f4f4f4;
		border-radius: 12px;
		box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
	}

	header {
		text-align: center;
		margin-bottom: 30px;
	}

	h1 {
		font-size: 3rem;
		color: #2c5282;
	}

	p {
		color: #4a5568;
		font-size: 1.2rem;
	}

	/* Search bar styling */
	.search-bar {
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
		gap: 15px;
		margin-bottom: 40px;
	}

	.search-input {
		width: 280px;
		padding: 12px 18px;
		border: 2px solid #ccc;
		border-radius: 8px;
		font-size: 1rem;
		transition: border 0.3s;
	}

	.search-input:focus {
		border-color: #2c5282;
		outline: none;
	}

	/* Buttons styling */
	.btn {
		padding: 12px 30px;
		border: none;
		border-radius: 8px;
		font-size: 1rem;
		cursor: pointer;
		transition:
			background 0.3s,
			transform 0.2s;
		color: #fff;
	}

	.search-btn {
		background: linear-gradient(135deg, #2c5282, #3182ce);
	}

	.reset-btn {
		background: linear-gradient(135deg, #48bb78, #38a169);
	}

	.btn:hover {
		transform: translateY(-3px);
	}

	/* Book list styling */
	.book-list {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
		gap: 30px;
	}

	.book-card {
		background: #ffffff;
		border-radius: 12px;
		box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
		overflow: hidden;
		transition: transform 0.3s ease-in-out;
		cursor: pointer;
	}

	.book-card:hover {
		transform: translateY(-10px);
		box-shadow: 0 15px 35px rgba(0, 0, 0, 0.15);
	}

	.book-cover {
		height: 200px;
		background: #edf2f7;
		overflow: hidden;
	}

	.book-thumbnail {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 0.3s ease;
	}

	.book-card:hover .book-thumbnail {
		transform: scale(1.05);
	}

	.book-info {
		padding: 20px;
	}

	.book-title {
		font-size: 1.8rem;
		font-weight: bold;
		color: #2d3748;
	}

	.book-author {
		font-size: 1.2rem;
		color: #4a5568;
		margin: 10px 0;
	}

	.book-year,
	.book-count {
		font-size: 1rem;
		color: #718096;
	}

	/* No books and error message */
	.no-books,
	.error-message {
		text-align: center;
		font-size: 1.5rem;
		color: #e53e3e;
	}

	/* Spinner styling */
	.spinner {
		border: 5px solid #f3f3f3;
		border-top: 5px solid #3498db;
		border-radius: 50%;
		width: 50px;
		height: 50px;
		animation: spin 1s linear infinite;
		margin: 20px auto;
	}

	@keyframes spin {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}

	/* Mobile responsive */
	@media (max-width: 768px) {
		.search-bar {
			flex-direction: column;
			align-items: center;
		}

		.book-list {
			grid-template-columns: 1fr;
		}
	}
</style>
