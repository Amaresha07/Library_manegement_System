<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { supabase } from '$lib/supabaseClient';

	let user;

	// Check if the user is authenticated on mount
	onMount(async () => {
		try {
			const {
				data: { user: currentUser }
			} = await supabase.auth.getUser();
			user = currentUser;
		} catch (error) {
			console.error('Error fetching user:', error);
		}
	});

	// Navigation handlers
	function navigateToUser() {
		goto(user ? '/userlogin' : '/userregister');
	}

	function navigateToReturn() {
		goto('/return');
	}

	function navigateToAdmin() {
		goto('/admin/login');
	}
</script>

<div class="min-h-screen bg-gray-100 text-gray-900">
	<!-- Hero Section -->
	<div class="hero flex items-center justify-center py-20 text-white">
		<div class="text-center">
			<h1 class="text-5xl leading-tight font-extrabold drop-shadow-lg">
				📚 Welcome to Library Management System
			</h1>
			<p class="mt-4 text-lg opacity-90">Effortlessly manage books, users, and returns.</p>
			<div class="mt-6 space-x-4">
				<a
					href="/books"
					class="rounded-lg bg-green-500 px-6 py-3 text-white transition hover:bg-green-600"
					>View Books</a
				>
				<button
					on:click={navigateToUser}
					class="rounded-lg bg-blue-500 px-6 py-3 text-white transition hover:bg-blue-600"
				>
					User Login
				</button>
			</div>
		</div>
	</div>

	<!-- Main Section -->
	<div class="container mx-auto px-6 py-12">
		<div class="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
			<!-- User Section -->
			<div class="card transform rounded-lg bg-white p-8 shadow-md transition hover:scale-105">
				<h2 class="text-3xl font-bold text-blue-600">👤 For Users</h2>
				<p class="mt-4 text-gray-700">Search, borrow, and track books easily.</p>
				<ul class="mt-4 space-y-2 text-gray-600">
					<li>✅ Browse available books</li>
					<li>✅ Search by title or author</li>
					<li>✅ Borrow books online</li>
					<li>✅ Track your borrowed books</li>
				</ul>
				<div class="mt-6 space-y-4">
					<button
						on:click={navigateToUser}
						class="w-full rounded-md bg-indigo-600 px-6 py-3 text-white transition hover:bg-indigo-700"
					>
						User Login
					</button>
					<button
						on:click={navigateToReturn}
						class="w-full rounded-md bg-green-600 px-6 py-3 text-white transition hover:bg-green-700"
					>
						Return Book
					</button>
				</div>
			</div>

			<!-- Admin Section -->
			<div class="card transform rounded-lg bg-white p-8 shadow-md transition hover:scale-105">
				<h2 class="text-3xl font-bold text-red-600">🛠️ For Admins</h2>
				<p class="mt-4 text-gray-700">Manage library operations efficiently.</p>
				<ul class="mt-4 space-y-2 text-gray-600">
					<li>✅ Manage book inventory</li>
					<li>✅ Track borrowed books</li>
					<li>✅ Handle user requests</li>
					<li>✅ Generate reports</li>
				</ul>
				<div class="mt-6">
					<button
						on:click={navigateToAdmin}
						class="w-full rounded-md bg-red-600 px-6 py-3 text-white transition hover:bg-red-700"
					>
						Admin Login
					</button>
				</div>
			</div>

			<!-- Info Section -->
			<div class="card transform rounded-lg bg-white p-8 shadow-md transition hover:scale-105">
				<h2 class="text-3xl font-bold text-gray-800">✨ Why Choose Us?</h2>
				<p class="mt-4 text-gray-700">
					Our system offers a seamless library experience with real-time updates and secure book
					management.
				</p>
				<ul class="mt-4 space-y-2 text-gray-600">
					<li>⚡ Real-time book availability</li>
					<li>🔐 Secure login and data management</li>
					<li>📊 Easy-to-use admin dashboard</li>
				</ul>
				<a
					href="/about"
					class="mt-6 inline-block rounded-md bg-gray-800 px-6 py-3 text-white transition hover:bg-gray-900"
				>
					Learn More →
				</a>
			</div>
		</div>
	</div>

	<!-- Footer -->
	<footer class="mt-12 bg-gray-800 py-6 text-center text-white">
		<p>&copy; 2025 Library Management System. All rights reserved.</p>
	</footer>
</div>

<style>
	/* Custom styles for extra polish */
	.hero {
		background: linear-gradient(135deg, #667eea, #764ba2);
	}
	.card:hover {
		transform: translateY(-10px);
		box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
	}
</style>
