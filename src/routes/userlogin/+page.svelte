<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	let Usn = '';
	let password = '';
	let errorMessage = '';
	let loading = false;

	const login = async () => {
		if (!Usn || !password) {
			errorMessage = 'Please fill in all fields';
			return;
		}

		loading = true;

		try {
			const response = await fetch('http://localhost:9090/userlogin', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ Usn, password })
			});

			if (!response.ok) {
				errorMessage = 'Invalid ID or password';
				loading = false;
				return;
			}

			const data = await response.json();
			localStorage.setItem('token', data.token);

			alert('✅ Login successful!');
			goto('/borrow'); // Redirect to books page
		} catch (error) {
			console.error('Failed to log in:', error);
			errorMessage = '❌ Failed to connect to server';
		} finally {
			loading = false;
		}
	};
</script>

<div class="bg-gradient flex min-h-screen items-center justify-center">
	<div class="card w-full max-w-md rounded-lg p-8">
		<h1 class="mb-6 text-center text-3xl font-bold text-gray-800">🔐 Welcome Back!</h1>

		{#if errorMessage}
			<div class="mb-4 rounded-md bg-red-100 p-4 text-red-600">
				{errorMessage}
			</div>
		{/if}

		<div class="mb-4">
			<label class="block text-sm font-medium text-gray-700">USN</label>
			<input
				type="text"
				bind:value={Usn}
				placeholder="Enter your USN"
				class="mt-1 w-full rounded-md border border-gray-300 p-3 shadow-sm transition focus:border-blue-500 focus:ring focus:ring-blue-200"
			/>
		</div>

		<div class="mb-4">
			<label class="block text-sm font-medium text-gray-700">Password</label>
			<input
				type="password"
				bind:value={password}
				placeholder="Enter your password"
				class="mt-1 w-full rounded-md border border-gray-300 p-3 shadow-sm transition focus:border-blue-500 focus:ring focus:ring-blue-200"
			/>
		</div>

		<!-- Login Button with loading animation -->
		<button
			on:click={login}
			class="flex w-full items-center justify-center rounded-md bg-blue-500 py-2 text-white transition hover:bg-blue-600"
			disabled={loading}
		>
			{#if loading}
				<div class="spinner mr-2"></div>
				<span>Logging in...</span>
			{:else}
				<span>Login</span>
			{/if}
		</button>

		<p class="mt-4 text-center text-gray-600">
			Don't have an account?
			<a href="/userregister" class="text-blue-500 hover:underline">Register here</a>
		</p>
	</div>
</div>

<style>
	/* Background gradient */
	.bg-gradient {
		background: linear-gradient(135deg, #667eea, #764ba2);
	}

	/* Card styling */
	.card {
		background: white;
		box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
		transition:
			transform 0.3s ease,
			box-shadow 0.3s ease;
	}
	.card:hover {
		transform: translateY(-5px);
		box-shadow: 0 15px 25px rgba(0, 0, 0, 0.15);
	}

	/* Button loading animation */
	.spinner {
		border: 4px solid rgba(0, 0, 0, 0.1);
		border-left-color: #4f46e5;
		animation: spin 1s linear infinite;
		border-radius: 50%;
		width: 20px;
		height: 20px;
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}
</style>
