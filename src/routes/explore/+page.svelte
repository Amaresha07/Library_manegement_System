<script lang="ts">
	let selectedBook: {
		series: string;
		volume: string;
		language: string;
		category: string;
		title: string;
		cover: string;
		description: string;
		file: string;
	} | null = null;

	let books = [
		// Magazines (English)
		{
			series: 'Science Monthly',
			volume: 'Vol. 1',
			language: 'English',
			category: 'Magazines',
			title: 'The Future of AI',
			cover: 'https://m.media-amazon.com/images/I/61+6j4yB9OL.jpg',
			description: 'Exploring the latest advancements in artificial intelligence.',
			file: 'https://www.example.com/ai-magazine'
		},
		{
			series: 'Science Monthly',
			volume: 'Vol. 2',
			language: 'English',
			category: 'Magazines',
			title: 'Space Exploration',
			cover: 'https://m.media-amazon.com/images/I/71QKQ9mwV7L.jpg',
			description: 'The latest updates on space travel and astronomy.',
			file: 'https://www.example.com/space-magazine'
		},
		{
			series: 'Science Monthly',
			volume: 'Vol. 2',
			language: 'English',
			category: 'Magazines',
			title: 'Space Exploration',
			cover: 'https://m.media-amazon.com/images/I/71QKQ9mwV7L.jpg',
			description: 'The latest updates on space travel and astronomy.',
			file: 'https://www.example.com/space-magazine'
		},
		{
			series: 'Science Monthly',
			volume: 'Vol. 2',
			language: 'English',
			category: 'Magazines',
			title: 'Space Exploration',
			cover: 'https://m.media-amazon.com/images/I/71QKQ9mwV7L.jpg',
			description: 'The latest updates on space travel and astronomy.',
			file: 'https://www.example.com/space-magazine'
		},
		{
			series: 'Science Monthly',
			volume: 'Vol. 2',
			language: 'English',
			category: 'Magazines',
			title: 'Space Exploration',
			cover: 'https://m.media-amazon.com/images/I/71QKQ9mwV7L.jpg',
			description: 'The latest updates on space travel and astronomy.',
			file: 'https://www.example.com/space-magazine'
		},
		{
			series: 'ವಿಜ್ಞಾನ ಮಾಸಪತ್ರಿಕೆ',
			volume: 'ಸಂಪುಟ 1',
			language: 'Kannada',
			category: 'Magazines',
			title: 'ಕೃತಕ ಬುದ್ಧಿಮತ್ತೆಯ ಭವಿಷ್ಯ',
			cover: 'https://m.media-amazon.com/images/I/71QKQ9mwV7L.jpg',
			description: 'ಕೃತಕ ಬುದ್ಧಿಮತ್ತೆಯ ಮುಂದಿನ ಹಂತಗಳನ್ನು ಅನ್ವೇಷಣೆ.',
			file: 'https://www.example.com/kannada-ai'
		},
		{
			series: 'ವಿಜ್ಞಾನ ಮಾಸಪತ್ರಿಕೆ',
			volume: 'ಸಂಪುಟ 1',
			language: 'Kannada',
			category: 'Magazines',
			title: 'ಕೃತಕ ಬುದ್ಧಿಮತ್ತೆಯ ಭವಿಷ್ಯ',
			cover: 'https://m.media-amazon.com/images/I/71QKQ9mwV7L.jpg',
			description: 'ಕೃತಕ ಬುದ್ಧಿಮತ್ತೆಯ ಮುಂದಿನ ಹಂತಗಳನ್ನು ಅನ್ವೇಷಣೆ.',
			file: 'https://www.example.com/kannada-ai'
		},

		// Magazines (Kannada)
		{
			series: 'ವಿಜ್ಞಾನ ಮಾಸಪತ್ರಿಕೆ',
			volume: 'ಸಂಪುಟ 1',
			language: 'Kannada',
			category: 'Magazines',
			title: 'ಕೃತಕ ಬುದ್ಧಿಮತ್ತೆಯ ಭವಿಷ್ಯ',
			cover: 'https://m.media-amazon.com/images/I/71QKQ9mwV7L.jpg',
			description: 'ಕೃತಕ ಬುದ್ಧಿಮತ್ತೆಯ ಮುಂದಿನ ಹಂತಗಳನ್ನು ಅನ್ವೇಷಣೆ.',
			file: 'https://www.example.com/kannada-ai'
		},

		// Storybooks (English)
		{
			series: 'Classic Tales',
			volume: 'Vol. 1',
			language: 'English',
			category: 'Storybooks',
			title: 'Alice in Wonderland',
			cover: 'https://m.media-amazon.com/images/I/81iqZ2HHD-L.jpg',
			description: 'A fantasy novel by Lewis Carroll about a girl in a magical world.',
			file: 'https://www.gutenberg.org/files/11/11-h/11-h.htm'
		},

		// Storybooks (Kannada)
		{
			series: 'ಜಾನಪದ ಕಥೆಗಳು',
			volume: 'ಸಂಪುಟ 1',
			language: 'Kannada',
			category: 'Storybooks',
			title: 'ತಿಂಗಳು ಮತ್ತು ಚಂದ್ರ',
			cover: 'https://m.media-amazon.com/images/I/51K84pomCRL.jpg',
			description: 'ಕನ್ನಡ ಜಾನಪದ ಕಥೆಗಳ ಸಂಗ್ರಹ.',
			file: 'https://www.example.com/kannada-story'
		}
	];

	function openBook(book) {
		selectedBook = book;
	}

	function closeBook() {
		selectedBook = null;
	}
