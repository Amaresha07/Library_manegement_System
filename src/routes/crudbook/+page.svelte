<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let books = [];

	const fetchBooks = async () => {
		try {
			const response = await fetch('http://localhost:9090/books');
			if (!response.ok) {
				throw new Error('Failed to fetch books');
			}

			const data = await response.json();
			console.log('Fetched books:', data);

			books = [...data]; // Assign the fetched books
		} catch (error) {
			console.error('Error fetching books:', error);
		}
	};

	onMount(() => {
		fetchBooks();
	});

	const navigateToAdd = () => {
		goto('/addbook');
	};

	const navigateToEdit = (id: number) => {
		goto(`/editbook/`);
	};

	const deleteBook = async (id: number) => {
		if (confirm(`Are you sure you want to delete book ID ${id}?`)) {
			try {
				const response = await fetch(`http://localhost:9090/books/delete/${id}`, {
					method: 'DELETE'
				});

				if (!response.ok) {
					throw new Error('Failed to delete book');
				}

				alert(`Book with ID ${id} deleted`);
				fetchBooks(); // Refresh the list after deletion
			} catch (error) {
				console.error('Error deleting book:', error);
			}
		}
	};
</script>

<div class="library-page">
	<div class="container">
		<div class="header">
			<h1>📚 Library Books</h1>
			<button class="add-btn" on:click={navigateToAdd}>➕ Add Book</button>
		</div>

		<div class="books-grid">
			{#each books as book (book.book_id)}
				<div class="book-card">
					<div class="card-content">
						<h2 class="card-header">{book.title}</h2>
						<p class="card-meta">👤 {book.author}</p>
						<p class="card-meta">📅 Year: {book.year}</p>
						<p class="card-meta">📚 Copies: {book.count}</p>
					</div>

					<div class="actions">
						<button class="update-btn" on:click={() => navigateToEdit(book.book_id)}>
							✏️ Edit
						</button>
						<button class="delete-btn" on:click={() => deleteBook(book.book_id)}>
							🗑️ Delete
						</button>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

<style>
	.library-page {
		min-height: 100vh;
		background: linear-gradient(135deg, #f1f1f1, #d1c4e9);
		padding: 40px;
	}

	.container {
		max-width: 1200px;
		margin: 0 auto;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 40px;
	}

	h1 {
		font-size: 3rem;
		color: #4a4a4a;
	}

	.add-btn {
		background: #4caf50;
		color: white;
		padding: 10px 20px;
		border: none;
		border-radius: 8px;
		cursor: pointer;
		transition: 0.3s;
	}

	.add-btn:hover {
		background: #45a049;
	}

	.books-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 30px;
	}

	.book-card {
		background: #ffffff;
		border-radius: 12px;
		box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
		overflow: hidden;
		transition:
			transform 0.3s,
			box-shadow 0.3s;
	}

	.book-card:hover {
		transform: translateY(-10px);
		box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
	}

	.card-content {
		padding: 20px;
	}

	.card-header {
		font-size: 1.5rem;
		font-weight: bold;
		color: #333;
		margin-bottom: 10px;
	}

	.card-meta {
		color: #777;
		font-size: 1rem;
		margin-bottom: 15px;
	}

	.actions {
		display: flex;
		justify-content: space-between;
		padding: 15px;
		background: #f9f9f9;
	}

	.actions button {
		padding: 10px 15px;
		border: none;
		cursor: pointer;
		font-size: 1rem;
		border-radius: 6px;
		transition: 0.3s;
	}

	.update-btn {
		background: #28a745;
		color: white;
	}

	.update-btn:hover {
		background: #218838;
	}

	.delete-btn {
		background: #dc3545;
		color: white;
	}

	.delete-btn:hover {
		background: #c82333;
	}
</style>
