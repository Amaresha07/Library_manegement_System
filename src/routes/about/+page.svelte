<script lang="ts">
	import { onMount } from 'svelte';
	let currentIndex = 0;

	const quotes = [
		'📚 A library is a portal to countless worlds, where every page is a new adventure.',
		'🌿 In the embrace of books, time becomes irrelevant and imagination infinite.',
		'✨ Libraries: where you find magic hidden between the covers.',
		'☕ A book, a quiet corner, and a cup of coffee—the perfect escape.'
	];

	const changeQuote = () => {
		currentIndex = (currentIndex + 1) % quotes.length;
	};

	onMount(() => {
		const interval = setInterval(changeQuote, 5000); // Change quote every 5 seconds
		return () => clearInterval(interval);
	});
</script>

<div class="container">
	<div class="library-card">
		<h1>📚 Welcome to the Library of Imagination ✨</h1>

		<p>
			Step into a world where time stands still, and every book is a new journey. From **ancient
			scrolls** to **futuristic tales**, the library holds the keys to countless realms. Find a
			corner, grab a book, and **let your imagination wander**.
		</p>

		<p class="quote">{quotes[currentIndex]}</p>

		<a href="/explore" class="explore-btn">Explore Books 📖</a>
	</div>
</div>

<style>
	body {
		margin: 0;
		padding: 0;
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		background: #f5f5f5;
		color: #333;
	}

	.container {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		background: linear-gradient(to right, #6a11cb, #2575fc);
		color: #fff;
		padding: 2rem;
		text-align: center;
		animation: fadeIn 1s ease-in-out;
	}

	.library-card {
		max-width: 900px;
		width: 90%;
		background: #ffffff;
		color: #333;
		padding: 2rem;
		border-radius: 12px;
		box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
		transition: transform 0.3s ease;
	}

	.library-card:hover {
		transform: translateY(-10px);
	}

	h1 {
		font-size: 3rem;
		margin-bottom: 1rem;
		color: #6a11cb;
	}

	p {
		font-size: 1.2rem;
		line-height: 1.8;
	}

	.quote {
		font-style: italic;
		color: #555;
		font-size: 1.5rem;
		margin: 2rem 0;
		transition: opacity 0.5s;
	}

	.quote::before,
	.quote::after {
		content: '“';
		color: #2575fc;
		font-size: 2rem;
	}

	.explore-btn {
		display: inline-block;
		background: #2575fc;
		color: #fff;
		padding: 0.8rem 2rem;
		font-size: 1rem;
		border-radius: 8px;
		text-decoration: none;
		transition: background 0.3s;
		margin-top: 1.5rem;
	}

	.explore-btn:hover {
		background: #6a11cb;
	}

	@keyframes fadeIn {
		0% {
			opacity: 0;
		}
		100% {
			opacity: 1;
		}
	}
</style>
