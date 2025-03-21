<script lang="ts">
	import { supabase } from '$lib/supabaseClient';

	let email = '';
	let password = '';
	let confirmPassword = '';
	let error = '';
	let success = '';

	async function handleRegister() {
		if (password !== confirmPassword) {
			error = 'Passwords do not match!';
			success = '';
			return;
		}

		try {
			const { data, error: authError } = await supabase.auth.signUp({
				email,
				password
			});

			if (authError) throw authError;

			success = '🎉 Registration successful! You can now log in.';
			error = '';
		} catch (e) {
			error = e.message;
			success = '';
		}
	}
</script>

<div class="container">
	<div class="card">
		<h2 class="heading">📚 Library Registration</h2>

		{#if error}
			<div class="message error">{error}</div>
		{/if}

		{#if success}
			<div class="message success">{success}</div>
		{/if}

		<form on:submit|preventDefault={handleRegister}>
			<div class="input-group">
				<label for="email">Email</label>
				<input type="email" id="email" bind:value={email} required />
			</div>

			<div class="input-group">
				<label for="password">Password</label>
				<input type="password" id="password" bind:value={password} required />
			</div>

			<div class="input-group">
				<label for="confirmPassword">Confirm Password</label>
				<input type="password" id="confirmPassword" bind:value={confirmPassword} required />
			</div>

			<button type="submit">🚀 Register</button>
		</form>

		<p class="mt-4 text-center text-sm text-gray-600">
			Already have an account?
			<a href="/login">Login here</a>
		</p>
	</div>
</div>

<style>
	/* Smooth transitions and user-friendly colors */
	.container {
		@apply flex min-h-screen items-center justify-center bg-gradient-to-br from-purple-600 to-blue-500;
	}

	.card {
		@apply w-full max-w-md rounded-xl bg-white p-8 shadow-lg transition-all duration-300 ease-in-out hover:shadow-2xl;
	}

	.heading {
		@apply mb-6 text-center text-3xl font-extrabold text-gray-800;
	}

	.input-group {
		@apply mb-4;
	}

	label {
		@apply block text-sm font-medium text-gray-700;
	}

	input {
		@apply mt-1 w-full rounded-lg border border-gray-300 p-3 text-gray-700 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:outline-none;
	}

	button {
		@apply w-full rounded-lg bg-blue-600 px-4 py-2 text-white transition-all duration-300 hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:outline-none;
	}

	.message {
		@apply mb-4 rounded-lg px-4 py-3 text-sm;
	}

	.success {
		@apply border border-green-400 bg-green-100 text-green-700;
	}

	.error {
		@apply border border-red-400 bg-red-100 text-red-700;
	}

	a {
		@apply text-blue-500 underline hover:text-blue-700;
	}
</style>
