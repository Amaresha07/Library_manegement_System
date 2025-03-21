<script lang="ts">
	import { onMount } from 'svelte';

	let books = [];

	// Fetch all books
	const fetchBooks = async () => {
		try {
			const response = await fetch('http://localhost:9090/books');
			if (!response.ok) throw new Error('Failed to fetch books');

			books = await response.json();
		} catch (error) {
			console.error('Error fetching books:', error);
		}
	};

	// Borrow book function
	const borrowBook = async (bookId) => {
		const token = localStorage.getItem('token');

		if (!token) {
			alert('You must be logged in!');
			return;
		}

		try {
			const response = await fetch(`http://localhost:9090/borrow/${bookId}`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				}
			});

			if (!response.ok) {
				const errorMessage = await response.text();
				alert(`Error: ${errorMessage}`);
				return;
			}

			const data = await response.json();
			alert(data.message);
			await fetchBooks(); // Refresh book list after borrowing
		} catch (error) {
			console.error('Failed to borrow book:', error);
		}
	};

	onMount(() => {
		fetchBooks();
	});
</script>

<!-- Navbar -->
<div class="navbar">
	<a href="/">🏠 Home</a>
	<a href="/books">📚 Books</a>
	<a href="/my-books">📖 My Borrowed Books</a>
	<a href="/profile">👤 Profile</a>
</div>

<!-- Main Container -->
<div class="container">
	<h1 class="library-title">📚 Library Book Collection</h1>

	{#if books.length > 0}
		<div class="book-grid">
			{#each books as book}
				<div class="book-card">
					<!-- Book Cover Image -->
					<img src="https://via.placeholder.com/300x200?text={book.title}" alt={book.title} />

					<!-- Book Details -->
					<div class="book-details">
						<h2 class="book-title">{book.title}</h2>
						<p class="book-author">Author: <strong>{book.author}</strong></p>
						<p>Available: <span class="book-count">{book.count}</span></p>

						<!-- Borrow Button -->
						<button
							class="borrow-btn"
							on:click={() => borrowBook(book.book_id)}
							disabled={book.count <= 0}
						>
							{book.count > 0 ? 'Borrow' : 'Out of Stock'}
						</button>
					</div>
				</div>
			{/each}
		</div>
	{:else}
		<p class="no-books">No books available at the moment. 📚</p>
	{/if}
</div>

<style>
	.navbar {
		background: #2c3e50;
		color: white;
		padding: 20px 50px;
		display: flex;
		justify-content: space-between;
		align-items: center;
		box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
	}
	.navbar a {
		color: white;
		text-decoration: none;
		font-size: 18px;
		margin: 0 15px;
		transition: color 0.3s;
	}
	.navbar a:hover {
		color: #f39c12;
	}
	.container {
		padding: 50px;
		background: #f4f4f9;
		min-height: 100vh;
	}
	.library-title {
		text-align: center;
		font-size: 48px;
		color: #2c3e50;
		margin-bottom: 30px;
	}
	.book-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 30px;
	}
	.book-card {
		background: #ffffff;
		border-radius: 10px;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
		overflow: hidden;
		transition: transform 0.3s;
	}
	.book-card:hover {
		transform: translateY(-10px);
	}
	.book-card img {
		width: 100%;
		height: 200px;
		object-fit: cover;
	}
	.book-details {
		padding: 20px;
	}
	.book-title {
		font-size: 24px;
		color: #34495e;
		margin-bottom: 10px;
	}
	.book-author {
		color: #7f8c8d;
		margin-bottom: 10px;
	}
	.book-count {
		font-weight: bold;
		color: #27ae60;
	}
	.borrow-btn {
		display: inline-block;
		width: 100%;
		padding: 12px 0;
		text-align: center;
		background: #3498db;
		color: white;
		font-size: 18px;
		border: none;
		cursor: pointer;
		transition: background 0.3s;
	}
	.borrow-btn:hover {
		background: #2980b9;
	}
	.no-books {
		text-align: center;
		color: #7f8c8d;
		font-size: 22px;
	}
</style>
