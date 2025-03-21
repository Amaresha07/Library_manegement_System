<script lang="ts">
	import { goto } from '$app/navigation';
	import { fade, fly } from 'svelte/transition';

	let Usn = '';
	let name = '';
	let password = '';
	let errorMessage = '';
	let successMessage = '';

	const register = async () => {
		try {
			const response = await fetch('http://localhost:9090/userregister', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ Usn, name, password })
			});

			if (!response.ok) {
				errorMessage = '⚠️ User ID already exists or invalid data';
				return;
			}

			successMessage = '✅ Registration successful! Redirecting to login...';
			setTimeout(() => goto('/userlogin'), 2000); // Redirect after success
		} catch (error) {
			console.error('Failed to register:', error);
			errorMessage = '❌ Failed to connect to server';
		}
	};
</script>

<div class="bg-gradient">
	<div class="card" in:fly={{ y: 50, duration: 500 }} out:fade={{ duration: 300 }}>
		<h1 class="mb-6 text-center text-3xl font-bold">📝 Create Your Account</h1>

		{#if errorMessage}
			<p class="message error">{errorMessage}</p>
		{/if}

		{#if successMessage}
			<p class="message success">{successMessage}</p>
		{/if}

		<div>
			<label class="block font-medium text-gray-700">📌 USN</label>
			<input type="text" bind:value={Usn} placeholder="Enter your USN" class="input" />
		</div>

		<div>
			<label class="block font-medium text-gray-700">👤 Name</label>
			<input type="text" bind:value={name} placeholder="Enter your full name" class="input" />
		</div>

		<div>
			<label class="block font-medium text-gray-700">🔑 Password</label>
			<input type="password" bind:value={password} placeholder="Create a password" class="input" />
		</div>

		<button class="button" on:click={register}>Register</button>

		<p class="mt-4 text-center">
			Already have an account?
			<a href="/userlogin" class="link">Login here</a>
		</p>
	</div>
</div>

<style>
	/* Background with gradient */
	.bg-gradient {
		background: linear-gradient(135deg, #667eea, #764ba2);
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.card {
		background: white;
		border-radius: 12px;
		box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
		padding: 40px;
		width: 100%;
		max-width: 450px;
		transition: all 0.3s ease;
	}

	.card:hover {
		box-shadow: 0 12px 25px rgba(0, 0, 0, 0.15);
		transform: translateY(-5px);
	}

	.input {
		width: 100%;
		padding: 12px;
		margin: 8px 0;
		border: 1px solid #ccc;
		border-radius: 8px;
		font-size: 14px;
		transition: border 0.3s;
	}

	.input:focus {
		outline: none;
		border-color: #764ba2;
		box-shadow: 0 0 8px rgba(118, 75, 162, 0.2);
	}

	.button {
		width: 100%;
		padding: 12px;
		background: #4caf50;
		color: white;
		border: none;
		border-radius: 8px;
		font-size: 16px;
		cursor: pointer;
		transition: background 0.3s;
	}

	.button:hover {
		background: #45a049;
	}

	.link {
		color: #667eea;
		text-decoration: none;
		transition: color 0.3s;
	}

	.link:hover {
		color: #764ba2;
	}

	.message {
		font-size: 14px;
		margin: 10px 0;
		padding: 10px;
		border-radius: 6px;
		text-align: center;
	}

	.error {
		background: #ffdddd;
		color: #d8000c;
		border: 1px solid #d8000c;
	}

	.success {
		background: #ddffdd;
		color: #4caf50;
		border: 1px solid #4caf50;
	}
</style>
