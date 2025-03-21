<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let id: number | null = null;
	let title = '';
	let author = '';
	let year = 0;
	let count = 0;
	let books = [];

	const fetchBooks = async () => {
		try {
			const response = await fetch('http://localhost:9090/books');
			if (!response.ok) {
				throw new Error('Failed to fetch books');
			}

			const data = await response.json();
			books = [...data];
		} catch (error) {
			console.error('Error fetching books:', error);
		}
	};

	onMount(() => {
		fetchBooks();
	});

	const updateBook = async () => {
		try {
			if (!id || id === 0) {
				alert('Please enter a valid Book ID');
				return;
			}

			const response = await fetch(`http://localhost:9090/update/${id}`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					title,
					author,
					count,
					year
				})
			});

			if (!response.ok) {
				const errorText = await response.text();
				try {
					const errorData = JSON.parse(errorText);
					alert(`Failed to update book: ${errorData.message}`);
				} catch {
					alert(`Failed to update book: ${errorText}`);
				}
				return;
			}

			alert(`Book Updated: ${title} by ${author} (${count} copies)`);
			goto('/books');
		} catch (error) {
			console.error('Error updating book:', error);
			alert('Failed to update book');
		}
	};
</script>

<!-- Real-world attractive UI -->
<div class="library-bg">
	<div class="form-container">
		<h1>📚 Update Book Details</h1>

		<form on:submit|preventDefault={updateBook}>
			<div class="input-group">
				<label>Book ID</label>
				<input type="number" placeholder="Enter Book ID" bind:value={id} required />
			</div>

			<div class="input-group">
				<label>Title</label>
				<input type="text" placeholder="Enter Book Title" bind:value={title} required />
			</div>

			<div class="input-group">
				<label>Author</label>
				<input type="text" placeholder="Enter Author Name" bind:value={author} required />
			</div>

			<div class="input-group">
				<label>Copies</label>
				<input type="number" placeholder="Number of Copies" bind:value={count} required />
			</div>

			<div class="input-group">
				<label>Year</label>
				<input type="number" placeholder="Publication Year" bind:value={year} required />
			</div>

			<button type="submit" class="update-btn">✅ Update Book</button>
		</form>

		<button class="back-btn" on:click={() => goto('/books')}>🔙 Back to Book List</button>
	</div>
</div>

<style>
	/* Background and layout */
	.library-bg {
		min-height: 100vh;
		background: linear-gradient(135deg, #f3f4f6, #d1c4e9);
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 20px;
	}

	.form-container {
		background: #ffffff;
		box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
		border-radius: 12px;
		padding: 40px;
		width: 100%;
		max-width: 600px;
		text-align: center;
		transition:
			transform 0.3s,
			box-shadow 0.3s;
	}

	.form-container:hover {
		transform: translateY(-10px);
		box-shadow: 0 15px 35px rgba(0, 0, 0, 0.3);
	}

	h1 {
		font-size: 2.5rem;
		color: #4a4a4a;
		margin-bottom: 20px;
	}

	/* Input fields */
	.input-group {
		margin-bottom: 20px;
		text-align: left;
	}

	.input-group label {
		display: block;
		margin-bottom: 8px;
		font-size: 1rem;
		color: #555;
	}

	.input-group input {
		width: 100%;
		padding: 12px 15px;
		font-size: 1rem;
		border: 1px solid #ccc;
		border-radius: 8px;
		transition: border 0.3s;
	}

	.input-group input:focus {
		border-color: #6200ea;
		outline: none;
	}

	/* Buttons */
	.update-btn {
		background: #4caf50;
		color: white;
		border: none;
		padding: 12px 30px;
		font-size: 1.1rem;
		border-radius: 8px;
		cursor: pointer;
		transition: 0.3s;
		width: 100%;
		margin-top: 20px;
	}

	.update-btn:hover {
		background: #45a049;
	}

	.back-btn {
		background: transparent;
		color: #6200ea;
		border: 2px solid #6200ea;
		padding: 12px 30px;
		font-size: 1.1rem;
		border-radius: 8px;
		cursor: pointer;
		transition: 0.3s;
		width: 100%;
		margin-top: 15px;
	}

	.back-btn:hover {
		background: #6200ea;
		color: white;
	}

	/* Responsiveness */
	@media (max-width: 768px) {
		.form-container {
			width: 90%;
			padding: 30px;
		}
	}
</style>