</script>

<!-- Library UI -->
<div class="container">
	<h1>📚 Digital Library</h1>
	<p>Explore Magazines and Storybooks in English & Kannada</p>

	<!-- Book Categories -->
	{#each ['Magazines', 'Storybooks'] as category}
		<h2>{category}</h2>
		<div class="book-grid">
			{#each books.filter((book) => book.category === category) as book}
				<div class="book-card" on:click={() => openBook(book)}>
					<img src={book.cover} alt={book.title} />
					<p class="book-title">{book.title}</p>
					<p class="book-series">{book.series} - {book.volume}</p>
				</div>
			{/each}
		</div>
	{/each}
</div>

<!-- Book Reading Modal -->
{#if selectedBook}
	<div class="modal">
		<div class="modal-content">
			<button class="close-btn" on:click={closeBook}>✖</button>
			<img src={selectedBook.cover} alt={selectedBook.title} class="book-cover" />
			<h2>{selectedBook.title}</h2>
			<h3>{selectedBook.series} - {selectedBook.volume}</h3>
			<p class="description">{selectedBook.description}</p>
			<a href={selectedBook.file} target="_blank" class="read-btn">📖 Read Now</a>
		</div>
	</div>
{/if}

<style>
	/* Page Styling */
	.container {
		text-align: center;
		padding: 20px;
		background: linear-gradient(to right, #fafafa, #f4f4f4);
		min-height: 100vh;
		font-family: Arial, sans-serif;
	}

	h1 {
		font-size: 2.5rem;
		color: #333;
	}

	h2 {
		margin-top: 30px;
		color: #444;
		border-bottom: 3px solid #ddd;
		display: inline-block;
		padding-bottom: 5px;
	}

	h3 {
		color: #666;
		margin-top: 10px;
	}

	/* Bookshelf Grid */
	.book-grid {
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
		gap: 20px;
		padding: 20px;
	}

	.book-card {
		width: 200px;
		background: white;
		padding: 10px;
		border-radius: 10px;
		cursor: pointer;
		transition:
			transform 0.3s,
			box-shadow 0.3s;
		box-shadow: 0 5px 10px rgba(0, 0, 0, 0.2);
		text-align: center;
	}

	.book-card img {
		width: 100%;
		height: 270px;
		object-fit: cover;
		border-radius: 10px;
	}

	.book-card:hover {
		transform: scale(1.05);
		box-shadow: 0 10px 20px rgba(0, 0, 0, 0.3);
	}

	.book-title {
		margin-top: 10px;
		font-weight: bold;
		color: #333;
	}

	.book-series {
		font-size: 0.9rem;
		color: #666;
	}

	/* Modal Styling */
	.modal {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0, 0, 0, 0.7);
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.modal-content {
		background: white;
		width: 60%;
		max-width: 800px;
		padding: 30px;
		border-radius: 12px;
		box-shadow: 0 15px 25px rgba(0, 0, 0, 0.5);
		text-align: center;
	}

	.book-cover {
		width: 150px;
		border-radius: 10px;
		margin-bottom: 10px;
	}

	.close-btn {
		position: absolute;
		top: 10px;
		right: 15px;
		font-size: 1.5rem;
		background: none;
		border: none;
		cursor: pointer;
		color: red;
	}

	.description {
		margin: 10px 0;
		color: #555;
		font-size: 1rem;
	}

	.read-btn {
		display: inline-block;
		padding: 10px 20px;
		background: #3498db;
		color: white;
		text-decoration: none;
		font-weight: bold;
		border-radius: 5px;
		margin-top: 15px;
		transition: background 0.3s;
	}

	.read-btn:hover {
		background: #2980b9;
		transform: scale(1.05);
		transition: 0.3s;
	}
</style>
